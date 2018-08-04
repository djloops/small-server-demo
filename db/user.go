package db

import "github.com/djloops/small-server-demo/models"

func CreateUser(u *models.User) error {
	return Mysql.Create(u).Error
}

func GetUser(ID int64) (*models.User, error) {
	user := models.User{}
	err := Mysql.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
