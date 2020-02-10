package main

import (
	"github.com/f1renze/the-architect/common/constant"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/common/constant/topic"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/f1renze/the-architect/srv/broker/sub"
)

func main() {
	common.Init()

	b := rabbitmq.NewBroker(
		broker.Addrs(infra.GetRabbitMqAddr()),
	)
	reg := etcd.NewRegistry(infra.GetRegistryOptions())

	err := b.Init(
		broker.Registry(reg),
	)
	err2 := b.Connect()
	if err = utils.NoErrors(err, err2);err != nil {
		log.Fatal("", err)
	}
	defer func() {
		if err = b.Disconnect(); err != nil {
			log.ErrorF("", err)
		}
	}()

	subHandler := sub.NewSubscriber()

	srv := micro.NewService(
		micro.Name("micro.arch.broker"),
		micro.Version("latest"),

		micro.Broker(b),

		micro.Registry(reg),
		micro.RegisterTTL(constant.RegisterTTL),
		micro.RegisterInterval(constant.RegisterInterval),
	)

	srv.Init()

	err = micro.RegisterSubscriber(topic.ConfirmEmailTopic, srv.Server(), subHandler.SendRegistrationEmail())
	err2 = micro.RegisterSubscriber(topic.ConfirmMobileTopic, srv.Server(), subHandler.SendRegistrationSms())
	if err = utils.NoErrors(err, err2); err != nil {
		log.Fatal("", err)
	}

	if err = srv.Run(); err != nil {
		log.Fatal("", err)
	}
}
