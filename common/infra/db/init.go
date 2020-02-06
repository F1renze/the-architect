package db

import (
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/f1renze/the-architect/common/config"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func Init(cmsCli config.CMSClient) {
	once.Do(func() {
		initMysql(cmsCli)
	})
}

func GetDB() *sqlx.DB {
	return db
}
