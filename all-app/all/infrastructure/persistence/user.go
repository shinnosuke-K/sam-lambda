package persistence

import (
	"all/domain"

	"gorm.io/gorm"
)

type User struct {
	DB    *gorm.DB
	Users domain.Users
}

func (u *User) Has() bool {
	return u.DB.Migrator().HasTable(&domain.User{})
}

func (u *User) CreateTable() error {
	return u.DB.Migrator().CreateTable(&domain.User{})
}

func (u *User) Insert() error {
	for _, user := range u.Users {
		if err := u.DB.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}
