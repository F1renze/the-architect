package main

import (
	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/f1renze/the-architect/srv/user/handler"
	pb "github.com/f1renze/the-architect/srv/user/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/etcd"
)


func main() {
	cmsCli := common.Init()

	srvCfg, err := config.GetSrvConfig(constant.UserSrvCfgName, cmsCli)
	if err != nil {
		log.Fatal("获取服务配置失败", err)
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

	_ = pb.RegisterUserHandler(srv.Server(), handler.NewHandler())

	if err = srv.Run(); err != nil {
		log.Fatal("启动 user 服务失败", err)
	}
}
