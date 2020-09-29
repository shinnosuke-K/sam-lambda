package domain

import "time"

type Users []User

type User struct {
	ID             int64
	Name           string
	Email          string
	CreateTime     time.Time
	OrganizationID int64
	Alias          string
	Role           string
}
