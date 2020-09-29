package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createConnect() string {
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(zendesk-db:3306)"
	DBNAME := "zendesk"
	OPTION := "charset=utf8mb4&parseTime=True&loc=Local"

	return fmt.Sprintf("%s:%s@%s/%s?%s", USER, PASS, PROTOCOL, DBNAME, OPTION)
}

func NewDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(createConnect()), nil)
}

func Exec(db *gorm.DB, query string) error {
	return nil
}
