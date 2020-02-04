package test

import (
	"fmt"
	"github.com/f1renze/the-architect/test"
	"testing"

	pb "github.com/f1renze/the-architect/srv/auth/proto"
)

func TestAuthModel(t *testing.T) {
	tc := []*pb.AuthInfo{
		{
			Uid:        1,
			AuthType:   pb.AuthType_Mobile,
			AuthId:     "13926999139",
			Credential: "test123",
		},
	}

	var (
		err  error
		info *pb.AuthInfo
	)

	for i := range tc {
		_, err = authModel.CreateCredential(tc[i])
		if err != nil {
			fmt.Println(err)
			//t.Fatal("create credential failed")
		}
		info, err = authModel.QueryCredential(tc[i].AuthType, tc[i].AuthId, tc[i].Credential)
		if err != nil {
			fmt.Println(err)
			t.Fatal("query failed")
		}
		test.AssertEqual(t, tc[i].Uid, info.Uid)
	}
}

