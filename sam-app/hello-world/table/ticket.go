package table

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type JsonTicket struct {
	Contents []Content `json:"ticket"`
	NextPage string    `json:"next_page"`
}

type Content struct {
	ID             int64     `json:"id"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
	Type           string    `json:"type"`
	Subject        string    `json:"subject"`
	Priority       string    `json:"priority"`
	Status         string    `json:"status"`
	Tags           []string  `json:"tags"`
	RequesterID    int64     `json:"requester_id"`
	AssigneeID     int64     `json:"assignee_id"`
	OrganizationID int64     `json:"organization_id"`
}

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

func (t Ticket) Mapping(jsonBody []byte) (JsonTicket, error) {
	var jt JsonTicket
	err := json.Unmarshal(jsonBody, &jt)
	if err != nil {
		return JsonTicket{}, err
	}
	return jt, nil
}

func (t Ticket) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&t)
}

func (t Ticket) CreateTable(db *gorm.DB) error {
	return db.Migrator().CreateTable(&t)
}

func (t Ticket) Insert(db *gorm.DB) {
	panic("implement me")
}

func NewTicket() Table {
	return &Ticket{}
}
