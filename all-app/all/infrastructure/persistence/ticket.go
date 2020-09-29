package persistence

import (
	"all/domain"

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

func NewTicket(db *gorm.DB) *Ticket {
	return &Ticket{DB: db}
}
