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
