package main

import (
	"github.com/f1renze/the-architect/common/constant/topic"
	"github.com/f1renze/the-architect/common/utils"
	"github.com/f1renze/the-architect/srv/auth/sub"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/broker/rabbitmq"

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

	b := rabbitmq.NewBroker(
		broker.Addrs(infra.GetRabbitMqAddr()),
	)
	reg := etcd.NewRegistry(infra.GetRegistryOptions())
	err = b.Init(
		broker.Registry(reg),
	)
	err2 := b.Connect()
	if err = utils.NoErrors(err, err2); err != nil {
		log.Fatal("rabbit mq connected error", err)
	}
	defer func() {
		if err = b.Disconnect(); err != nil {
			log.ErrorF("disconnect rabbit mq error", err)
		}
	}()

	srv := micro.NewService(
		micro.Name(srvCfg.Name),
		micro.Version(srvCfg.Version),

		micro.Broker(b),

		micro.Registry(reg),
		micro.RegisterTTL(constant.RegisterTTL),
		micro.RegisterInterval(constant.RegisterInterval),
	)
	srv.Init()

	subHandler := sub.NewSubscriber()
	err = micro.RegisterSubscriber(topic.ConfirmEmailTopic, srv.Server(), subHandler.SendRegistrationEmail())
	err2 = micro.RegisterSubscriber(topic.ConfirmMobileTopic, srv.Server(), subHandler.SendRegistrationSms())
	if err = utils.NoErrors(err, err2); err != nil {
		log.Fatal("register subscriber error", err)
	}

	_ = pb.RegisterAuthHandler(srv.Server(), handler.NewHandler())

	if err = srv.Run(); err != nil {
		log.Fatal("launch auth service failed", err)
	}
}
