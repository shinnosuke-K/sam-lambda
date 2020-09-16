package table

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID             int64
	CreateTime     time.Time
	UpdateTime     time.Time
	Type           string
	Subject        string
	Priority       string
	Status         string
	Tag            string
	RequesterID    int64
	AssigneeID     int64
	OrganizationID int64
}

func (t *Ticket) Mapping(byteBody []byte) error {
	return json.Unmarshal(byteBody, &t)
}

func (t Ticket) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&t)
}

func (t Ticket) CreateTable(db *gorm.DB) {
	panic("implement me")
}

func (t Ticket) Insert(db *gorm.DB) {
	panic("implement me")
}

func NewTicket() Table {
	return &Ticket{}
}
