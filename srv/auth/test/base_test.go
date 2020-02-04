package test

import (
	"github.com/f1renze/the-architect/common/constant"
	"os"
	"testing"

	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/srv/auth/model"
)

var (
	authModel model.AuthModel
)

func TestMain(m *testing.M) {
	os.Setenv(constant.CfgCenterAddrEnv, "127.0.0.1:9689")
	common.Init()
	authModel = model.NewModel()

	m.Run()
}