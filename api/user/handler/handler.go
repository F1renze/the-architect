package handler

import (
	"context"
	"github.com/f1renze/the-architect/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"

	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	userPb "github.com/f1renze/the-architect/srv/user/proto"
	authPb "github.com/f1renze/the-architect/srv/auth/proto"
)

type UserApi interface {
	Login(*gin.Context)
	Logout(*gin.Context)
	Register(*gin.Context)
}

func NewHandler(cmsCli config.CMSClient) (UserApi, error) {
	userCfg, err := config.GetSrvConfig(constant.UserSrvCfgName, cmsCli)
	authCfg, err := config.

	if err != nil {
		return nil, err
	}
	return &Handler{
		userCli: userPb.NewUserService(userCfg.Name, client.DefaultClient),
		authCli: authPb.NewAuthService()
	}, nil
}

type Handler struct {
	userCli userPb.UserService
	authCli authPb.AuthService
}

func (h *Handler) Login(c *gin.Context) {

}

func (h *Handler) Logout(c *gin.Context) {

}

func (h *Handler) Register(c *gin.Context) {
	username := c.PostForm("username")
	// using email as default auth type
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	confirmPwd := c.PostForm("confirm_password")

	if utils.ValidateEmailFormat(email) || pwd != confirmPwd || username == "" || pwd == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"msg": "请求参数错误",
		})
		return
	}

	ctx := context.TODO()
	req := &userPb.Request{
		User: &userPb.User{
			Name: username,
		},
	}
	resp, err := h.userCli.CreateUser(ctx, req)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"msg": "服务器出小差了，请联系管理员",
		})
		return
	}
	if !resp.Success {
		if resp.Error.Code == constant.MySQLDupEntryErr {
			c.AbortWithStatusJSON(400, gin.H{
				"msg": "用户名已被注册",
			})
			return
		}
		c.AbortWithStatusJSON(int(resp.Error.Code), gin.H{
			"msg": resp.Error.Detail,
		})
		return
	}

	resp, err :=
}
