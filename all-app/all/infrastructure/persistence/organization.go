package persistence

import (
	"all/domain"

	"gorm.io/gorm"
)

type Organization struct {
	DB            *gorm.DB
	Organizations domain.Organizations
}

func (t *Organization) Has() bool {
	return t.DB.Migrator().HasTable(&domain.Organization{})
}

func (t *Organization) CreateTable() error {
	return t.DB.Migrator().CreateTable(&domain.Organization{})
}

func (t *Organization) Insert() error {
	return nil
}

func NewOrg(db *gorm.DB) *Organization {
	return &Organization{DB: db}
}
