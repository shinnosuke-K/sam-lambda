package table

import (
	"encoding/json"
	"fmt"
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

	test := JsonOrgs{}
	for n := 0; n < 3; n++ {
		test.Orgs = append(test.Orgs, Organization{
			ID:         int64(n + 2),
			Name:       fmt.Sprintf("株式会社%d", n),
			CreateTime: time.Now().Add(time.Duration(n)),
		})
	}
	j, _ := json.Marshal(test)
	return j
}

func (c *JsonOrgs) Insert(db *gorm.DB) {

	for n := range c.Orgs {
		org := Organization{
			ID:         c.Orgs[n].ID,
			Name:       c.Orgs[n].Name,
			CreateTime: c.Orgs[n].CreateTime.Local(),
		}
		db.Create(&org)
	}
}

func NewOrganization() *JsonOrgs {
	return new(JsonOrgs)
}
