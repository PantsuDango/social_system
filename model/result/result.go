package result

type ListAllPost struct {
	ID           int      `json:"ID"            form:"ID"`
	Title        string   `json:"Title"         form:"Title"`
	Content      string   `json:"Content"       form:"Content"`
	PictureUrl   []string `json:"PictureUrl"    form:"PictureUrl"`
	Type         int      `json:"Type"          form:"Type"`
	FromId       int      `json:"FromId"        form:"FromId"`
	CreatedAt    string   `json:"CreateTime"    form:"CreateTime"`
	UpdatedAt    string   `json:"UpdateTime"    form:"UpdateTime"`
	UserInfo     UserInfo `json:"UserInfo"      form:"UserInfo"`
	CommentCount int      `json:"CommentCount"  form:"CommentCount"`
	QuotedCount  int      `json:"QuotedCount"   form:"QuotedCount"`
	StartCount   int      `json:"StartCount"    form:"StartCount"`
}

type UserInfo struct {
	ID        int    `json:"ID"          form:"ID"`
	Nick      string `json:"Nick"        form:"Nick"`
	Username  string `json:"Username"    form:"Username"`
	Sex       int    `json:"Sex"         form:"Sex"`
	HeadImage string `json:"HeadImage"   form:"HeadImage"`
	Email     string `json:"Email"       form:"Email"`
	Phone     string `json:"Phone"       form:"Phone"`
}
