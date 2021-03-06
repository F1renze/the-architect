package test

import (
	"github.com/f1renze/the-architect/common/constant"
	"os"
	"testing"

	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/srv/user/model"
	userPb "github.com/f1renze/the-architect/srv/user/proto"
)

var (
	userModel model.UserModel
	userCli   userPb.UserService
)

func TestMain(m *testing.M) {
	os.Setenv(constant.CfgCenterAddrEnv, "127.0.0.1:9689")
	common.Init()
	userModel = model.NewModel()

	m.Run()
}
