package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/constant/topic"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/common/infra"
	"github.com/f1renze/the-architect/common/utils"
	"github.com/f1renze/the-architect/common/utils/log"
	authPb "github.com/f1renze/the-architect/srv/auth/proto"
	msgPb "github.com/f1renze/the-architect/srv/broker/proto"
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
	SendSmsCode(*gin.Context)
}

func NewHandler(cmsCli config.CMSClient) (UserApi, error) {
	userCfg, err := config.GetSrvConfig(constant.UserSrvCfgName, cmsCli)
	authCfg, err2 := config.GetSrvConfig(constant.AuthSrvCfgName, cmsCli)

	b := rabbitmq.NewBroker(
		broker.Addrs(infra.GetRabbitMqAddr()),
	)
	reg := etcd.NewRegistry(infra.GetRegistryOptions())
	err3 := b.Init(
		broker.Registry(reg),
	)
	err4 := b.Connect()
	if err = utils.NoErrors(err, err2, err3, err4); err != nil {
		return nil, err
	}

	defer func() {
		if err = b.Disconnect(); err != nil {
			log.ErrorF("[api.user.handler::NewHandler] disconnect mq error", err)
		}
	}()

	rpcClient := client.NewClient(
		client.Registry(reg),
		client.Broker(b),
	)

	return &Handler{
		userCli: userPb.NewUserService(userCfg.Name, rpcClient),
		authCli: authPb.NewAuthService(authCfg.Name, rpcClient),

		emailPub:  micro.NewPublisher(topic.ConfirmEmailTopic, rpcClient),
		mobilePub: micro.NewPublisher(topic.ConfirmMobileTopic, rpcClient),
	}, nil
}

type Handler struct {
	userCli userPb.UserService
	authCli authPb.AuthService

	emailPub  micro.Publisher
	mobilePub micro.Publisher
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
		log.ErrorF("[api.user.handler::Login] call \"srv.auth::CheckCredential\" failed, %s", err)
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
		log.ErrorF("[api.user.handler::Login] call \"srv.auth::SignOn\" failed", err)
		return
	}
	if !resp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": resp.Error.Code,
			"msg":  resp.Error.Detail,
		})
		return
	}

	c.SetCookie("_token", resp.Token, int(constant.JwtExpiredTime.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"code":  errno.OK.Code,
		"msg":   errno.OK.Message,
		"token": resp.Token,
	})
}

// 注销 token
func (h *Handler) Logout(c *gin.Context) {
	token, err := c.Cookie("_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code": -1,
			"msg":  "You must login first",
		})
		return
	}

	ctx := context.TODO()
	resp, err := h.authCli.SignOff(ctx, &authPb.Request{
		Token: token,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errno.GetRespFromErr(errno.RemoteServiceErr))
		log.ErrorF("[api.user.handler::Logout] call \"srv.auth::SignOff\" failed", err)
		return
	}
	if !resp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": resp.Error.Code,
			"msg":  resp.Error.Detail,
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

// 默认注册使用邮箱进行注册
// 也可使用手机验证码一键登录或三方登录，当手机号未注册过时自动添加用户
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
		log.ErrorF("[api.user.handler::Register] call \"srv.user::CreateUser\" failed, %s", err)
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
		log.ErrorF("[api.user.handler::Register] call \"srv.auth::AddLoginCredential\" failed, %s", err)
		return
	}
	if !authResp.Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": authResp.Error.Code,
			"msg":  authResp.Error.Detail,
		})
		return
	}

	go func() {
		msg := &msgPb.ConfirmEmail{
			Msg: &msgPb.BaseMessage{
				Id:      uuid.New().String(),
				Time:    time.Now().Unix(),
				Message: "",
			},
			Username: form.Name,
			AuthId:   form.Email,
		}
		err = h.mobilePub.Publish(ctx, msg)
		if err != nil {
			log.Error("[api.user.handler::Register] Send Event failed", log.Any{
				"error": err,
				"event": msg,
			})
		}
		log.Info("[api.user.handler::Register] event published", log.Any{
			"event": msg,
		})
	}()

	// TODO 发送验证邮件
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
		log.ErrorF("[api.user.handler::Register] call \"srv.auth::SignOn\" failed", err)
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

type SendSmsCodeForm struct {
	// src: web / ios / etc..
	Src    string `form:"src" binding:"required" valid:""`
	Mobile string `form:"mobile" binding:"required" valid:""`
	// 登录 / 注册 / 绑定
	Action string `form:"action" binding:"required" valid:""`
}

func (h *Handler) SendSmsCode(c *gin.Context) {
	var form SendSmsCodeForm
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "表单字段不可为空",
		})
		return
	}
	// govalidator
	if !utils.ValidateMobile(form.Mobile) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "手机号码不合法",
		})
	}

	msg := &msgPb.ConfirmMobile{
		Msg: &msgPb.BaseMessage{
			Id:      uuid.New().String(),
			Time:    time.Now().Unix(),
			Message: "",
		},
		// todo token rpc
		OneTimeToken: "899998",
		Mobile:       form.Mobile,
	}
	ctx := context.TODO()
	if err := h.mobilePub.Publish(ctx, msg); err != nil {
		log.Error("[api.user.handler::SendSmsCode] publish event failed", log.Any{
			"error": err,
			"event": msg,
		})
	}

	log.Info("[api.user.handler::SendSmsCode] event published", log.Any{
		"event": msg,
	})

	c.JSON(http.StatusOK, errno.GetRespFromErr(errno.OK))
}
