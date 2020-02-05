package user

import (
	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"os"
	"testing"
)

var (
	cmsCli config.CMSClient
)

func TestMain(m *testing.M) {
	os.Setenv(constant.CfgCenterAddrEnv, "127.0.0.1:9689")
	cmsCli = common.Init()

	m.Run()
}
