package handler

import (
	"context"
	"io/ioutil"
	"testing"

	pb "github.com/f1renze/the-architect/srv/casbin/proto"
)

func TestRBACModel(t *testing.T) {
	h := NewHandler()
	ctx := context.Background()

	resp := &pb.NewAdapterReply{}
	err := h.NewAdapter(ctx, &pb.NewAdapterRequest{DriverName: "file", ConnectString: "../example/rbac_policy.csv"}, resp)
	if err != nil {
		t.Error(err)
		return
	}

	modelText, err := ioutil.ReadFile("../conf/rbac_model.conf")
	if err != nil {
		t.Error(err)
		return
	}

	resp2 := &pb.NewEnforcerReply{}
	err = h.NewEnforcer(ctx, &pb.NewEnforcerRequest{ModelText: string(modelText), AdapterHandle: 0}, resp2)
	if err != nil {
		t.Error(err)
		return
	}
	e := resp2.Handler

	sub := "alice"
	obj := "data1"
	act := "read"
	dom := "domain1"

	resp3 := &pb.EmptyReply{}
	err = h.Enforce(ctx, &pb.EnforceRequest{EnforcerHandler: e, Params: []string{sub, obj, act, dom}}, resp3)
	if err != nil {
		t.Error(err)
		return
	}

	if resp3.Success != true {
		t.Errorf("%s, %s, %s: %t, supposed to be %t", sub, obj, act, resp3.Success, true)
		return
	}
}