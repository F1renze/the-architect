package main

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"

	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils/log"
	pb "github.com/f1renze/the-architect/srv/casbin/proto"
)

func InitCasbin() {
	rpcCli := client.NewClient(
		client.Registry(etcd.NewRegistry(infra.GetRegistryOptions())),
	)
	casbinCli := pb.NewCasbinService("micro.arch.srv.casbin", rpcCli)

	ctx := context.TODO()
	resp, err := casbinCli.NewAdapter(ctx, &pb.NewAdapterRequest{
		DriverName: "mysql",
	})
	if err != nil {
		log.Fatal("add adapter error", err)
	}
	adapterHandler := resp.Handler
//	todo create user in superadmin domain
}


func main() {
	common.Init()
	_ = plugin.Register(cors.NewPlugin())

	cmd.Init(
		micro.RegisterTTL(constant.RegisterTTL),
		micro.RegisterInterval(constant.RegisterInterval),
	)
}
