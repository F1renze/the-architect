package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/f1renze/the-architect/common/config"
	"github.com/f1renze/the-architect/common/utils/log"
)

type mysqlConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	DBName      string `json:"db_name"`
	MaxIdleConn int    `json:"max_idle_conn"`
	MaxOpenConn int    `json:"max_open_conn"`
}

func initMysql(cmsCli config.CMSClient) {
	cfg := &mysqlConfig{}
	err := cmsCli.Scan("infra.mysql", cfg)
	if err != nil {
		log.Fatal("获取 mysql 配置失败", err)
	}

	// TODO delete
	//cfgCenterAddr := os.Getenv(constant.CfgCenterAddrEnv)
	//if strings.HasPrefix(cfgCenterAddr, "127.0.0.1") {
	//	cfg.Host = "127.0.0.1"
	//}

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, true, "Local",
	)
	log.DebugF("url: %s", url)

	// 创建连接
	db, err = sqlx.Connect("mysql", url)
	if err != nil {
		log.Fatal("建立 mysql 连接失败", err)
	}

	// 最大连接数
	db.SetMaxOpenConns(cfg.MaxOpenConn)

	// 最大闲置数
	db.SetMaxIdleConns(cfg.MaxIdleConn)

	// 激活链接
	if err = db.Ping(); err != nil {
		log.Fatal("无法激活 mysql 连接", err)
	}
}
