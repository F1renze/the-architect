package handler

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/srv/casbin/adapter"
	pb "github.com/f1renze/the-architect/srv/casbin/proto"
)

func NewHandler() pb.CasbinHandler {
	return &Handler{
		enforcerMap: make(map[int]*casbin.Enforcer),
		adapterMap:  make(map[int]persist.Adapter),
	}
}

type Handler struct {
	enforcerMap map[int]*casbin.Enforcer
	adapterMap  map[int]persist.Adapter
}

func (h *Handler) getEnforcer(handler int) (*casbin.Enforcer, error) {
	if _, ok := h.enforcerMap[handler]; ok {
		return h.enforcerMap[handler], nil
	} else {
		return nil, errno.EnforcerNotFound
	}
}

func (h *Handler) getAdapter(handle int) (persist.Adapter, error) {
	if _, ok := h.adapterMap[handle]; ok {
		return h.adapterMap[handle], nil
	} else {
		return nil, errno.AdapterNotFound
	}
}

func (h *Handler) addEnforcer(e *casbin.Enforcer) int {
	cnt := len(h.enforcerMap)
	h.enforcerMap[cnt] = e
	return cnt
}

func (h *Handler) addAdapter(a persist.Adapter) int {
	cnt := len(h.adapterMap)
	h.adapterMap[cnt] = a
	return cnt
}

func (h *Handler) NewEnforcer(ctx context.Context, req *pb.NewEnforcerRequest, resp *pb.NewEnforcerReply) error {
	var a persist.Adapter
	var e *casbin.Enforcer

	if req.AdapterHandle != -1 {
		var err error
		a, err = h.getAdapter(int(req.AdapterHandle))
		if err != nil {
			resp.Handler = -1
			return err
		}
	}

	if req.ModelText == "" {
		resp.Handler = -1
		return errno.InvalidModel.Add("model must not empty")
	}

	if a == nil {
		m, err := model.NewModelFromString(req.ModelText)
		if err != nil {
			resp.Handler = -1
			return errno.InvalidModel.With(err)
		}

		e, err = casbin.NewEnforcer(m)
		if err != nil {
			resp.Handler = -1
			return errno.InvalidModel.With(err)
		}
	} else {
		m, err := model.NewModelFromString(req.ModelText)
		if err != nil {
			resp.Handler = -1
			return errno.InvalidModel.With(err)
		}

		e, err = casbin.NewEnforcer(m, a)
		if err != nil {
			resp.Handler = -1
			return errno.InvalidModel.With(err)
		}
	}
	n := h.addEnforcer(e)

	resp.Handler = int32(n)
	return nil
}

func (h *Handler) NewAdapter(ctx context.Context, req *pb.NewAdapterRequest, resp *pb.NewAdapterReply) error {
	a, err := adapter.NewAdapter(req)
	if err != nil {
		resp.Handler = -1
		return err
	}

	n := h.addAdapter(a)
	resp.Handler = int32(n)
	return nil
}

func (h *Handler) Enforce(ctx context.Context, req *pb.EnforceRequest, resp *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(req.EnforcerHandler))
	if err != nil {
		resp.Success = false
		return err
	}
	var param interface{}
	params := make([]interface{}, 0, len(req.Params))
	m := e.GetModel()["m"]["m"].Value

	for index := range req.Params {
		param, m = parseParam(req.Params[index], m)
		params = append(params, param)
	}
	//e.EnableLog(true)

	res, err := e.EnforceWithMatcher(m, params...)
	resp.Success = res
	return err
}

// todo complete error code
func (h *Handler) LoadPolicy(ctx context.Context, req *pb.EmptyRequest, resp *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(req.Handler))
	if err != nil {
		resp = &pb.EmptyReply{Success: false}
		return err
	}

	err = e.LoadPolicy()
	resp = &pb.EmptyReply{Success: true}
	if err != nil {
		resp.Success = false
	}

	return err
}

func (h *Handler) SavePolicy(ctx context.Context, req *pb.EmptyRequest, resp *pb.EmptyReply) error {
	e, err := h.getEnforcer(int(req.Handler))
	if err != nil {
		resp = &pb.EmptyReply{Success: false}
		return err
	}

	err = e.SavePolicy()
	resp = &pb.EmptyReply{Success: true}
	if err != nil {
		resp.Success = false
	}
	return err
}
