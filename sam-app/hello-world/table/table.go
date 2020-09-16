package table

import "gorm.io/gorm"

type Table interface {
	Mapping(byteBody []byte) error
	HasTable(db *gorm.DB) bool
	CreateTable(db *gorm.DB)
	Insert(db *gorm.DB)
}

func Mapping(t Table, byteBody []byte) error {
	return t.Mapping(byteBody)
}

func HasTable(t Table, db *gorm.DB) bool {
	return t.HasTable(db)
}

func CreateTable(t Table, db *gorm.DB) {
	t.CreateTable(db)
}
