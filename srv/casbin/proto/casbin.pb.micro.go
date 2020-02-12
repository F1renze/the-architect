// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/casbin/proto/casbin.proto

package micro_arch_srv_casbin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Casbin service

type CasbinService interface {
	NewEnforcer(ctx context.Context, in *NewEnforcerRequest, opts ...client.CallOption) (*NewEnforcerReply, error)
	NewAdapter(ctx context.Context, in *NewAdapterRequest, opts ...client.CallOption) (*NewAdapterReply, error)
	Enforce(ctx context.Context, in *EnforceRequest, opts ...client.CallOption) (*EmptyReply, error)
	LoadPolicy(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyReply, error)
	SavePolicy(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyReply, error)
	GetRolesForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*ArrayReply, error)
	GetImplicitRolesForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*ArrayReply, error)
	GetUsersForRole(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*ArrayReply, error)
	HasRoleForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error)
	AddRoleForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error)
	// DeleteRoleForUser deletes a role for a user.
	// Returns false if the user does not have the role (aka not affected).
	DeleteRoleForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error)
	// DeleteRolesForUser deletes all roles for a user.
	// Returns false if the user does not have any roles (aka not affected).
	DeleteRolesForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error)
	DeleteRole(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error)
}

type casbinService struct {
	c    client.Client
	name string
}

func NewCasbinService(name string, c client.Client) CasbinService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "micro.arch.srv.casbin"
	}
	return &casbinService{
		c:    c,
		name: name,
	}
}

func (c *casbinService) NewEnforcer(ctx context.Context, in *NewEnforcerRequest, opts ...client.CallOption) (*NewEnforcerReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.NewEnforcer", in)
	out := new(NewEnforcerReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) NewAdapter(ctx context.Context, in *NewAdapterRequest, opts ...client.CallOption) (*NewAdapterReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.NewAdapter", in)
	out := new(NewAdapterReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) Enforce(ctx context.Context, in *EnforceRequest, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.Enforce", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) LoadPolicy(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.LoadPolicy", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) SavePolicy(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.SavePolicy", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) GetRolesForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*ArrayReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.GetRolesForUser", in)
	out := new(ArrayReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) GetImplicitRolesForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*ArrayReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.GetImplicitRolesForUser", in)
	out := new(ArrayReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) GetUsersForRole(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*ArrayReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.GetUsersForRole", in)
	out := new(ArrayReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) HasRoleForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.HasRoleForUser", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) AddRoleForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.AddRoleForUser", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) DeleteRoleForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.DeleteRoleForUser", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) DeleteRolesForUser(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.DeleteRolesForUser", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) DeleteRole(ctx context.Context, in *UserRoleRequest, opts ...client.CallOption) (*EmptyReply, error) {
	req := c.c.NewRequest(c.name, "Casbin.DeleteRole", in)
	out := new(EmptyReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Casbin service

type CasbinHandler interface {
	NewEnforcer(context.Context, *NewEnforcerRequest, *NewEnforcerReply) error
	NewAdapter(context.Context, *NewAdapterRequest, *NewAdapterReply) error
	Enforce(context.Context, *EnforceRequest, *EmptyReply) error
	LoadPolicy(context.Context, *EmptyRequest, *EmptyReply) error
	SavePolicy(context.Context, *EmptyRequest, *EmptyReply) error
	GetRolesForUser(context.Context, *UserRoleRequest, *ArrayReply) error
	GetImplicitRolesForUser(context.Context, *UserRoleRequest, *ArrayReply) error
	GetUsersForRole(context.Context, *UserRoleRequest, *ArrayReply) error
	HasRoleForUser(context.Context, *UserRoleRequest, *EmptyReply) error
	AddRoleForUser(context.Context, *UserRoleRequest, *EmptyReply) error
	// DeleteRoleForUser deletes a role for a user.
	// Returns false if the user does not have the role (aka not affected).
	DeleteRoleForUser(context.Context, *UserRoleRequest, *EmptyReply) error
	// DeleteRolesForUser deletes all roles for a user.
	// Returns false if the user does not have any roles (aka not affected).
	DeleteRolesForUser(context.Context, *UserRoleRequest, *EmptyReply) error
	DeleteRole(context.Context, *UserRoleRequest, *EmptyReply) error
}

func RegisterCasbinHandler(s server.Server, hdlr CasbinHandler, opts ...server.HandlerOption) error {
	type casbin interface {
		NewEnforcer(ctx context.Context, in *NewEnforcerRequest, out *NewEnforcerReply) error
		NewAdapter(ctx context.Context, in *NewAdapterRequest, out *NewAdapterReply) error
		Enforce(ctx context.Context, in *EnforceRequest, out *EmptyReply) error
		LoadPolicy(ctx context.Context, in *EmptyRequest, out *EmptyReply) error
		SavePolicy(ctx context.Context, in *EmptyRequest, out *EmptyReply) error
		GetRolesForUser(ctx context.Context, in *UserRoleRequest, out *ArrayReply) error
		GetImplicitRolesForUser(ctx context.Context, in *UserRoleRequest, out *ArrayReply) error
		GetUsersForRole(ctx context.Context, in *UserRoleRequest, out *ArrayReply) error
		HasRoleForUser(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error
		AddRoleForUser(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error
		DeleteRoleForUser(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error
		DeleteRolesForUser(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error
		DeleteRole(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error
	}
	type Casbin struct {
		casbin
	}
	h := &casbinHandler{hdlr}
	return s.Handle(s.NewHandler(&Casbin{h}, opts...))
}

type casbinHandler struct {
	CasbinHandler
}

func (h *casbinHandler) NewEnforcer(ctx context.Context, in *NewEnforcerRequest, out *NewEnforcerReply) error {
	return h.CasbinHandler.NewEnforcer(ctx, in, out)
}

func (h *casbinHandler) NewAdapter(ctx context.Context, in *NewAdapterRequest, out *NewAdapterReply) error {
	return h.CasbinHandler.NewAdapter(ctx, in, out)
}

func (h *casbinHandler) Enforce(ctx context.Context, in *EnforceRequest, out *EmptyReply) error {
	return h.CasbinHandler.Enforce(ctx, in, out)
}

func (h *casbinHandler) LoadPolicy(ctx context.Context, in *EmptyRequest, out *EmptyReply) error {
	return h.CasbinHandler.LoadPolicy(ctx, in, out)
}

func (h *casbinHandler) SavePolicy(ctx context.Context, in *EmptyRequest, out *EmptyReply) error {
	return h.CasbinHandler.SavePolicy(ctx, in, out)
}

func (h *casbinHandler) GetRolesForUser(ctx context.Context, in *UserRoleRequest, out *ArrayReply) error {
	return h.CasbinHandler.GetRolesForUser(ctx, in, out)
}

func (h *casbinHandler) GetImplicitRolesForUser(ctx context.Context, in *UserRoleRequest, out *ArrayReply) error {
	return h.CasbinHandler.GetImplicitRolesForUser(ctx, in, out)
}

func (h *casbinHandler) GetUsersForRole(ctx context.Context, in *UserRoleRequest, out *ArrayReply) error {
	return h.CasbinHandler.GetUsersForRole(ctx, in, out)
}

func (h *casbinHandler) HasRoleForUser(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error {
	return h.CasbinHandler.HasRoleForUser(ctx, in, out)
}

func (h *casbinHandler) AddRoleForUser(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error {
	return h.CasbinHandler.AddRoleForUser(ctx, in, out)
}

func (h *casbinHandler) DeleteRoleForUser(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error {
	return h.CasbinHandler.DeleteRoleForUser(ctx, in, out)
}

func (h *casbinHandler) DeleteRolesForUser(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error {
	return h.CasbinHandler.DeleteRolesForUser(ctx, in, out)
}

func (h *casbinHandler) DeleteRole(ctx context.Context, in *UserRoleRequest, out *EmptyReply) error {
	return h.CasbinHandler.DeleteRole(ctx, in, out)
}
