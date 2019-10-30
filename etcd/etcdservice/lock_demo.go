package etcdservice

import "github.com/coreos/etcd/clientv3/concurrency"

func getIDEtcdWithLock(ctx context.Context) error {
	session, err := etcdservice.NewSession(15)
	if err != nil {
		glog.Error(" new session failed")
		return err
	}
	defer session.Close()

	mutexKey := fmt.Sprintf("/mutex/idKeyManager/%s", ctx.Value("lockKey"))
	mutex := concurrency.NewMutex(session, mutexKey)
	if err := mutex.Lock(ctx); err != nil {
		glog.Errorf("etcd lock failed:%s", err)
		return err
	}
	defer mutex.Unlock(ctx)

	return nil
}
