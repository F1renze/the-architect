package redis

import (
	"fmt"
	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/go-redis/redis"
	"sync"
)

var (
	cli *redis.Client
	once sync.Once
)

func Init(cmsCli config.CMSClient) {
	once.Do(func() {
		initRedis(cmsCli)
	})
}

func GetClient() *redis.Client{
	return cli
}

type redisConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
	DBNum int `json:"db_num"`
	Password string `json:"password"`
	Timeout int `json:"timeout"`
}

func initRedis(cmsCli config.CMSClient) {
	cfg := new(redisConfig)
	err := cmsCli.Scan(constant.RedisCfgName, cfg)
	if err != nil {
		log.Fatal("获取 Redis 配置失败", err)
	}

	cli = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB: cfg.DBNum,
	})
	log.InfoF("初始化 Redis, 检测连接...")

	pong, err := cli.Ping().Result()
	if err != nil {
		log.Fatal("Redis 连接不可用, 请联系系统管理员", err)
	}
	log.InfoF("Redis 已连接, ping: %v", pong)
}


