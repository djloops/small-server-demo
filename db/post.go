package db

import (
	"github.com/djloops/small-server-demo/models"
	"fmt"
)

func CreatePost(u *models.Post) error {
	return Mysql.Create(u).Error
}

func GetPosts(board int8, offset int64, count int32) ([]models.Post, error) {
	posts := []models.Post{}

	err := Mysql.Where("board = ? ", board).Order("update_time").Offset(offset).Limit(count).Find(&posts).Error
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	return posts, err
}