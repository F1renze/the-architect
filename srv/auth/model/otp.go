package model

import (
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/errno"
	redis2 "github.com/f1renze/the-architect/common/infra/redis"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/go-redis/redis"
	"github.com/pquerna/otp/totp"
	"time"
)

type OtpSrv interface {
	GenerateCode(authId string) (code string, err error)
	ValidateCode(authId string, code string) (bool, error)
}

func NewOtpSrv() OtpSrv {
	return &otpSrv{
		rCli: redis2.GetClient(),
	}
}

type otpSrv struct {
	rCli *redis.Client
}

func (s *otpSrv)GenerateCode(authId string) (code string, err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      constant.TokenIssuer,
		AccountName: authId,
	})
	if err != nil {
		err = errno.SystemErr.With(err)
		log.Error("[srv.auth.model::GenerateCode] totp generate key error", log.Any{
			"error": err, "auth_id": authId,
		})
		return
	}

	code, err = totp.GenerateCode(key.Secret(), time.Now().UTC())
	if err != nil {
		err = errno.SystemErr.With(err)
		log.Error("[srv.auth.model::GenerateCode] totp generate code error", log.Any{
			"error": err, "auth_id": authId,
		})
		return
	}

	err = s.rCli.Set(s.getKeyFromAuthId(authId) , code, constant.SmsCodeExpiredTime).Err()
	if err != nil {
		log.ErrorF("[srv.auth.model::GenerateCode] redis opeation err: %s", err)
		return "", errno.RedisOperationErr.With(err)
	}

	return
}

func (s *otpSrv) ValidateCode(authId string, code string) (bool, error) {
	secret, err := s.rCli.Get(s.getKeyFromAuthId(authId)).Result()
	if err == redis.Nil {
		return false, errno.TokenExpired.With(err)
	} else if err != nil {
		return false, errno.RedisOperationErr.With(err)
	}

	return totp.Validate(code, secret), nil
}

func (s *otpSrv) getKeyFromAuthId(authId string) string {
	return constant.SmsCodeKeyPrefix + authId
}
