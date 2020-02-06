package test

import (
	"database/sql"
	"fmt"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/test"
	"testing"

	pb "github.com/f1renze/the-architect/srv/auth/proto"
)

func TestAuthModel(t *testing.T) {
	tc := []struct {
		pb.AuthInfo
		err error
	}{
		{
			pb.AuthInfo{
				Uid:        1,
				AuthType:   pb.AuthType_Mobile,
				AuthId:     "13926799139",
				Credential: "test123",
			},
			errno.AuthTypeAlreadyBind,
		},
		{
			pb.AuthInfo{
				Uid:        2,
				AuthType:   pb.AuthType_Mobile,
				AuthId:     "13926799139",
				Credential: "test123",
			},
			errno.AuthIdAlreadyUsed,
		},
	}

	var (
		err  error
		info *pb.AuthInfo
	)

	for i := range tc {
		t.Logf("test case %d", i)
		_, err = authModel.CreateCredential(&tc[i].AuthInfo)
		if err != nil && err != tc[i].err {
			t.Fatal("create credential failed:", err)
		} else if err != nil {
			continue
		}

		info, err = authModel.QueryCredential(tc[i].AuthType, tc[i].AuthId, tc[i].Credential)
		if err != nil {
			t.Fatal("query failed", err, tc[i], err == sql.ErrNoRows)
		}
		// mysql tinyint(1) 直接转换为 bool
		fmt.Println("verified", info.Verified)
		test.AssertEqual(t, tc[i].Uid, info.Uid)
		err = authModel.RefreshLoginTime(info.Id, "127.0.0.1")
		if err != nil {
			t.Fatal("refresh login time failed")
		}
	}
}
