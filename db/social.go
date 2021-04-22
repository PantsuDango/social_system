package db

import "social_system/model/tables"

type SocialDB struct{}

func (SocialDB) GetUserInfo(UserName string) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("username = ? AND status= 0 ", UserName).First(&user).Error
	return user, err
}

func (SocialDB) QueryUserById(userId int) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("id = ? AND status = 0", userId).First(&user).Error
	return user, err
}

func (SocialDB) CreateUser(user tables.User) error {
	err := exeDB.Create(&user).Error
	return err
}
