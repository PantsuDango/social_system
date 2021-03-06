package result

type ListAllPost struct {
	AllPostCount int        `json:"AllPostCount"  form:"AllPostCount"`
	PostInfo     []PostInfo `json:"PostInfo"      form:"PostInfo"`
}

type PostInfo struct {
	ID           int      `json:"ID"              form:"ID"`
	Title        string   `json:"Title"           form:"Title"`
	Content      string   `json:"Content"         form:"Content"`
	PictureUrl   []string `json:"PictureUrl"      form:"PictureUrl"`
	Type         int      `json:"Type"            form:"Type"`
	FromId       int      `json:"FromId"          form:"FromId"`
	CreatedAt    string   `json:"CreateTime"      form:"CreateTime"`
	UpdatedAt    string   `json:"UpdateTime"      form:"UpdateTime"`
	UserInfo     UserInfo `json:"UserInfo"        form:"UserInfo"`
	CommentCount int      `json:"CommentCount"    form:"CommentCount"`
	QuotedCount  int      `json:"QuotedCount"     form:"QuotedCount"`
	StarCount    int      `json:"StarCount"       form:"StarCount"`
	IsStar       bool     `json:"IsStar"          form:"IsStar"`
}

type UserInfo struct {
	ID             int    `json:"ID"              form:"ID"`
	Nick           string `json:"Nick"            form:"Nick"`
	Username       string `json:"Username"        form:"Username"`
	Sex            int    `json:"Sex"             form:"Sex"`
	HeadImage      string `json:"HeadImage"       form:"HeadImage"`
	Email          string `json:"Email"           form:"Email"`
	Phone          string `json:"Phone"           form:"Phone"`
	AttentionCount int    `json:"AttentionCount"  form:"AttentionCount"`
	IsAttention    bool   `json:"IsAttention"     form:"IsAttention"`
	PostCount      int    `json:"PostCount"       form:"PostCount"`
}

type PersonalUserInfo struct {
	ID             int    `json:"ID"              form:"ID"`
	Nick           string `json:"Nick"            form:"Nick"`
	Username       string `json:"Username"        form:"Username"`
	Sex            int    `json:"Sex"             form:"Sex"`
	HeadImage      string `json:"HeadImage"       form:"HeadImage"`
	Email          string `json:"Email"           form:"Email"`
	Phone          string `json:"Phone"           form:"Phone"`
	CommentCount   int    `json:"CommentCount"    form:"CommentCount"`
	StarCount      int    `json:"StarCount"       form:"StarCount"`
	AttentionCount int    `json:"AttentionCount"  form:"AttentionCount"`
	PostCount      int    `json:"PostCount"       form:"PostCount"`
}

type UserInfoResult struct {
	UserInfo UserInfo       `json:"UserInfo"  form:"UserInfo"`
	PostInfo []UserPostInfo `json:"PostInfo"  form:"PostInfo"`
}

type PersonalUserInfoResult struct {
	UserInfo PersonalUserInfo `json:"UserInfo"  form:"UserInfo"`
	PostInfo []UserPostInfo   `json:"PostInfo"  form:"PostInfo"`
}

type UserPostInfo struct {
	ID           int      `json:"ID"              form:"ID"`
	Title        string   `json:"Title"           form:"Title"`
	Content      string   `json:"Content"         form:"Content"`
	PictureUrl   []string `json:"PictureUrl"      form:"PictureUrl"`
	Type         int      `json:"Type"            form:"Type"`
	FromId       int      `json:"FromId"          form:"FromId"`
	CreatedAt    string   `json:"CreateTime"      form:"CreateTime"`
	UpdatedAt    string   `json:"UpdateTime"      form:"UpdateTime"`
	CommentCount int      `json:"CommentCount"    form:"CommentCount"`
	QuotedCount  int      `json:"QuotedCount"     form:"QuotedCount"`
	StarCount    int      `json:"StarCount"       form:"StarCount"`
	IsStar       bool     `json:"IsStar"          form:"IsStar"`
}

type ShowPost struct {
	ID           int           `json:"ID"              form:"ID"`
	Title        string        `json:"Title"           form:"Title"`
	Content      string        `json:"Content"         form:"Content"`
	PictureUrl   []string      `json:"PictureUrl"      form:"PictureUrl"`
	Type         int           `json:"Type"            form:"Type"`
	FromId       int           `json:"FromId"          form:"FromId"`
	CreatedAt    string        `json:"CreateTime"      form:"CreateTime"`
	UpdatedAt    string        `json:"UpdateTime"      form:"UpdateTime"`
	UserInfo     UserInfo      `json:"UserInfo"        form:"UserInfo"`
	CommentCount int           `json:"CommentCount"    form:"CommentCount"`
	QuotedCount  int           `json:"QuotedCount"     form:"QuotedCount"`
	StarCount    int           `json:"StarCount"       form:"StarCount"`
	CommentInfo  []CommentInfo `json:"CommentInfo"     form:"CommentInfo"`
	IsStar       bool          `json:"IsStar"          form:"IsStar"`
}

type CommentInfo struct {
	ID        int      `json:"ID"          form:"ID"`
	Content   string   `json:"Content"     form:"Content"`
	CreatedAt string   `json:"CreateTime"  form:"CreateTime"`
	UserInfo  UserInfo `json:"UserInfo"    form:"UserInfo"`
}

type ShowUserAttention struct {
	ID        int    `json:"ID"              form:"ID"`
	Nick      string `json:"Nick"            form:"Nick"`
	Username  string `json:"Username"        form:"Username"`
	Sex       int    `json:"Sex"             form:"Sex"`
	HeadImage string `json:"HeadImage"       form:"HeadImage"`
	Email     string `json:"Email"           form:"Email"`
	Phone     string `json:"Phone"           form:"Phone"`
}
