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

func (SocialDB) CreatePost(post *tables.Post) error {
	err := exeDB.Create(&post).Error
	return err
}

func (SocialDB) CreatePostPictureMap(post_picture_map tables.PostPictureMap) error {
	err := exeDB.Create(&post_picture_map).Error
	return err
}

//func (SocialDB) SelectAllUser() ([]tables.User) {
//	var user []tables.User
//	exeDB.Where("status = 0").Find(&user)
//	return user
//}

func (SocialDB) SelectAllPost(offset, limit int) []tables.Post {
	var post []tables.Post
	exeDB.Offset(offset).Limit(limit).Order(`createtime desc`).Find(&post)
	return post
}
