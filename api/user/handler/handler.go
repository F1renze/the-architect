package handler

import (
	"github.com/gin-gonic/gin"

	userPb "github.com/f1renze/the-architect/srv/user/proto"
)

type UserApi interface {
	Login(*gin.Context)
	Logout(*gin.Context)
	Register(*gin.Context)
}

type Handler struct {
	userCli userPb.UserService
}

func (h *Handler) Login(c *gin.Context) {

}

func (h *Handler) Logout(c *gin.Context) {}

func (h *Handler) Register(c *gin.Context) {

}
