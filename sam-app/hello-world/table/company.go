package table

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type JsonCompanies struct {
	Companies []Company `json:"company"`
	NextPage  string    `json:"next_page"`
}

type Company struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
}

func (c *JsonCompanies) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&Company{})
}

func (c *JsonCompanies) CreateTable(db *gorm.DB) error {
	return db.Migrator().CreateTable(&Company{})
}

func (c *JsonCompanies) Mapping(jsonBody []byte) error {
	err := json.Unmarshal(jsonBody, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *JsonCompanies) GetBody() []byte {
	return nil
}

func (c *JsonCompanies) Insert(db *gorm.DB) {
	panic("implement me")
}

func NewCompany() *JsonCompanies {
	return new(JsonCompanies)
}
