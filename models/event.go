package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaeljesus/tracing-rest/db"
)

type Event struct {
	gorm.Model
	Name    string      `json: "name"`
	Status  string      `json: "status"`
	Payload PropertyMap `json: "payload" type:jsonb not null default '{}'::jsonb`
}

func Search(q Query) interface{} {
	return db.Repo.
		Where("name = ?", q.Name).
		Find(&Event{}).
		Value
}

func (e *Event) Save() *gorm.DB {
	return db.Repo.Create(e)
}
