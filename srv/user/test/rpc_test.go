package test

import (
	"context"
	"github.com/f1renze/the-architect/test"
	"strconv"
	"testing"
	"time"

	pb "github.com/f1renze/the-architect/srv/user/proto"
)

func TestSrv(t *testing.T) {
	ctx := context.TODO()
	name := strconv.FormatInt(time.Now().Unix(), 10)
	resp, err := userCli.CreateUser(ctx, &pb.Request{
		User: &pb.User{
			Name: name,
		},
	})
	if err != nil {
		t.Fatal("create user by rpc call failed", err)
	}
	test.Equals(t, true, resp.Success)

	uid := resp.User.Id
	resp, err = userCli.QueryUser(context.TODO(), &pb.Request{
		User: &pb.User{
			Id: uid,
		},
	})
	if err != nil {
		t.Fatal("query user by rpc call failed", err)
	}
	test.Equals(t, true, resp.Success)
}
