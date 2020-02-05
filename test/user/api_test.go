package user

import (
	"encoding/json"
	"fmt"
	"github.com/f1renze/the-architect/api/user/handler"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/test"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
		{
			"test1",
			"123345345",
			"test123",
			"test123",
			400,
			errno.InvalidEmail.Code,
		},
		{
			"test1",
			"Eugene@gmail.com",
			"test123",
			"test123",
			200,
			0,
		},
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
