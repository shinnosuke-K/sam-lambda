package table

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
}

func (c *Company) Mapping(jsonBody []byte) error {
	return json.Unmarshal(jsonBody, &c)
}

func (c Company) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&c)
}

func (c Company) CreateTable(db *gorm.DB) error {
	return db.Migrator().CreateTable(&c)
}

func (c Company) Insert(db *gorm.DB) {
	panic("implement me")
}

func NewCompany() Table {
	return &Company{}
}
