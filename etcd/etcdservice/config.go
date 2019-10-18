package etcdservice

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/golang/glog"
)

var gConfigCache = &ConfigCache{
	kvs: map[string][]byte{},
}

type ConfigCache struct {
	once     sync.Once
	revision int64
	kvs      map[string][]byte
}

func GetConfig(key string) ([]byte, int64, bool) {
	return gConfigCache.Get(key)
}

func (cache *ConfigCache) Get(key string) ([]byte, int64, bool) {
	cache.once.Do(cache.reload)
	val, ok := cache.kvs[key]
	return val, cache.revision, ok
}

func (cache *ConfigCache) GetAll() (map[string][]byte, int64) {
	cache.once.Do(cache.reload)
	return cache.kvs, cache.revision
}

func (cache *ConfigCache) reload() {
	cli, err := GetClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	resp, err := cli.Get(ctx, "/config/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	for _, kv := range resp.Kvs {
		key := string(kv.Key)
		cache.kvs[key] = kv.Value
	}
	cache.revision = resp.Header.GetRevision()
}

//从etcd中加载一个配置, 失败则退出
func MustLoadConfig(name string, cfg interface{}) {
	key := "/config/" + name
	val, _, ok := gConfigCache.Get(key)
	if !ok {
		glog.Fatal("config ", key, " not exists!")
	}

	if err := toml.Unmarshal(val, cfg); err != nil {
		glog.Fatal("load config ", key, " error:", err)
	}
}

//从etcd中加载一个配置，不存在则使用默认值; 配置格式错误的话，仍然会panic
func LoadConfig(name string, cfg interface{}) bool {
	key := "/config/" + name
	val, _, ok := gConfigCache.Get(key)
	if !ok {
		return false
	}

	if err := toml.Unmarshal(val, cfg); err != nil {
		glog.Fatal("load config ", key, " error:", err)
	}
	return true
}

type ConfigWatcher struct {
	ctx    context.Context
	cancel context.CancelFunc
	next   chan []byte
}

func NewConfigWatcher(parent context.Context, name string) *ConfigWatcher {
	ctx, cancel := context.WithCancel(parent)
	w := &ConfigWatcher{
		ctx:    ctx,
		cancel: cancel,
		next:   make(chan []byte, 1),
	}
	go w.watch(name)
	return w
}

func (this *ConfigWatcher) watch(name string) {
	key := "/config/" + name
	val, revision, _ := gConfigCache.Get(key)
	//klog.Printf("config[%s] loaded, rev=%d val=%s", name, revision, val)
	this.next <- val

	for {
		cli, err := GetClient()
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		ch := cli.Watch(this.ctx, key, clientv3.WithRev(revision+1), clientv3.WithPrevKV())
		for resp := range ch {
			for _, ev := range resp.Events {
				switch ev.Type {
				case mvccpb.DELETE:
					glog.Infof("config[%s] deleted, rev=%d", name, ev.Kv.ModRevision)
					this.next <- nil
				case mvccpb.PUT:
					glog.Infof("config[%s] updated, rev=%d", name, ev.Kv.ModRevision)
					this.next <- ev.Kv.Value
				}
			}
		}
	}
}

func (this *ConfigWatcher) Stop() {
	this.cancel()
}

func (this *ConfigWatcher) Next() ([]byte, error) {
	select {
	case val := <-this.next:
		return val, nil
	case <-this.ctx.Done():
		return nil, this.ctx.Err()
	}
}

func WatchConfig(name string, watcher func([]byte)) *ConfigWatcher {
	w := NewConfigWatcher(context.Background(), name)
	go func() {
		for {
			val, err := w.Next()
			if err != nil {
				return
			}
			watcher(val)
		}
	}()
	return w
}
