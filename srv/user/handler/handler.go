package handler

import (
	"context"
	"database/sql"
	"github.com/f1renze/the-architect/common/utils/log"
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
		resp.Error = &pb.Error{
			Code: 500,
			Detail: err.Error(),
		}

		if model.IsEmptyUserNameErr(err) {
			resp.Error.Code = 400
		} else {
			log.Error("create user failed", log.Any{
				"error": err,
				"req": req,
			})
		}
		return err
	}

	resp.Success = true
	resp.User = user
	return nil
}

func (h *Handler) QueryUser(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	if req.User == nil {
		resp.Success = false
		resp.Error = &pb.Error{
			Code: 400,
			Detail: "invalid request param",
		}
		return nil
	}
	user, err := h.model.QueryUser(req.User.Id)

	if err != nil {
		resp.Success = false
		resp.Error = &pb.Error{
			Code: 500,
			Detail: err.Error(),
		}
		if err == sql.ErrNoRows {
			resp.Error.Code = 400
		} else {
			log.Error("query user failed", log.Any{
				"error": err,
				"uid": req.User.Id,
			})
		}
		return err
	}

	resp.User = user
	resp.Success = true
	return nil
}