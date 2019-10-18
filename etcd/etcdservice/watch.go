package etcdservice

import (
	"bytes"
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/golang/glog"
	"google.golang.org/grpc/naming"
	"time"
)

var servicePrefix = "/service/"
var sessionPrefix = "/session/"
var sessionEditFormat = "edit/%s"
var sessionPreviewFormat = "preview/%s/v/%d"
var sessionCompareFormat = "compare/%s/v/%d"
var sessionDebugFormat = "debug/%s@%s"

var defaultResolver = &Resolver{}

func GetSessionPreviewFormat() string {
	return sessionPreviewFormat
}

func GetSessionEditFormat() string {
	return sessionEditFormat
}

func GetSessionDebugFormat() string {
	return sessionDebugFormat
}

func GetSessionCompareFormat() string {
	return sessionCompareFormat
}

func GetSessionPrefix() string {
	return sessionPrefix
}

type Resolver struct {
	naming.Resolver
}

func NewResolver() naming.Resolver {
	return defaultResolver
}

func (resolver *Resolver) Resolve(name string) (naming.Watcher, error) {
	return newWatcher(name), nil
}

type watchResponse struct {
	updates []*naming.Update
	err     error
}

type Watcher struct {
	naming.Watcher

	name  string
	addrs map[string]bool
	next  chan []*naming.Update
	close chan bool
}

func newWatcher(name string) *Watcher {
	watcher := &Watcher{
		name:  name,
		addrs: map[string]bool{},
		next:  make(chan []*naming.Update, 1),
		close: make(chan bool, 1),
	}
	go func() {
		for {
			watcher.watch()
			time.Sleep(time.Second)
		}
	}()
	return watcher
}

func (watcher *Watcher) Next() ([]*naming.Update, error) {
	next := <-watcher.next
	return next, nil
}

func (watcher *Watcher) Close() {
	select {
	case watcher.close <- true:
	default:
	}
}

func getKeyAddr(key []byte) string {
	addr := ""
	if idx := bytes.LastIndexByte(key, '/'); idx >= 0 {
		addr = string(key[idx+1:])
	}
	return addr
}

func (watcher *Watcher) initWatch() (*clientv3.Client, *concurrency.Session, error) {
	if etcd, err := newClient(); err != nil {
		return nil, nil, err
	} else if session, err := concurrency.NewSession(etcd, concurrency.WithTTL(60)); err != nil {
		etcd.Close()
		return nil, nil, err
	} else {
		glog.Infof("watcher[%s] session started, leaseid:%v", watcher.name, session.Lease())
		return etcd, session, nil
	}
}

func (watcher *Watcher) watch() {
	glog.Info("Watcher...")
	//复用etcd client
	cli, session, err := clientSession.Get()
	if err != nil {
		glog.Infof("watcher[%s] get client session failed: %s", watcher.name, err)
		return
	}

	prefix := servicePrefix + watcher.name + "/"
	resp, err := cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		glog.Infof("watcher[%s] list server failed: %s", watcher.name, err)
		return
	}
	updates := make([]*naming.Update, 0, len(resp.Kvs)+len(watcher.addrs))

	oldAddrs := watcher.addrs
	watcher.addrs = map[string]bool{}
	//添加新服务器
	for _, kv := range resp.Kvs {
		addr := getKeyAddr(kv.Key)

		//已存在的服务器，不再添加
		if oldAddrs[addr] {
			delete(oldAddrs, addr)
			continue
		}

		update := &naming.Update{
			Op:   naming.Add,
			Addr: addr,
		}
		updates = append(updates, update)
		watcher.addrs[update.Addr] = true
		glog.Infof("init add server:%s, address:%s", watcher.name, update.Addr)
	}
	//删除无用的服务器
	for addr := range oldAddrs {
		update := &naming.Update{
			Op:   naming.Delete,
			Addr: addr,
		}
		updates = append(updates, update)
		glog.Infof("init remove server:%s, address:%s", watcher.name, update.Addr)
	}

	lastRevision := resp.Header.GetRevision()
	watcher.next <- updates

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//watch changes, lastRevision已经观测到， 这里需要+1
	events := cli.Watch(ctx, prefix, clientv3.WithPrefix(), clientv3.WithRev(lastRevision+1))
	for {
		select {
		case <-session.Done():
			glog.Errorf("watcher[%s] session done, leaseid:%v", watcher.name, session.Lease())
			return

		case resp := <-events:
			if err := resp.Err(); err != nil {
				glog.Errorf("watcher[%s] error: %s", watcher.name, err)
				return
			}

			updates := make([]*naming.Update, 0, len(resp.Events))
			for _, ev := range resp.Events {
				addr := ""
				if idx := bytes.LastIndexByte(ev.Kv.Key, '/'); idx >= 0 {
					addr = string(ev.Kv.Key[idx+1:])
				}

				switch ev.Type {
				case mvccpb.DELETE:
					update := &naming.Update{
						Op:   naming.Delete,
						Addr: addr,
					}
					updates = append(updates, update)
					delete(watcher.addrs, addr)
					glog.Infof("remove server:%s, address:%s", watcher.name, addr)
				case mvccpb.PUT:
					update := &naming.Update{
						Op:   naming.Add,
						Addr: addr,
					}
					updates = append(updates, update)
					watcher.addrs[addr] = true
					glog.Infof("add server:%s, address:%s", watcher.name, addr)
				}
			}
			watcher.next <- updates
		case <-watcher.close:
			return //don't call cancel() here
		}
	}
}
