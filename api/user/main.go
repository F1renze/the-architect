package main

import (
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"

	"github.com/f1renze/the-architect/api/user/handler"
	"github.com/f1renze/the-architect/api/user/router"
	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils"
	"github.com/f1renze/the-architect/common/utils/log"
)

func main() {
	cmsCli := common.Init()

	webCfg, err := config.GetWebConfig(constant.UserApiCfgName, cmsCli)
	h, err2 := handler.NewHandler(cmsCli)

	if err = utils.NoErrors(err, err2); err != nil {
		log.Fatal("初始化 api.user 相关配置失败!", err)
	}

	reg := etcd.NewRegistry(infra.GetRegistryOptions())

	srv := web.NewService(
		web.Name(webCfg.Name),
		web.Version(webCfg.Version),
		web.Address(webCfg.Address),

		web.Registry(reg),
		web.RegisterInterval(constant.RegisterInterval),
		web.RegisterTTL(constant.RegisterTTL),
	)
	_ = srv.Init()

	srv.Handle("/", router.Default(h))

	if err = srv.Run(); err != nil {
		log.Fatal("launched api.user failed", err)
	}
}
