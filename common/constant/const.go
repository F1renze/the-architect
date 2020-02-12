package constant

import "time"

// jwt
const (
	JwtExpiredTime = 24 * time.Hour
	TokenIssuer = "micro.architect"
)

// redis
const (
	RedisKey4Log = "arch_log"
	JwtKeyPrefix = "arch_token_"

	SmsCodeKeyPrefix = "arch_sms_"
	SmsCodeExpiredTime = 5 * time.Minute
)

// config center
const (
	CfgPrefix        = "arch"
	CfgCenterAddrEnv = "ARCH_CONFIG_ADDR"
)

// micro
const (
	// Interval 决定服务每隔多久向服务发现重新注册
	RegisterInterval = time.Second * 45
	// TTL 决定多久后服务发现将失效的服务移除
	RegisterTTL = time.Second * 10
)

// tracing

// services

const (
	UserSrvCfgName = "srv.user"
	UserApiCfgName = "api.user"

	CasbinSrvCfgName = "srv.casbin"
	AuthSrvCfgName = "srv.auth"

	MysqlCfgName = "infra.mysql"
	RedisCfgName = "infra.redis"
)
