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

func (SocialDB) SelectAllPost(offset, limit int) []tables.Post {
	var post []tables.Post
	exeDB.Offset(offset).Limit(limit).Order(`createtime desc`).Find(&post)
	return post
}

func (SocialDB) SelectAllPostCount() int {
	var count int
	exeDB.Model(&tables.Post{}).Count(&count)
	return count
}

func (SocialDB) SelectPostPictureMap(post_id int) []tables.PostPictureMap {
	var post_picture_map []tables.PostPictureMap
	exeDB.Where(`post_id = ?`, post_id).Find(&post_picture_map)
	return post_picture_map
}

func (SocialDB) SelectCommentCount(post_id int) int {
	var count int
	exeDB.Model(&tables.PostCommentMap{}).Where(`post_id = ?`, post_id).Count(&count)
	return count
}

func (SocialDB) SelectQuotedCount(post_id int) int {
	var count int
	exeDB.Model(&tables.PostQuotedMap{}).Where(`post_id = ?`, post_id).Count(&count)
	return count
}

func (SocialDB) SelectStartCount(post_id int) int {
	var count int
	exeDB.Model(&tables.PostStarMap{}).Where(`post_id = ?`, post_id).Count(&count)
	return count
}

func (SocialDB) SelectAttentionCount(user_id int) int {
	var count int
	exeDB.Model(&tables.UserAttentionMap{}).Where(`user_id = ?`, user_id).Count(&count)
	return count
}

func (SocialDB) SelectUserPost(user_id int) []tables.Post {
	var post []tables.Post
	exeDB.Where(`user_id = ?`, user_id).Find(&post)
	return post
}

func (SocialDB) ModifyUser(user tables.User) error {
	err := exeDB.Save(&user).Error
	return err
}
