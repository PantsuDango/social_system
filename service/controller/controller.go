package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"social_system/model/params"
	"social_system/model/result"
	"social_system/model/tables"
)

// 新建帖子
func (Controller Controller) AddPost(ctx *gin.Context, user tables.User) {
	var AddPostParams params.AddPostParams
	if err := ctx.ShouldBindBodyWith(&AddPostParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	var post tables.Post
	post.Title = AddPostParams.Title
	post.Content = AddPostParams.Content
	post.UserId = user.ID
	post.Type = 0
	err := Controller.SocialDB.CreatePost(&post)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "create post error", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	if len(AddPostParams.PictureUrl) > 0 {
		var post_picture_map tables.PostPictureMap
		for _, picture_url := range AddPostParams.PictureUrl {
			post_picture_map.PostId = post.ID
			post_picture_map.PictureUrl = picture_url
			err = Controller.SocialDB.CreatePostPictureMap(post_picture_map)
		}
		if err != nil {
			JSONFail(ctx, http.StatusOK, AccessDBError, "create post_picture_map error", gin.H{
				"Code":    AccessDBError,
				"Message": err.Error(),
			})
			return
		}
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 公共大厅查询所有帖子
func (Controller Controller) ListAllPost(ctx *gin.Context, user tables.User) {

	var ListAllPost params.ListAllPost
	if err := ctx.ShouldBindBodyWith(&ListAllPost, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	var ListAllPostResult []result.ListAllPost
	post := Controller.SocialDB.SelectAllPost(ListAllPost.Offset, ListAllPost.Limit)
	for _, tmp := range post {
		var ListAllPost result.ListAllPost
		ListAllPost.ID = tmp.ID
		ListAllPost.Title = tmp.Title
		ListAllPost.Content = tmp.Content
		ListAllPost.Type = tmp.Type
		ListAllPost.FromId = tmp.FromId
		ListAllPost.CreatedAt = tmp.CreatedAt.Format("2006-01-02 15:04:05")
		ListAllPost.UpdatedAt = tmp.UpdatedAt.Format("2006-01-02 15:04:05")
		ListAllPost.PictureUrl = make([]string, 0)
		post_picture_map := Controller.SocialDB.SelectPostPictureMap(tmp.ID)
		for _, val := range post_picture_map {
			ListAllPost.PictureUrl = append(ListAllPost.PictureUrl, val.PictureUrl)
		}

		people, _ := Controller.SocialDB.QueryUserById(tmp.UserId)
		ListAllPost.UserInfo.ID = people.ID
		ListAllPost.UserInfo.Nick = people.Nick
		ListAllPost.UserInfo.Username = people.Username
		ListAllPost.UserInfo.Email = people.Email
		ListAllPost.UserInfo.Phone = people.Phone
		ListAllPost.UserInfo.HeadImage = people.HeadImage
		ListAllPost.UserInfo.Sex = people.Sex

		count := Controller.SocialDB.SelectCommentCount(tmp.ID)
		ListAllPost.CommentCount = count
		count = Controller.SocialDB.SelectQuotedCount(tmp.ID)
		ListAllPost.QuotedCount = count
		count = Controller.SocialDB.SelectStartCount(tmp.ID)
		ListAllPost.StartCount = count

		ListAllPostResult = append(ListAllPostResult, ListAllPost)
	}

	JSONSuccess(ctx, http.StatusOK, ListAllPostResult)
}
