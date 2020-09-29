package _interface

import "time"

type TicketsResponse struct {
	Tickets  []Ticket `json:"ticket"`
	NextPage string   `json:"next_page"`
}

type Ticket struct {
	ID             int64     `json:"id"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
	Type           string    `json:"type"`
	Subject        string    `json:"subject"`
	Priority       string    `json:"priority"`
	Status         string    `json:"status"`
	Tags           []string  `json:"tags"`
	RequesterID    int64     `json:"requester_id"`
	AssigneeID     int64     `json:"assignee_id"`
	OrganizationID int64     `json:"organization_id"`
}

type UsersResponse struct {
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

type OrgsResponse struct {
	Orgs     []Organization `json:"organization"`
	NextPage string         `json:"next_page"`
}

type Organization struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
}
