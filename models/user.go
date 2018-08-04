package models

import (
	"time"
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	AvatarUrl  string    `json:"avatar_url"`
}
func (User) TableName() string {
	return "user"
}