package tables

import "time"

type User struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	Nick      string    `json:"Nick"        form:"Nick"        gorm:"column:nick"`
	Username  string    `json:"Username"    form:"Username"    gorm:"column:username"`
	Password  string    `json:"Password"    form:"Password"    gorm:"column:password"`
	Salt      string    `json:"Salt"        form:"Salt"        gorm:"column:salt"`
	Sex       int       `json:"Sex"         form:"Sex"         gorm:"column:sex"`
	HeadImage string    `json:"HeadImage"   form:"HeadImage"   gorm:"column:head_image"`
	Email     string    `json:"Email"       form:"Email"       gorm:"column:email"`
	Phone     string    `json:"Phone"       form:"Phone"       gorm:"column:phone"`
	Status    int       `json:"Status"      form:"Status"      gorm:"column:status"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  form:"UpdateTime"  gorm:"column:lastupdate"`
}

func (User) TableName() string {
	return "user"
}

type UserAttentionMap struct {
	ID         int       `json:"ID"          form:"ID"          gorm:"column:id"`
	UserId     int       `json:"UserId"      form:"UserId"      gorm:"column:user_id"`
	FollowerId int       `json:"FollowerId"  form:"FollowerId"  gorm:"column:follower_id"`
	CreatedAt  time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
}

func (UserAttentionMap) TableName() string {
	return "user_attention_map"
}

type Post struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	Title     string    `json:"Title"       form:"Title"       gorm:"column:title"`
	Content   string    `json:"Content"     form:"Content"     gorm:"column:content"`
	UserId    int       `json:"UserId"      form:"UserId"      gorm:"column:user_id"`
	Type      int       `json:"Type"        form:"Type"        gorm:"column:type"`
	FromId    int       `json:"FromId"      form:"FromId"      gorm:"column:from_id"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  form:"UpdateTime"  gorm:"column:lastupdate"`
}

func (Post) TableName() string {
	return "post"
}

type PostCommentMap struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	PostId    int       `json:"PostId"      form:"PostId"      gorm:"column:post_id"`
	Content   string    `json:"Content"     form:"Content"     gorm:"column:content"`
	UserId    int       `json:"UserId"      form:"UserId"      gorm:"column:user_id"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
}

func (PostCommentMap) TableName() string {
	return "post_comment_map"
}

type PostPictureMap struct {
	ID         int       `json:"ID"          form:"ID"          gorm:"column:id"`
	PostId     int       `json:"PostId"      form:"PostId"      gorm:"column:post_id"`
	PictureUrl string    `json:"PictureUrl"  form:"PictureUrl"  gorm:"column:picture_url"`
	CreatedAt  time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
}

func (PostPictureMap) TableName() string {
	return "post_picture_map"
}

type PostQuotedMap struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	PostId    int       `json:"PostId"      form:"PostId"      gorm:"column:post_id"`
	UserId    int       `json:"UserId"      form:"UserId"      gorm:"column:user_id"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
}

func (PostQuotedMap) TableName() string {
	return "post_quoted_map"
}

type PostStarMap struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	PostId    int       `json:"PostId"      form:"PostId"      gorm:"column:post_id"`
	UserId    int       `json:"UserId"      form:"UserId"      gorm:"column:user_id"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
}

func (PostStarMap) TableName() string {
	return "post_star_map"
}
