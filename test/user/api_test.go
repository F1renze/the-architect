package user

import (
	"encoding/json"
	"fmt"
	"github.com/f1renze/the-architect/api/user/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/f1renze/the-architect/api/user/handler"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/test"
	"github.com/gin-gonic/gin"
)

func TestRegisterApi(t *testing.T) {
	tc := []struct {
		Username       string
		Email          string
		Pwd            string
		ConfirmPwd     string
		ExceptHttpCode int
		ExceptCode     int
	}{
		// test form validation
		{
			"",
			"",
			"",
			"",
			400,
			-1,
		},
		{
			"test1",
			"123345345",
			"test123",
			"test1234",
			400,
			-1,
		},
		{
			"test1",
			"123345345",
			"test123",
			"test123",
			400,
			-1,
		},
		// test normal register
		{
			"test1",
			"Eugene@gmail.com",
			"test123",
			"test123",
			200,
			0,
		},
		// test repeated email
		{
			"test2",
			"Eugene@gmail.com",
			"test123",
			"test123",
			200,
			errno.AuthIdAlreadyUsed.Code,
		},
	}

	h, err := handler.NewHandler(cmsCli)
	if err != nil {
		t.Fatal("init handler failed", err)
	}
	router := gin.Default()
	router.POST("/register", h.Register)

	for i, tCase := range tc {
		body := fmt.Sprintf("username=%s&email=%s&password=%s&confirm_password=%s",
			tCase.Username, tCase.Email, tCase.Pwd, tCase.ConfirmPwd)
		r := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(body))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		t.Logf("%d, resp: %v", i, w.Body)
		test.AssertEqual(t, tCase.ExceptHttpCode, w.Code)

		var resp map[string]interface{}
		err = json.NewDecoder(w.Body).Decode(&resp)

		code, ok := resp["code"].(float64)
		if ok {
			test.AssertEqual(t, tCase.ExceptCode, int(code))
		}

	}
	// post json:
	//mBody := map[string]interface{}{
	//	"username":         tCase.Username,
	//	"email":            tCase.Email,
	//	"password":         tCase.Pwd,
	//	"confirm_password": tCase.ConfirmPwd,
	//}
	//body, _ := json.Marshal(mBody)
	// bytes.NewReader
}

func TestSendSms(t *testing.T) {
	h, err := handler.NewHandler(cmsCli)
	if err != nil {
		t.Fatal("init handler failed", err)
	}
	r := router.Default(h)

	body := fmt.Sprintf("src=%s&mobile=%s&action=%s", "web", "13924776036", "register")

	req := httptest.NewRequest(http.MethodPost, "/send-sms", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	t.Logf("resp: %v", resp.Body)

}