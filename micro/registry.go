package micro

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type Registry interface {
	Address() []string
}

func etcdRegistry(c *EtcdRegistry) (reg registry.Registry) {
	return etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = c.Address()
	})
}
