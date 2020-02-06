package infra

import (
	"fmt"
	"github.com/f1renze/the-architect/common/utils/log"
	"sync"

	"github.com/micro/go-micro/registry"

	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/infra/db"
)

var (
	once sync.Once
	etcd *EtcdConfig
)

func Init(cmsCli config.CMSClient) {
	once.Do(func() {
		db.Init(cmsCli)

		etcd = new(EtcdConfig)
		err := cmsCli.Scan("infra.etcd", etcd)
		if err != nil {
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
