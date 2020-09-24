package table

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type JsonUsers struct {
	Users    []User `json:"users"`
	NextPage string `json:"next_page"`
}

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	CreateTime     time.Time `json:"create_time"`
	OrganizationID int64     `json:"organization_id"`
	Alias          string    `json:"alias"`
	Role           string    `json:"role"`
}

func (u *JsonUsers) HasTable(db *gorm.DB) bool {
	return db.Migrator().HasTable(&User{})
}

func (u *JsonUsers) CreateTable(db *gorm.DB) error {
	return db.Migrator().CreateTable(&User{})
}

func (u *JsonUsers) Mapping(jsonBody []byte) error {

	err := json.Unmarshal(jsonBody, &u)
	if err != nil {
		return err
	}
	return nil
}

func (u *JsonUsers) GetBody() []byte {

	test := JsonUsers{}
	for n := 0; n < 3; n++ {
		test.Users = append(test.Users, User{
			ID:             int64(n + 1),
			Name:           fmt.Sprintf("%d", n+1),
			Email:          "",
			CreateTime:     time.Now().Local().Add(time.Duration(n)),
			OrganizationID: int64(n + 2),
			Alias:          "",
			Role:           "",
		})
	}
	j, _ := json.Marshal(test)
	return j
}

func (u JsonUsers) Insert(db *gorm.DB) {

	for n := range u.Users {
		user := User{
			ID:             u.Users[n].ID,
			Name:           u.Users[n].Name,
			Email:          u.Users[n].Email,
			CreateTime:     u.Users[n].CreateTime.Local(),
			OrganizationID: u.Users[n].OrganizationID,
			Alias:          u.Users[n].Alias,
			Role:           u.Users[n].Role,
		}
		db.Create(&user)
	}
}

func NewUser() *JsonUsers {
	return new(JsonUsers)
}
