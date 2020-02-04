package constant

import "time"

// redis
const (
	RedisKey4Log = "arch_log"
)

// config center
const (
	CfgPrefix = "arch"
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

