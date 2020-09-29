package persistence

import (
	"all/domain"

	"gorm.io/gorm"
)

type User struct {
	DB    *gorm.DB
	Users domain.Users
}

func (t *User) Has() bool {
	return t.DB.Migrator().HasTable(&domain.Ticket{})
}

func (t *User) CreateTable() error {
	return t.DB.Migrator().CreateTable(&domain.Ticket{})
}

func (t *User) Insert() error {
	return nil
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}
