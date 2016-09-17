package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rafaeljesus/tracing-rest/support"
	"os"
)

var Repo *gorm.DB

func Connect() error {
	if os.Getenv("TRACKING_REST_DB") == "" {
		os.Setenv("TRACKING_REST_DB", support.Config["TRACKING_REST_DB"])
	}

	conn, err := gorm.Open("postgres", os.Getenv("TRACKING_REST_DB"))
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
