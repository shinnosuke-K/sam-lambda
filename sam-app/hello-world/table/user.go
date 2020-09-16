package table

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int64
	Name           string
	Email          string
	CreateTime     time.Time
	OrganizationID int64
	Alias          string
	Role           string
}

func (u *User) Mapping(byteBody []byte) error {
	return json.Unmarshal(byteBody, &u)
}

func (u User) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&u)
}

func (u User) CreateTable(db *gorm.DB) {
	panic("implement me")
}

func (u User) Insert(db *gorm.DB) {
	panic("implement me")
}

func NewUser() Table {
	return &User{}
}
