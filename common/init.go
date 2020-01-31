package common

import (
	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils/log"
)

func Init() config.CMSClient {
	log.Init()

	cmsCli := config.GetCMSClient()

	infra.Init(cmsCli)

	return cmsCli
}
