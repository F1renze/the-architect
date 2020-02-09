package handler

import (
	"context"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/etcd"

	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils/log"
	authPb "github.com/f1renze/the-architect/srv/auth/proto"
	userPb "github.com/f1renze/the-architect/srv/user/proto"
)

func init() {
	govalidator.CustomTypeTagMap.Set("pwdeq", func(i interface{}, o interface{}) bool {
		if i.(string) == o.(RegisterForm).Pwd {
			return true
		}
		return false
	})
}

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

type LoginForm struct {
	AuthId     string `form:"auth_id" binding:"required" valid:"stringlength(6|255)"`
	Credential string `form:"credential" binding:"required" valid:"stringlength(6|255)"`
}

// todo login using username, email, phone
// default create a username auth record
func (h *Handler) Login(c *gin.Context) {
	var form LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "表单字段不可为空",
		})
		return
	}

	_, err = govalidator.ValidateStruct(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx := context.TODO()
	resp, err := h.authCli.CheckCredential(ctx, &authPb.Request{
		Login: true,
		Info: &authPb.AuthInfo{
			AuthId:     form.AuthId,
			Credential: form.Credential,
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errno.RemoteServiceErr.Code,
			"msg":  errno.RemoteServiceErr.Message,
		})
		log.ErrorF("api.user.login: call auth srv failed, %s", err)
		return
	}
	if !resp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": resp.Error.Code,
			"msg":  resp.Error.Detail,
		})
		return
	}

	resp, err = h.authCli.SignOn(ctx, &authPb.Request{
		Info: &authPb.AuthInfo{
			Uid:    resp.Info.Uid,
			AuthId: resp.Info.AuthId,
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errno.GetRespFromErr(errno.RemoteServiceErr))
		log.ErrorF("api.user.login: call srv.auth.signOn failed", err)
		return
	}
	if !resp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": resp.Error.Code,
			"msg": resp.Error.Detail,
		})
		return
	}

	c.SetCookie("_token", resp.Token, int(constant.JwtExpiredTime.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"code": errno.OK.Code,
		"msg":  errno.OK.Message,
		"token": resp.Token,
	})
}

// 注销 token
func (h *Handler) Logout(c *gin.Context) {
	token, err := c.Cookie("_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code": -1,
			"msg": "You must login first",
		})
	}

	ctx := context.TODO()
	resp, err := h.authCli.SignOff(ctx, &authPb.Request{
		Token: token,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errno.GetRespFromErr(errno.RemoteServiceErr))
		log.ErrorF("api.user.logout: call srv.auth.signOff failed", err)
		return
	}
	if !resp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": resp.Error.Code,
			"msg": resp.Error.Detail,
		})
		return
	}
	// refresh cookie
	c.SetCookie("_token", "", int(constant.JwtExpiredTime.Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, errno.GetRespFromErr(errno.OK.Add("you're already logout.")))
}

type RegisterForm struct {
	Name       string `form:"username" binding:"required" valid:"stringlength(1|32)"`
	Email      string `form:"email" binding:"required" valid:"email"`
	Pwd        string `form:"password" binding:"required" valid:"stringlength(6|255)"`
	ConfirmPwd string `form:"confirm_password" binding:"required" valid:"pwdeq"`
}

func (h *Handler) Register(c *gin.Context) {
	var form RegisterForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "表单字段不可为空",
		})
		return
	}

	_, err = govalidator.ValidateStruct(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx := context.TODO()
	resp, err := h.userCli.CreateUser(ctx, &userPb.Request{
		User: &userPb.User{
			Name: form.Name,
		},
	})
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
			AuthId:     form.Email,
			Credential: form.Pwd,
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errno.RemoteServiceErr.Code,
			"msg":  errno.RemoteServiceErr.Message,
		})
		log.ErrorF("api.user.register: call auth srv failed, %s", err)
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
	authResp, err = h.authCli.SignOn(ctx, &authPb.Request{
		Info: &authPb.AuthInfo{
			Uid:    resp.User.Id,
			AuthId: form.Email,
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errno.RemoteServiceErr.Code,
			"msg":  errno.RemoteServiceErr.Message,
		})
		log.ErrorF("api.user.register: call srv.auth.signOn failed", err)
		return
	}
	if !authResp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": authResp.Error.Code,
			"msg":  authResp.Error.Detail,
		})
		return
	}
	// cookie expires base on server side, set an expiry date
	// cookie maxAge base on client side, set time in secs for how long it will be deleted
	// https://mrcoles.com/blog/cookies-max-age-vs-expires/
	c.SetCookie("_token", authResp.Token, int(constant.JwtExpiredTime.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"code":  errno.OK.Code,
		"msg":   errno.OK.Message,
		"token": authResp.Token,
	})
}
