package models

type FeedCard struct {
	User User `json:"user"`
	Post Post	`json:"post"`
	CommentCount int64 `json:"comment_count"`
}
