package infra

import (
	"fmt"
	"github.com/f1renze/the-architect/common/utils"
	"sync"

	"github.com/micro/go-micro/registry"

	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/infra/db"
	"github.com/f1renze/the-architect/common/infra/redis"
	"github.com/f1renze/the-architect/common/utils/log"
)

var (
	once sync.Once
	etcd *EtcdConfig
	rabbit *RabbitMqConfig
)

func Init(cmsCli config.CMSClient) {
	once.Do(func() {
		db.Init(cmsCli)
		redis.Init(cmsCli)

		etcd = new(EtcdConfig)
		rabbit = new(RabbitMqConfig)
		err := cmsCli.Scan("infra.etcd", etcd)
		err2 := cmsCli.Scan("infra.rabbitmq", rabbit)
		if err = utils.NoErrors(err, err2); err != nil {
			log.Fatal("initialize infra config failed", err)
		}
	})
}

type EtcdConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func GetRegistryOptions() func(*registry.Options) {
	return func(ops *registry.Options) {
		ops.Addrs = []string{
			fmt.Sprintf("%s:%d", etcd.Host, etcd.Port),
		}
	}
}


type RabbitMqConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
}

func GetRabbitMqAddr() string{
	return fmt.Sprintf("amqp://%s:%s@%s:%d", rabbit.User, rabbit.Password, rabbit.Host, rabbit.Port)
}