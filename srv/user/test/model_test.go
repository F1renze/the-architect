package test

import (
	"strconv"
	"testing"
	"time"

	"github.com/f1renze/the-architect/test"
)

func TestModel_CreateUser(t *testing.T) {
	name := strconv.FormatInt(time.Now().Unix(), 10)
	u, err := userModel.CreateUser(name, "")
	if err != nil {
		t.Fatal("create user failed")
	}

	test.AssertEqual(t, name, u.Name)
}

func TestModel_QueryUser(t *testing.T) {

}
