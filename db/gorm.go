package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "github.com/rafaeljesus/tracing-rest/models"
)

var Repo *gorm.DB

func Connect() error {
  conn, err := gorm.Open("postgres", "host=localhost dbname=tracing_rest_dev sslmode=disable")

  if err != nil {
    panic("failed to connect database")
  }

  if err := conn.DB().Ping(); err != nil {
    panic(err)
  }

  conn.DB().SetMaxIdleConns(10)
  conn.DB().SetMaxOpenConns(100)
  conn.AutoMigrate(&models.Event{})

  Repo = conn

  return nil
}
