package persistence

import (
	"all/domain"

	"gorm.io/gorm"
)

type Organization struct {
	DB            *gorm.DB
	Organizations domain.Organizations
}

func (o *Organization) Has() bool {
	return o.DB.Migrator().HasTable(&domain.Organization{})
}

func (o *Organization) CreateTable() error {
	return o.DB.Migrator().CreateTable(&domain.Organization{})
}

func (o *Organization) Insert() error {
	for _, org := range o.Organizations {
		if err := o.DB.Create(&org).Error; err != nil {
			return err
		}
	}
	return nil
}

func NewOrg(db *gorm.DB) *Organization {
	return &Organization{DB: db}
}
