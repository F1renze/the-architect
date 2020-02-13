package test

import (
	"context"
	"github.com/f1renze/the-architect/common"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/infra/db"
	"github.com/f1renze/the-architect/srv/casbin/handler"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	pb "github.com/f1renze/the-architect/srv/casbin/proto"
)

func TestMain(m *testing.M) {
	os.Setenv(constant.CfgCenterAddrEnv, "127.0.0.1:9689")
	common.Init()

	m.Run()
}

func TestRBACModel(t *testing.T) {
	tc := []struct{
		// user
		sub string
		// domain
		dom string
		// resource
		obj string
		// operation
		act string
		expect bool
	}{
		{"bob", "domain2", "data1", "read", false},
		{"alice", "domain1", "data2", "read", false},
		{"alice", "domain1", "data1", "read", true},
		{"bob", "domain2", "data2", "read", true},
		{"user", "root", "data1", "write", true},
		{"user", "root", "data2", "write", true},
	}

	h := handler.NewHandler()
	ctx := context.Background()

	adapterRep := &pb.NewAdapterReply{}
	err := h.NewAdapter(ctx, &pb.NewAdapterRequest{DriverName: "file", ConnectString: "./rbac_policy.csv"}, adapterRep)
	if err != nil {
		t.Fatal(err)
	}

	modelText, err := ioutil.ReadFile("./rbac_model.conf")
	if err != nil {
		t.Fatal(err)
	}

	enforcerRep := &pb.NewEnforcerReply{}
	err = h.NewEnforcer(ctx, &pb.NewEnforcerRequest{ModelText: string(modelText), AdapterHandle: 0}, enforcerRep)
	if err != nil {
		t.Fatal(err)
	}
	e := enforcerRep.Handler

	var emptyRep *pb.EmptyReply
	for i := range tc {
		emptyRep = new(pb.EmptyReply)
		err = h.Enforce(ctx, &pb.EnforceRequest{EnforcerHandler: e, Params: []string{tc[i].sub, tc[i].dom, tc[i].obj, tc[i].act}}, emptyRep)
		if err != nil {
			t.Fatal(err)
		}
		if emptyRep.Success != tc[i].expect {
			t.Fatalf("%d: expect %t, got %t", i, tc[i].expect, emptyRep.Success)
		}
	}

	// p type
	policy := []struct{
		sub string
		dom string
		obj string
		act string
	} {
		{"admin", "domain1", "data1", "read"},
		{"admin", "domain1", "data1", "write"},
		{"admin", "domain2", "data2", "read"},
		{"admin", "domain2", "data2", "write"},
		{"g", "alice", "admin", "domain1"},
	}

	groupingPolicy := []struct{
		user string
		role string
		domain string
	} {
		{"alice", "admin", "domain1"},
		{"bob", "admin", "domain2"},
		{"user", "admin", "root"},
	}

	
	// change adapter
	url := strings.Split(db.GetMysqlUrl(), "/")[0] + "/"
	t.Log(url)
	err = h.NewAdapter(ctx, &pb.NewAdapterRequest{DriverName: "mysql", ConnectString: url}, adapterRep)
	if err != nil {
		t.Fatal(err)
	}
	err = h.NewEnforcer(ctx, &pb.NewEnforcerRequest{ModelText: string(modelText), AdapterHandle: adapterRep.Handler}, enforcerRep)
	e = enforcerRep.Handler

	for _, p := range policy {
		emptyRep = new(pb.EmptyReply)
		err = h.AddPolicy(ctx, &pb.PolicyRequest{EnforcerHandler: e, Params: []string{p.sub, p.dom, p.obj, p.act}}, emptyRep)
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, g := range groupingPolicy {
		emptyRep = new(pb.EmptyReply)
		err = h.AddGroupingPolicy(ctx, &pb.PolicyRequest{EnforcerHandler:e, Params: []string{g.user, g.role, g.domain}}, emptyRep)
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := range tc {
		emptyRep = new(pb.EmptyReply)
		err = h.Enforce(ctx, &pb.EnforceRequest{EnforcerHandler: e, Params: []string{tc[i].sub, tc[i].dom, tc[i].obj, tc[i].act}}, emptyRep)
		if err != nil {
			t.Fatal(err)
		}
		if emptyRep.Success != tc[i].expect {
			t.Fatalf("%d: expect %t, got %t", i, tc[i].expect, emptyRep.Success)
		}
	}
}