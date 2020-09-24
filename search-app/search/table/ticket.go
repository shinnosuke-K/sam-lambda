package table

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/jinzhu/gorm"
)

func createConnect() string {
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(zendesk-db:3306)"
	DBNAME := "zendesk"
	OPTION := "charset=utf8mb4&parseTime=True&loc=Local"

	return fmt.Sprintf("%s:%s@%s/%s?%s", USER, PASS, PROTOCOL, DBNAME, OPTION)
}

func dbOpen() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(createConnect()), nil)
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
