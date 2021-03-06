package errno

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Errno struct {
	Code    int
	Message string
}

func (e *Errno) Error() string {
	return e.Message
}

func (e *Errno) Add(format string, v ...interface{}) error {
	e.Message += ", " + fmt.Sprintf(format, v...)
	return e
}

func (e *Errno) With(err error) error {
	return &Err{e, err}
}

func newCode(code int, msg string) *Errno {
	return &Errno{
		Code:    code,
		Message: msg,
	}
}

type Err struct {
	*Errno
	err error
}

func (e *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", e.Code, e.Message, e.err)
}

func DecodeInt32Err(err error) (int32, string) {
	if err == nil {
		return int32(OK.Code), OK.Message
	}

	// cannot fallthrough in type switch
	switch t := err.(type) {
	case *Errno:
		return int32(t.Code), t.Message
	case *Err:
		return int32(t.Code), t.Message
	default:
	}

	return int32(SystemErr.Code), SystemErr.Message
}

func GetRespFromErr(err error) gin.H {
	code, msg := DecodeInt32Err(err)
	return gin.H{
		"code": code,
		"msg": msg,
	}
}
