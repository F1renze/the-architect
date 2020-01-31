package config

import (
	"fmt"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-plugins/config/source/grpc"
	"os"

	"strings"
	"sync"
)

var (
	once sync.Once
	c    *Client
)

// 配置中心客户端
type CMSClient interface {
	// 读取配置并赋值到对应 struct
	Scan(configName string, config interface{}) error
	// 增加配置源接口
	AppendSources(sources ...source.Source) error
}

type Client struct {
	// manages all config, abstracting away manager, encoders and the reader.
	manager config.Config
}

func (c *Client) Scan(cfgName string, config interface{}) error {
	keys := strings.Split(cfgName, ".")
	v := c.manager.Get(keys...)
	if v == nil {
		return fmt.Errorf("%s: 配置不存在", cfgName)
	}
	log.InfoF("查询配置: %s", cfgName)

	err := v.Scan(config)
	if err != nil {
		log.Error("查询配置失败", log.Any{
			"err":      err,
			"cfg_name": cfgName,
		})
	}
	return err
}

func (c *Client) AppendSources(sources ...source.Source) error {
	log.InfoF("增加配置源")
	return c.manager.Load(sources...)
}

func GetCMSClient() CMSClient {
	once.Do(func() {
		addr := os.Getenv(constant.CfgCenterAddrEnv)
		log.InfoF("config address: %s", addr)

		src := grpc.NewSource(
			// grpc 配置服务地址
			grpc.WithAddress(addr),
			// 默认的 req.path
			grpc.WithPath(constant.CfgPrefix),
		)
		c = new(Client)
		c.manager = config.NewConfig()

		if err := c.manager.Load(src); err != nil {
			log.Fatal("CMS Client: load src failed", err)
		}

		go func() {
			watcher, err := c.manager.Watch()
			if err != nil {
				log.Fatal("CMS Client: 监听配置失败...", err)
			}
			log.InfoF("CMS Client: 监听配置中...")

			for {
				v, err := watcher.Next()
				if err != nil {
					log.Fatal("CMS Client: 监听配置变动错误", err)
				}
				log.Info("配置变化", log.Any{
					"changes": string(v.Bytes()),
				})
			}
		}()
	})

	return c
}
