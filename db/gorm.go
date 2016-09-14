package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Repo *gorm.DB

func Connect() error {
	cfg, err := dbConfig()
	if err != nil {
		panic(err)
	}

	conn, err := gorm.Open(cfg.drv, cfg.open)
	if err != nil {
		panic(err)
	}

	if err := conn.DB().Ping(); err != nil {
		panic(err)
	}

	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(100)

	Repo = conn

	return nil
}
