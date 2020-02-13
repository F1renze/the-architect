package handler

import (
	"context"

	pb "github.com/f1renze/the-architect/srv/casbin/proto"
)

func (h *Handler) AddPolicy(ctx context.Context, in *pb.PolicyRequest, out *pb.EmptyReply) error {
	in.PType = "p"
	return h.AddNamedPolicy(ctx, in, out)
}

func (h *Handler) AddNamedPolicy(ctx context.Context, in *pb.PolicyRequest, out *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		out.Success = false
		return err
	}

	ok, err := e.AddNamedPolicy(in.PType, in.Params)
	out.Success = ok
	return err
}

func (h *Handler) AddGroupingPolicy(ctx context.Context, in *pb.PolicyRequest, out *pb.EmptyReply) error {
	in.PType = "g"
	return h.AddNamedGroupingPolicy(ctx, in, out)
}

func (h *Handler) AddNamedGroupingPolicy(ctx context.Context, in *pb.PolicyRequest, out *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		out.Success = false
		return err
	}

	ok, err := e.AddNamedGroupingPolicy(in.PType, in.Params)
	out.Success = ok
	return err
}
