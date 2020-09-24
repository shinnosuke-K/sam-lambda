package table

import "gorm.io/gorm"

type Table interface {
	HasTable(db *gorm.DB) bool
	CreateTable(db *gorm.DB) error
	GetBody() []byte
	Mapping(jsonBody []byte) error
	Insert(db *gorm.DB)
}

func New() []Table {
	return []Table{NewTicket(), NewUser(), NewOrganization()}
}
