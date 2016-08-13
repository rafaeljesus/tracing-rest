package models

import (
  "github.com/jinzhu/gorm"
)

type Event struct {
  gorm.Model
  Name string
  Status string
  Payload map[string]interface{} `type: jsonb not null default '{}'::jsonb`
}
