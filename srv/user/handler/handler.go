package handler

import (
	"context"

	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/srv/user/model"
	pb "github.com/f1renze/the-architect/srv/user/proto"
)

func NewHandler() pb.UserHandler {
	return &Handler{
		model: model.NewModel(),
	}
}

type Handler struct {
	model model.UserModel
}

func (h *Handler) CreateUser(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	user, err := h.model.CreateUser(req.User.Name, req.User.Avatar)

	if err != nil {
		resp.Success = false
		resp.Error = new(pb.Error)
		resp.Error.Code, resp.Error.Detail = errno.DecodeInt32Err(err)
		return nil
	}

	resp.Success = true
	resp.User = user
	return nil
}

func (h *Handler) QueryUser(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	if req.User == nil {
		resp.Success = false
		resp.Error = new(pb.Error)
		resp.Error.Code, resp.Error.Detail = errno.DecodeInt32Err(errno.UidIsNull)
		return nil
	}
	user, err := h.model.QueryUser(req.User.Id)

	if err != nil {
		resp.Success = false
		resp.Error = new(pb.Error)
		resp.Error.Code, resp.Error.Detail = errno.DecodeInt32Err(err)
		return nil
	}

	resp.User = user
	resp.Success = true
	return nil
}
