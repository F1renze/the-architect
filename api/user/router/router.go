package router

import (
	"github.com/f1renze/the-architect/api/user/handler"
	"github.com/gin-gonic/gin"
)

func Default(h handler.UserApi) *gin.Engine {
	r := gin.Default()

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)

	r.POST("/send-sms", h.SendSmsCode)

	return r
}
