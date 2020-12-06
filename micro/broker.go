package micro

import (
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/mqtt/v2"
	"github.com/micro/go-plugins/broker/nsq/v2"
	"github.com/micro/go-plugins/broker/redis/v2"
)

type Broker interface {
	Address() []string
}

func nsqBroker(c *NsqBroker) broker.Broker {
	return nsq.NewBroker(func(options *broker.Options) {
		options.Addrs = c.Addrs
	})
}

func redisBroker(c *RedisBroker) broker.Broker {
	return redis.NewBroker(func(options *broker.Options) {
		options.Addrs = c.Addrs
	})
}

func mqttBroker(c Broker) broker.Broker {
	return mqtt.NewBroker(func(options *broker.Options) {
		options.Addrs = c.Address()
		options.Secure = true
	})
}
