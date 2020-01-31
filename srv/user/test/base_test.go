package test

import (
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
	common.Init()
	userModel = model.NewModel()

	m.Run()
}
