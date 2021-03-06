package persistence

import (
	"all/domain"

	"gorm.io/gorm"
)

type Ticket struct {
	DB      *gorm.DB
	Tickets domain.Tickets
}

func (t *Ticket) Has() bool {
	return t.DB.Migrator().HasTable(&domain.Ticket{})
}

func (t *Ticket) CreateTable() error {
	return t.DB.Migrator().CreateTable(&domain.Ticket{})
}

func (t *Ticket) Insert() error {
	for _, ticket := range t.Tickets {
		if err := t.DB.Create(&ticket).Error; err != nil {
			return err
		}
	}
	return nil
}

func NewTicket(db *gorm.DB) *Ticket {
	return &Ticket{DB: db}
}
