package main

import (

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"

	"github.com/f1renze/the-architect/common/constant"
)

const srvName = "infra.gateway"

func main() {
	_ = plugin.Register(cors.NewPlugin())

	cmd.Init(
		micro.RegisterTTL(constant.RegisterTTL),
		micro.RegisterInterval(constant.RegisterInterval),
	)
}
