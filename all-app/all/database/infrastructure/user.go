package infrastructure

import (
	"all/database/domain"

	"gorm.io/gorm"
)

type User struct {
	DB    *gorm.DB
	Users domain.Users
}

func (t *User) Has(db *gorm.DB) bool {
	return t.DB.Migrator().HasTable(&domain.Ticket{})
}

func (t *User) CreateTable(db *gorm.DB) error {
	return t.DB.Migrator().CreateTable(&domain.Ticket{})
}

func (t *User) Insert(db *gorm.DB) error {
	return nil
}
