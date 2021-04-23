package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"social_system/model/params"
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
