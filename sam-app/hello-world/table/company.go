package table

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID         int64
	Name       string
	CreateTime time.Time
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
