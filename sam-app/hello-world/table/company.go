package table

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type JsonOrgs struct {
	Orgs     []Organization `json:"organization"`
	NextPage string         `json:"next_page"`
}

type Organization struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
}

func (c *JsonOrgs) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&Organization{})
}

func (c *JsonOrgs) CreateTable(db *gorm.DB) error {
	return db.Migrator().CreateTable(&Organization{})
}

func (c *JsonOrgs) Mapping(jsonBody []byte) error {
	err := json.Unmarshal(jsonBody, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *JsonOrgs) GetBody() []byte {
	return nil
}

func (c *JsonCompanies) Insert(db *gorm.DB) {
	panic("implement me")
}

func NewOrganization() *JsonOrgs {
	return new(JsonOrgs)
}
