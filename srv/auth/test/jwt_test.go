package test

import (
	"github.com/f1renze/the-architect/srv/auth/model"
	"testing"
	"time"
)

func TestJwt_CreateAndExpire(t *testing.T) {
	token, err := jwtSrv.NewJwt(1, "example@example.com")
	if err != nil {
		t.Fatal("create token failed", err)
	}

	payload, err := jwtSrv.ParseJwt(token)
	if err != nil {
		t.Fatal("parse token error", err)
	}
	ok, err := jwtSrv.IsValid(payload.UId, payload.Id)
	if err != nil || !ok {
		t.Fatal("token invalid", err)
	}
	err = jwtSrv.ExpireJwt(payload.UId)
	if err != nil {
		t.Fatal("sign off failed", err)
	}
	ok, err = jwtSrv.IsValid(payload.UId, payload.Id)
	if ok {
		t.Fatal("sign off failed, token still valid", err)
	}
}

func TestJwt_Refresh(t *testing.T) {
	model.JwtExpiredTime = 5 * time.Second
	token, err := jwtSrv.NewJwt(1, "example@example.com")
	if err != nil {
		t.Fatal("create token failed", err)
	}
	payload, err := jwtSrv.ParseJwt(token)
	if err != nil {
		t.Fatal("parse token error", err)
	}

	time.Sleep(3 * time.Second)
	token2, err := jwtSrv.NewJwt(1, "example@example.com")
	if err != nil {
		t.Fatal("create token failed", err)
	}
	payload2, err := jwtSrv.ParseJwt(token2)
	if err != nil {
		t.Fatal("parse token error", err)
	}

	ok, err := jwtSrv.IsValid(payload.UId, payload.Id)
	if err != nil || !ok {
		t.Fatal("token invalid", err)
	}
	ok, err = jwtSrv.IsValid(payload2.UId, payload2.Id)
	if err != nil || !ok {
		t.Fatal("token invalid", err)
	}

	err = jwtSrv.ExpireJwt(payload.UId)
	if err != nil {
		t.Fatal("sign off failed", err)
	}
	ok, err = jwtSrv.IsValid(payload.UId, payload.Id)
	if ok {
		t.Fatal("sign off failed, token still valid", err)
	}
	ok, err = jwtSrv.IsValid(payload2.UId, payload2.Id)
	if ok {
		t.Fatal("sign off failed, token still valid", err)
	}
}
