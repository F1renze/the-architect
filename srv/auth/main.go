package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/etcd"

	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/f1renze/the-architect/srv/auth/handler"

	pb "github.com/f1renze/the-architect/srv/auth/proto"
)

func main() {
	cmsCli := common.Init()

	srvCfg, err := config.GetSrvConfig(constant.AuthSrvCfgName, cmsCli)
	if err != nil {
		log.Fatal("get auth service config failed", err)
	}

	reg := etcd.NewRegistry(infra.GetRegistryOptions())

	srv := micro.NewService(
		micro.Name(srvCfg.Name),
		micro.Registry(reg),
		micro.Version(srvCfg.Version),

		micro.RegisterTTL(constant.RegisterTTL),
		micro.RegisterInterval(constant.RegisterInterval),
	)
	srv.Init()

	_ = pb.RegisterAuthHandler(srv.Server(), handler.NewHandler())

	if err = srv.Run(); err != nil {
		log.Fatal("launch auth service failed", err)
	}
}
