package params

type ModActIndex struct {
	Module string `json:"Module" form:"Module" binding:"required"`
	Action string `json:"Action" form:"Action" binding:"required"`
}

type AddPostParams struct {
	Title      string   `json:"Title"       form:"Title"       binding:"required"`
	Content    string   `json:"Content"     form:"Content"`
	PictureUrl []string `json:"PictureUrl"  form:"PictureUrl"`
}

type ListAllPostParams struct {
	Offset int `json:"Offset" form:"Offset"`
	Limit  int `json:"Limit"  form:"Limit"   binding:"required"`
}

type UserInfoParams struct {
	Nick        string `json:"Nick"         form:"Nick"`
	Password    string `json:"Password"     form:"Password"`
	OldPassword string `json:"OldPassword"  form:"OldPassword"`
	Sex         int    `json:"Sex"          form:"Sex"`
	HeadImage   string `json:"HeadImage"    form:"HeadImage"`
	Email       string `json:"Email"        form:"Email"`
	Phone       string `json:"Phone"        form:"Phone"`
}

type ShowUserInfoParams struct {
	ID     int `json:"ID"     form:"ID"      binding:"required"`
	Offset int `json:"Offset" form:"Offset"`
	Limit  int `json:"Limit"  form:"Limit"   binding:"required"`
}

type AddCommentParams struct {
	ID      int    `json:"ID"       form:"ID"       binding:"required"`
	Content string `json:"content"  form:"content"  binding:"required"`
}
