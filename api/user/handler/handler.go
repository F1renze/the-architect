package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/etcd"

	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils"
	"github.com/f1renze/the-architect/common/utils/log"
	authPb "github.com/f1renze/the-architect/srv/auth/proto"
	userPb "github.com/f1renze/the-architect/srv/user/proto"
)

type UserApi interface {
	Login(*gin.Context)
	Logout(*gin.Context)
	Register(*gin.Context)
}

func NewHandler(cmsCli config.CMSClient) (UserApi, error) {
	userCfg, err := config.GetSrvConfig(constant.UserSrvCfgName, cmsCli)
	authCfg, err := config.GetSrvConfig(constant.AuthSrvCfgName, cmsCli)

	if err != nil {
		return nil, err
	}
	return &Handler{
		userCli: userPb.NewUserService(userCfg.Name, client.NewClient(
			client.Registry(etcd.NewRegistry(infra.GetRegistryOptions())),
		)),
		authCli: authPb.NewAuthService(authCfg.Name, client.NewClient(
			client.Registry(etcd.NewRegistry(infra.GetRegistryOptions())),
		)),
	}, nil
}

type Handler struct {
	userCli userPb.UserService
	authCli authPb.AuthService
}

func (h *Handler) Login(c *gin.Context) {
	// todo validate
}

func (h *Handler) Logout(c *gin.Context) {

}

func (h *Handler) Register(c *gin.Context) {
	username := c.PostForm("username")
	// using email as default auth type
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	confirmPwd := c.PostForm("confirm_password")

	body := gin.H{}
	body["code"], body["msg"] = errno.DecodeInt32Err(errno.OK)

	switch {
	case !utils.ValidateEmailFormat(email):
		err := errno.InvalidEmail.Add(`"` + email + `"`)
		body["code"], body["msg"] = errno.DecodeInt32Err(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, body)
		return
	case pwd == "":
		body["msg"] = "密码不可为空"
		c.AbortWithStatusJSON(http.StatusBadRequest, body)
		return
	case pwd != confirmPwd:
		body["msg"] = "两次密码不一致"
		c.AbortWithStatusJSON(http.StatusBadRequest, body)
		return
	case username == "":
		body["code"], body["msg"] = errno.DecodeInt32Err(errno.InvalidUserName)
		c.AbortWithStatusJSON(http.StatusBadRequest, body)
		return
	default:
	}

	ctx := context.TODO()
	req := &userPb.Request{
		User: &userPb.User{
			Name: username,
		},
	}
	resp, err := h.userCli.CreateUser(ctx, req)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errno.RemoteServiceErr.Code,
			"msg":  errno.RemoteServiceErr.Message,
		})
		log.ErrorF("call user srv failed, %s", err)
		return
	}
	if !resp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": resp.Error.Code,
			"msg":  resp.Error.Detail,
		})
		return
	}

	authResp, err := h.authCli.AddLoginCredential(ctx, &authPb.Request{
		Login: true,
		Info: &authPb.AuthInfo{
			Uid:        resp.User.Id,
			AuthType:   authPb.AuthType_Email,
			AuthId:     email,
			Credential: pwd,
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errno.RemoteServiceErr.Code,
			"msg":  errno.RemoteServiceErr.Message,
		})
		log.ErrorF("call auth srv failed, %s", err)
		return
	}
	if !authResp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": authResp.Error.Code,
			"msg":  authResp.Error.Detail,
		})
		return
	}

	// TODO 消息队列发送验证邮件
	// 生成 jwt 返回
	// 若只有一个登录方式且未 Verified 禁止任何带权限操作

	c.JSON(http.StatusOK, gin.H{
		"code": errno.OK.Code,
		"msg":  errno.OK.Message,
	})
}
