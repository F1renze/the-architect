package main

import (
	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/f1renze/the-architect/srv/casbin/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/etcd"

	pb "github.com/f1renze/the-architect/srv/casbin/proto"
)

func main() {
	cmsCli := common.Init()

	srvCfg, err := config.GetSrvConfig(constant.CasbinSrvCfgName, cmsCli)
	if err != nil {
		log.Fatal("get config failed", err)
	}

	reg := etcd.NewRegistry(infra.GetRegistryOptions())

	srv := micro.NewService(
		micro.Name(srvCfg.Name),
		micro.Version(srvCfg.Version),

		micro.Registry(reg),
		micro.RegisterTTL(constant.RegisterTTL),
		micro.RegisterInterval(constant.RegisterInterval),
	)
	srv.Init()

	_ = pb.RegisterCasbinHandler(srv.Server(), handler.NewHandler())
	if err = srv.Run(); err != nil {
		log.Fatal("launch casbin srv failed", err)
	}
}
