package nacosx

import (
	"sync"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
)

type NacosConf[T any] struct {
	lock  *sync.RWMutex
	data  T
	kconf config.Config
	log   *log.Helper
}

func NewNacosConf[T any](kconf config.Config, logger log.Logger) *NacosConf[T] {
	return &NacosConf[T]{
		lock:  new(sync.RWMutex),
		kconf: kconf,
		log:   log.NewHelper(log.With(logger, "x_module", "nacos")),
	}
}

func (n *NacosConf[T]) Get() T {
	n.lock.RLock()
	defer n.lock.RUnlock()
	return n.data
}

func (n *NacosConf[T]) Scan() error {
	n.lock.Lock()
	defer n.lock.Unlock()
	err := n.kconf.Scan(&n.data)
	if err != nil {
		n.log.Error(err)
		return err
	}
	return nil
}

func (n *NacosConf[T]) Watch(key string) error {
	if err := n.kconf.Watch(key, func(key string, value config.Value) {
		n.log.Infof("config changed: %s = %v\n", key, value)
		err := n.Scan()
		if err != nil {
			n.log.Error(err)
			return
		}
		// 在这里写回调的逻辑
	}); err != nil {
		n.log.Error(err)
		return err
	}
	return nil
}
