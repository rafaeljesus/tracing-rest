package events

import (
  "github.com/jinzhu/gorm"
)

type Event struct {
  gorm.Model
  Name string
  Status string
  Payload string `sql: "type:JSONB DEFAULT '{}'::JSONB"`
}
