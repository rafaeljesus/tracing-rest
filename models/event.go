package models

import (
  "github.com/jinzhu/gorm"
)

type Event struct {
  gorm.Model
  Name string
  Status string
  Payload PropertyMap `type: jsonb not null default '{}'::jsonb`
}
