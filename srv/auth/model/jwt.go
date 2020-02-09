package model

import (
	"github.com/f1renze/the-architect/common/constant"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/f1renze/the-architect/common/errno"
	redis2 "github.com/f1renze/the-architect/common/infra/redis"
	"github.com/f1renze/the-architect/common/utils/log"
)

type JwtSrv interface {
	NewJwt(uid uint32, authId string) (encoded string, err error)
	ParseJwt(encoded string) (*JwtPayload, error)
	IsValid(uid uint32, jwtId string) (bool, error)
	ExpireJwt(uid uint32) error
}

var (
	JwtExpiredTime = 24 * time.Hour
	JwtIssuer      = "the-architect"
)

type JwtPayload struct {
	jwt.StandardClaims

	UId    uint32 `json:"uid"`
	AuthId string `json:"auth_id"`
}

func NewJwtSrv() JwtSrv {
	return &jwtSrv{
		rCli: redis2.GetClient(),
	}
}

type jwtSrv struct {
	rCli   *redis.Client
	secret string
}

func (s *jwtSrv) getKeyFromUid(uid uint32) string {
	return constant.JwtKeyPrefix + strconv.FormatUint(uint64(uid), 10)
}

func (s *jwtSrv) IsValid(uid uint32, jwtId string) (bool, error) {
	max := strconv.FormatInt(time.Now().Add(JwtExpiredTime).Unix(), 10)
	min := strconv.FormatInt(time.Now().Add(-JwtExpiredTime).Unix(), 10)

	r, err := s.rCli.ZRangeByScore(s.getKeyFromUid(uid), redis.ZRangeBy{
		Max: max, Min: min,
	}).Result()
	if err != nil {
		return false, errno.RedisOperationErr.With(err)
	}

	for i := range r {
		if jwtId == r[i] {
			return true, nil
		}
	}
	return false, errno.TokenExpired
}

func (s *jwtSrv) ExpireJwt(uid uint32) error {
	key := s.getKeyFromUid(uid)
	err := s.rCli.Del(key).Err()
	// todo test nil key
	if err != nil {
		log.Error("srv.auth.jwt: del key error", log.Any{
			"error": err, "key": key,
		})
		return errno.RedisOperationErr.With(err)
	}
	return nil
}

func (s *jwtSrv) ParseJwt(encoded string) (*JwtPayload, error) {
	token, err := jwt.ParseWithClaims(encoded, &JwtPayload{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errno.SignatureMethodRejected.Add("signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, errno.InvalidToken.With(err)
	}

	if claim, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return claim, nil
	}

	return nil, errno.InvalidToken
}

func (s *jwtSrv) NewJwt(uid uint32, authId string) (encoded string, err error) {
	now := time.Now()
	expiredTime := now.Add(JwtExpiredTime)
	claim := JwtPayload{
		UId:    uid,
		AuthId: authId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			NotBefore: now.Unix(),
			Id:        uuid.New().String(),
			Issuer:    JwtIssuer,
			IssuedAt:  now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	encoded, err = token.SignedString([]byte(s.secret))
	if err != nil {
		log.ErrorF("srv.auth.jwt: signing token failed", err)
		return "", errno.SignedTokenFailed.With(err)
	}

	err = s.rCli.ZAdd(s.getKeyFromUid(uid), redis.Z{
		Member: claim.Id,
		Score:  float64(expiredTime.Unix()),
	}).Err()
	if err != nil {
		log.ErrorF("srv.auth.jwt: push token id into redis error", err)
		return "", errno.RedisOperationErr.With(err)
	}

	// 摊销清理过期 token, 在当前时间或之前过期的 token 都会被清理
	go s.cleanExpiredJwt(s.getKeyFromUid(uid), 0, now.Unix())
	return
}

func (s *jwtSrv) cleanExpiredJwt(key string, min, max int64) {
	s.rCli.ZRemRangeByScore(key,
		strconv.FormatInt(min, 10),
		strconv.FormatInt(max, 10))
}
