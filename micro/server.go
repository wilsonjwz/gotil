package micro

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

func InitServer(name, version string, registry *EtcdRegistry, broker *MqttBroker, fn func(s server.Server)) {
	var (
		s    micro.Service
		opts []micro.Option
	)
	opts = []micro.Option{
		micro.Name(name),
		micro.WrapHandler(LogWrapper),
		micro.Version(version),
	}
	if registry != nil {
		opts = append(opts, micro.Registry(etcdRegistry(registry)))
	}
	if broker != nil {
		opts = append(opts, micro.Broker(mqttBroker(broker)))
	}
	s = micro.NewService(opts...)
	//s.Init()
	fn(s.Server())

	go func() {
		if err := s.Run(); err != nil {
			panic(fmt.Sprintf("[%s] micro server run error(%+v).", name, err))
		}
	}()
}
