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
		jwt: model.NewJwtSrv(),
	}
}

type Handler struct {
	model model.AuthModel
	jwt model.JwtSrv
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

	authInfo, err := h.model.QueryCredential(req.Info.AuthId, req.Info.Credential)
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
	resp.Success = true
	resp.Info = authInfo
	return nil
}

func (h *Handler) VerifyToken(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	payload, err := h.jwt.ParseJwt(req.Token)
	if err == nil {
		ok, err := h.jwt.IsValid(payload.UId, payload.Id)
		if err == nil && ok {
			resp.Success = true
			resp.Valid = true
			resp.Token = req.Token
			return nil
		}
	}

	resp.Success = false
	resp.Error = new(pb.Error)
	resp.Error.Code, resp.Error.Detail = errno.DecodeInt32Err(err)
	return nil
}

func (h *Handler) SignOn(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	token, err := h.jwt.NewJwt(req.Info.Uid, req.Info.AuthId)
	if err != nil {
		resp.Success = false
		resp.Error = new(pb.Error)
		resp.Error.Code, resp.Error.Detail = errno.DecodeInt32Err(err)
		return nil
	}
	resp.Success = true
	resp.Token = token
	resp.Valid = true
	return nil
}

func (h *Handler) SignOff(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	payload, err := h.jwt.ParseJwt(req.Token)
	if err == nil {
		ok, err := h.jwt.IsValid(payload.UId, payload.Id)
		if err == nil && ok {
			if err = h.jwt.ExpireJwt(payload.UId); err == nil {
				resp.Success = true
				return nil
			}
		}
	}

	resp.Success = false
	resp.Error = new(pb.Error)
	resp.Error.Code, resp.Error.Detail = errno.DecodeInt32Err(err)
	return nil
}
