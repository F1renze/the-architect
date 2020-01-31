package model

import "errors"

var (
	EmptyUserNameErr = errors.New("invalid param: name must not empty")
)

func IsEmptyUserNameErr(err error) bool {
	return err == EmptyUserNameErr
}

