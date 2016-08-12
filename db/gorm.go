package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "github.com/rafaeljesus/tracing-rest/api/events"
)

func Connect() error {
  db, err := gorm.Open("postgres", "dbname=tracing_rest_dev sslmode=disable")

  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&events.Event{})

  return nil
}
