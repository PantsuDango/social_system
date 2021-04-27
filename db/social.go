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

func (SocialDB) SelectAttentionCountByFollowerId(follower_id int) int {
	var count int
	exeDB.Model(&tables.UserAttentionMap{}).Where(`follower_id = ?`, follower_id).Count(&count)
	return count
}

func (SocialDB) SelectUserPost(user_id, offset, limit int) []tables.Post {
	var post []tables.Post
	exeDB.Where(`user_id = ?`, user_id).Offset(offset).Limit(limit).Order(`createtime desc`).Find(&post)
	return post
}

func (SocialDB) SelectUserPostCount(user_id int) int {
	var count int
	exeDB.Where(`user_id = ?`, user_id).Count(&count)
	return count
}

func (SocialDB) ModifyUser(user tables.User) error {
	err := exeDB.Save(&user).Error
	return err
}

func (SocialDB) SelectUserAttentionMap(user_id, follower_id int) (tables.UserAttentionMap, error) {
	var user_attention_map tables.UserAttentionMap
	err := exeDB.Where(`user_id = ? and follower_id = ?`, user_id, follower_id).First(&user_attention_map).Error
	return user_attention_map, err
}

func (SocialDB) DeleteUserAttentionMap(user_attention_map tables.UserAttentionMap) {
	exeDB.Delete(&user_attention_map)
}

func (SocialDB) CreateUserAttentionMap(user_attention_map tables.UserAttentionMap) {
	exeDB.Create(&user_attention_map)
}

func (SocialDB) SelectPostInfo(id int) (tables.Post, error) {
	var post tables.Post
	err := exeDB.Where(`id = ?`, id).First(&post).Error
	return post, err
}

func (SocialDB) SelectPostCommentMap(post_id int) []tables.PostCommentMap {
	var post_comment_map []tables.PostCommentMap
	exeDB.Where(`post_id = ?`, post_id).Find(&post_comment_map)
	return post_comment_map
}

func (SocialDB) SelectPostStarMap(post_id, user_id int) (tables.PostStarMap, error) {
	var post_star_map tables.PostStarMap
	err := exeDB.Where(`post_id = ? and user_id = ?`, post_id, user_id).First(&post_star_map).Error
	return post_star_map, err
}

func (SocialDB) DeletePostStarMap(post_star_map tables.PostStarMap) {
	exeDB.Delete(&post_star_map)
}

func (SocialDB) CreatePostStarMap(post_star_map tables.PostStarMap) {
	exeDB.Create(&post_star_map)
}

func (SocialDB) CreatePostCommentMap(post_comment_map tables.PostCommentMap) {
	exeDB.Create(&post_comment_map)
}
