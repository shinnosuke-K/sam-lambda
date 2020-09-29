package domain

import "time"

type Organizations []Organization

type Organization struct {
	ID         int64
	Name       string
	CreateTime time.Time
}
