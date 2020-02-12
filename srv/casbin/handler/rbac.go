package handler

import (
	"context"

	pb "github.com/f1renze/the-architect/srv/casbin/proto"
	)

func (h *Handler) GetRolesForUser(ctx context.Context, req *pb.UserRoleRequest, resp *pb.ArrayReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp = &pb.ArrayReply{}
		return err
	}

	res, _ := e.GetModel()["g"]["g"].RM.GetRoles(req.User)
	resp = &pb.ArrayReply{
		Array: res,
	}
	return nil
}

func (h *Handler) GetImplicitRolesForUser(ctx context.Context, req *pb.UserRoleRequest, resp *pb.ArrayReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp = &pb.ArrayReply{}
		return err
	}
	res, err := e.GetImplicitRolesForUser(req.User)
	resp = &pb.ArrayReply{
		Array: res,
	}
	return nil
}

func (h *Handler) GetUsersForRole(ctx context.Context, req *pb.UserRoleRequest, resp *pb.ArrayReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp = &pb.ArrayReply{}
		return err
	}

	res, _ := e.GetModel()["g"]["g"].RM.GetUsers(req.User)

	resp = &pb.ArrayReply{
		Array: res,
	}
	return nil
}

func (h *Handler) HasRoleForUser(ctx context.Context, req *pb.UserRoleRequest, resp *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp = &pb.EmptyReply{}
		return err
	}

	roles, err := e.GetRolesForUser(req.User)
	if err != nil {
		resp = &pb.EmptyReply{}
		return err
	}

	for _, r := range roles {
		if r == req.Role {
			resp = &pb.EmptyReply{Success:true}
			return nil
		}
	}

	resp = &pb.EmptyReply{Success:false}
	return nil
}

func (h *Handler) AddRoleForUser(ctx context.Context, req *pb.UserRoleRequest, resp *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp = &pb.EmptyReply{}
		return err
	}

	ruleAdded, err := e.AddGroupingPolicy(req.User, req.Role)
	resp = &pb.EmptyReply{Success:ruleAdded}
	return err
}

func (h *Handler) DeleteRoleForUser(ctx context.Context, req *pb.UserRoleRequest, resp *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp = &pb.EmptyReply{}
		return err
	}

	ruleRemoved, err := e.RemoveGroupingPolicy(req.User, req.Role)
	resp = &pb.EmptyReply{Success:ruleRemoved}
	return err
}
func (h *Handler) DeleteRolesForUser(ctx context.Context, req *pb.UserRoleRequest, resp *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp = &pb.EmptyReply{}
		return err
	}

	ruleRemoved, err := e.RemoveFilteredGroupingPolicy(0, req.User)
	resp = &pb.EmptyReply{Success:ruleRemoved}
	return err
}
func (h *Handler) DeleteRole(ctx context.Context, req *pb.UserRoleRequest, resp *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp = &pb.EmptyReply{}
		return err
	}

	ok, err := e.DeleteRole(req.Role)
	resp = &pb.EmptyReply{Success:ok}
	return err
}