package models

import "time"

type Post struct {
	ID         int64     `json:"id"`
	OwnerID		int64    `json:"owner_id"`
	Board		int8 	`json:"board"`
	Subject       string    `json:"subject"`
	Content       string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
func (Post) TableName() string {
	return "post"
}
