package test

import (
	"github.com/f1renze/the-architect/test"
	"testing"
)

func TestModel_CreateUser(t *testing.T) {
	//name := strconv.FormatInt(time.Now().Unix(), 10)
	name := "1580816469"
	u, err := userModel.CreateUser(name, "")
	if err != nil {
		t.Fatal("create user failed:", err)
	}

	test.AssertEqual(t, name, u.Name)
}

func TestModel_QueryUser(t *testing.T) {

}
