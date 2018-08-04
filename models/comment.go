package models

import "time"

type Comment struct {
	ID         int64     `json:"id"`
	PostID		int64    `json:"post_id"`
	Content       string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
func (Comment) TableName() string {
	return "comment"
}
