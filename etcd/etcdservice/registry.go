package etcdservice

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"sync"
	"time"
)

type Registry struct {
	mu         sync.Mutex
	client     *clientv3.Client
	session    *concurrency.Session
	key, value string
	Status     bool
}

func RegisterForever(name, addr string, data string) {
	go func() {
		for {
			reg, err := Register(name, addr, data)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			reg.Wait()
		}
	}()
}

//向etcd注册一个服务
func Register(name, addr string, data string) (*Registry, error) {
	//获取etcd连接客户端
	client, _, err := clientSession.Get()
	if err != nil {
		return nil, err
	}
	registry := &Registry{
		client: client,
		key:    servicePrefix + name + "/" + addr,
		value:  data,
	}

	if session, err := registry.newSession(); err != nil {
		return nil, err
	} else {
		registry.session = session
		registry.Status = true
	}

	//go registry.waitSession()
	return registry, nil
}

func (registry *Registry) String() string {
	return registry.key
}

func (registry *Registry) Close() {
	registry.closeSession()
}

func (registry *Registry) closeSession() {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	if registry.session != nil {
		registry.session.Close()
		log.Printf("server %s unregistered", registry.key)
		registry.session = nil
		registry.Status = false
	}
}

//wait for session closed
func (registry *Registry) Wait() {
	<-registry.session.Done()
	fmt.Println("closeSession key:", registry.key, "  value:", registry.value)
	registry.closeSession()
}

//建立有时间时效(租约)的链接
func (registry *Registry) newSession() (*concurrency.Session, error) {
	session, err := concurrency.NewSession(registry.client, concurrency.WithTTL(etcdTTL))
	if err != nil {
		return nil, err
	}

	leaseID := session.Lease()
	if _, err := registry.client.Put(context.Background(), registry.key, registry.value, clientv3.WithLease(leaseID)); err != nil {
		return nil, err
	} else {
		log.Printf("server %s registered", registry.key)
	}
	return session, nil
}

func (registry *Registry) Done() <-chan struct{} {
	return registry.session.Done()
}
