package db

import "github.com/djloops/small-server-demo/models"

func CreateComment(u *models.Comment) error {
	return Mysql.Create(u).Error
}

func GetCommentCount(postID int64) (*int64, error) {
	var count int64
	err := Mysql.Model(models.Comment{}).Where("post_id = ? ", postID).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &count, nil
}
