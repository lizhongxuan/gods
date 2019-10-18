package etcdservice

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/golang/glog"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var etcdTTL = 30
var etcdTimeout = time.Second * 30

var etcdEndpoints []string

var clientSession ClientSession

type ClientSession struct {
	*clientv3.Client

	mu      sync.Mutex
	session *concurrency.Session
}

func init() {
	if s := os.Getenv("ETCD_ENDPOINTS"); s != "" {
		etcdEndpoints = strings.Split(s, ",")
	} else {
		etcdEndpoints = []string{"http://127.0.0.1:2379"}
	}
}

func newClient() (*clientv3.Client, error) {
	config := clientv3.Config{
		Endpoints:            etcdEndpoints,
		DialTimeout:          etcdTimeout,
		DialKeepAliveTime:    etcdTimeout,
		DialKeepAliveTimeout: etcdTimeout,
	}
	glog.Info("config:", config)
	c, err := clientv3.New(config)
	return c, err
}

// 命令行指定etcd endpoints时设置
func SetEndpoints(endpoints string) {
	etcdEndpoints = strings.Split(endpoints, ",")
}

func (cs *ClientSession) Get() (*clientv3.Client, *concurrency.Session, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if cs.Client == nil {
		if etcd, err := newClient(); err != nil {
			glog.Error(err)
			return nil, nil, err
		} else {
			log.Print("client created")
			cs.Client = etcd
		}
	}

	if cs.session == nil {
		if session, err := concurrency.NewSession(cs.Client, concurrency.WithTTL(60)); err != nil {
			return nil, nil, err
		} else {
			log.Printf("session created, leaseid:%v", session.Lease())
			cs.session = session
			go cs.waitSession(session)
		}
	}
	return cs.Client, cs.session, nil
}

func (cs *ClientSession) waitSession(session *concurrency.Session) {
	<-session.Done()

	cs.mu.Lock()
	defer cs.mu.Unlock()
	if cs.session == session {
		cs.session = nil
		log.Printf("session done, leaseid:%v", session.Lease())
	}
}

type Client interface {
	clientv3.KV
	clientv3.Watcher
}

//返回clientv3.KV，不暴露Close等方法
func GetClient() (Client, error) {
	cli, _, err := clientSession.Get()
	return cli, err
}

//返回clientv3.KV，不暴露Close等方法
func GetClientSession() (clientv3.KV, *concurrency.Session, error) {
	cli, session, err := clientSession.Get()
	return cli, session, err
}

// GetClientWithoutLease 获取etcd client，不需要建立session
func GetClientWithoutLease() (*clientv3.Client, error) {
	cli, _, err := clientSession.Get()
	return cli, err
}

func NewSession(ttl int) (*concurrency.Session, error) {
	if cli, _, err := clientSession.Get(); err != nil {
		return nil, err
	} else if session, err := concurrency.NewSession(cli, concurrency.WithTTL(ttl)); err != nil {
		return nil, err
	} else {
		return session, nil
	}
}

func NewSessionWithoutTTL() (clientv3.KV, *concurrency.Session, error) {
	if cli, _, err := clientSession.Get(); err != nil {
		return nil, nil, err
	} else if session, err := concurrency.NewSession(cli); err != nil {
		return nil, nil, err
	} else {
		return cli, session, nil
	}
}
