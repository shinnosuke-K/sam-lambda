package table

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	CreateTime     time.Time `json:"create_time"`
	OrganizationID int64     `json:"organization_id"`
	Alias          string    `json:"alias"`
	Role           string    `json:"role"`
}

func (u *User) Mapping(jsonBody []byte) error {
	return json.Unmarshal(jsonBody, &u)
}

func (u User) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&u)
}

func (u User) CreateTable(db *gorm.DB) error {
	return db.Migrator().CreateTable(&u)
}

func (u User) Insert(db *gorm.DB) {
	panic("implement me")
}

func NewUser() Table {
	return &User{}
}
