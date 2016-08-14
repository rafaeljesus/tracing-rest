package models

import (
  "github.com/jinzhu/gorm"
  "github.com/rafaeljesus/tracing-rest/db"
)

type Event struct {
  gorm.Model
  Name string
  Status string
  Payload PropertyMap `type: jsonb not null default '{}'::jsonb`
}

func Search(q Query) interface{} {
  return db.Repo.
    Where("name = ?", q.Name).Or("status = ?",  q.Status).
    Find(&Event{}).
    Value
}

func (e *Event) Save() *gorm.DB {
  return db.Repo.Create(e)
}

