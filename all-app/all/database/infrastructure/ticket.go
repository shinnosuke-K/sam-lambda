package infrastructure

import (
	"all/database/domain"

	"gorm.io/gorm"
)

type Ticket struct {
	DB      *gorm.DB
	Tickets domain.Tickets
}

func (t *Ticket) Has(db *gorm.DB) bool {
	return t.DB.Migrator().HasTable(&domain.Ticket{})
}

func (t *Ticket) CreateTable(db *gorm.DB) error {
	return t.DB.Migrator().CreateTable(&domain.Ticket{})
}

func (t *Ticket) Insert(db *gorm.DB) error {
	return nil
}
