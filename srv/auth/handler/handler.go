package handler

import (
	"context"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/srv/auth/model"
	pb "github.com/f1renze/the-architect/srv/auth/proto"
)

func NewHandler() pb.AuthHandler {
	return &Handler{
		model: model.NewModel(),
	}
}

type Handler struct {
	model model.AuthModel
}

func (h *Handler) AddLoginCredential(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	_, err := h.model.CreateCredential(req.Info)
	if err != nil {
		resp.Success = false
		resp.Error = new(pb.Error)
		resp.Error.Code, resp.Error.Detail = errno.DecodeInt32Err(err)
		return nil
	}

	resp.Success = true
	return nil
}

func (h *Handler) CheckCredential(ctx context.Context, req *pb.Request, resp *pb.Response) error {

	authInfo, err := h.model.QueryCredential(req.Info.AuthType, req.Info.AuthId, req.Info.Credential)
	if err != nil {
		resp.Success = false
		resp.Error = new(pb.Error)
		resp.Error.Code, resp.Error.Detail = errno.DecodeInt32Err(err)
		return nil
	}
	// todo verify 留 api 层处理

	if req.Login {
		_ = h.model.RefreshLoginTime(authInfo.Id, req.Info.IpAddr)
	}

	return nil
}
