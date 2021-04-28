package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
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
			"Code":    IllegalRequestParameter,
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

	var ListAllPostParams params.ListAllPostParams
	if err := ctx.ShouldBindBodyWith(&ListAllPostParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	var PostInfo []result.PostInfo
	PostInfo = make([]result.PostInfo, 0)
	post := Controller.SocialDB.SelectAllPost(ListAllPostParams.Offset, ListAllPostParams.Limit)
	for _, tmp := range post {
		var ListAllPost result.PostInfo
		_, err := Controller.SocialDB.SelectPostStarMap(tmp.ID, user.ID)
		if err == gorm.ErrRecordNotFound {
			ListAllPost.IsStar = false
		} else {
			ListAllPost.IsStar = true
		}
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
		_, err = Controller.SocialDB.SelectUserAttentionMap(people.ID, user.ID)
		if err == gorm.ErrRecordNotFound {
			ListAllPost.UserInfo.IsAttention = false
		} else {
			ListAllPost.UserInfo.IsAttention = true
		}

		count := Controller.SocialDB.SelectCommentCount(tmp.ID)
		ListAllPost.CommentCount = count
		count = Controller.SocialDB.SelectQuotedCount(tmp.ID)
		ListAllPost.QuotedCount = count
		count = Controller.SocialDB.SelectStartCount(tmp.ID)
		ListAllPost.StarCount = count

		PostInfo = append(PostInfo, ListAllPost)
	}

	var ListAllPostResult result.ListAllPost
	count := Controller.SocialDB.SelectAllPostCount()
	ListAllPostResult.AllPostCount = count
	ListAllPostResult.PostInfo = PostInfo

	JSONSuccess(ctx, http.StatusOK, ListAllPostResult)
}

// 个人信息
func (Controller Controller) UserInfo(ctx *gin.Context, user tables.User) {

	var ListAllPostParams params.ListAllPostParams
	if err := ctx.ShouldBindBodyWith(&ListAllPostParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	var UserInfo result.PersonalUserInfo
	UserInfo.ID = user.ID
	UserInfo.Nick = user.Nick
	UserInfo.Username = user.Username
	UserInfo.Email = user.Email
	UserInfo.Phone = user.Phone
	UserInfo.HeadImage = user.HeadImage
	UserInfo.Sex = user.Sex
	// 我关注的用户数
	count := Controller.SocialDB.SelectAttentionCountByFollowerId(user.ID)
	UserInfo.AttentionCount = count
	// 我的帖子总数
	count = Controller.SocialDB.SelectUserPostCount(user.ID)
	UserInfo.PostCount = count

	var UserPostInfo []result.UserPostInfo
	UserPostInfo = make([]result.UserPostInfo, 0)
	post := Controller.SocialDB.SelectUserPost(user.ID, ListAllPostParams.Offset, ListAllPostParams.Limit)
	for _, tmp := range post {
		var ListAllPost result.UserPostInfo
		_, err := Controller.SocialDB.SelectPostStarMap(tmp.ID, user.ID)
		if err == gorm.ErrRecordNotFound {
			ListAllPost.IsStar = false
		} else {
			ListAllPost.IsStar = true
		}
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
		// 帖子评论数
		count = Controller.SocialDB.SelectCommentCount(tmp.ID)
		UserInfo.CommentCount += count
		ListAllPost.CommentCount = count
		// 帖子转发数
		count = Controller.SocialDB.SelectQuotedCount(tmp.ID)
		ListAllPost.QuotedCount = count
		// 帖子点赞数
		count = Controller.SocialDB.SelectStartCount(tmp.ID)
		UserInfo.StarCount += count
		ListAllPost.StarCount = count

		UserPostInfo = append(UserPostInfo, ListAllPost)
	}

	var UserInfoResult result.PersonalUserInfoResult
	UserInfoResult.UserInfo = UserInfo
	UserInfoResult.PostInfo = UserPostInfo
	JSONSuccess(ctx, http.StatusOK, UserInfoResult)
}

// 公共大厅查询所有帖子
func (Controller Controller) ModifyUser(ctx *gin.Context, user tables.User) {

	var UserInfoParams params.UserInfoParams
	if err := ctx.ShouldBindBodyWith(&UserInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	if UserInfoParams.Sex > 0 && user.Sex == 0 {
		user.Sex = UserInfoParams.Sex
	}
	if len(UserInfoParams.Nick) > 0 {
		user.Nick = UserInfoParams.Nick
	}
	if len(UserInfoParams.Phone) > 0 {
		user.Phone = UserInfoParams.Phone
	}
	if len(UserInfoParams.Email) > 0 {
		user.Email = UserInfoParams.Email
	}
	if len(UserInfoParams.HeadImage) > 0 {
		user.HeadImage = UserInfoParams.HeadImage
	}
	if len(UserInfoParams.Password) > 0 {
		s := UserInfoParams.OldPassword + user.Salt
		if fmt.Sprintf("%x", md5.Sum([]byte(s))) == user.Password {
			user.Salt = GetRandomString(8)
			s := UserInfoParams.Password + user.Salt
			user.Password = fmt.Sprintf("%x", md5.Sum([]byte(s)))
		} else {
			JSONFail(ctx, http.StatusOK, PasswordError, "OldPassword error.", gin.H{
				"Code":    "InvalidJSON",
				"Message": "OldPassword error.",
			})
			return
		}
	}

	err := Controller.SocialDB.ModifyUser(user)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "update user error", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 浏览别人信息
func (Controller Controller) ShowUserInfo(ctx *gin.Context, user tables.User) {

	var ShowUserInfoParams params.ShowUserInfoParams
	if err := ctx.ShouldBindBodyWith(&ShowUserInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	user_info, err := Controller.SocialDB.QueryUserById(ShowUserInfoParams.ID)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "select user error", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	var UserInfo result.UserInfo
	UserInfo.ID = user_info.ID
	UserInfo.Nick = user_info.Nick
	UserInfo.Username = user_info.Username
	UserInfo.Email = user_info.Email
	UserInfo.Phone = user_info.Phone
	UserInfo.HeadImage = user_info.HeadImage
	UserInfo.Sex = user_info.Sex
	count := Controller.SocialDB.SelectAttentionCount(ShowUserInfoParams.ID)
	UserInfo.AttentionCount = count
	_, err = Controller.SocialDB.SelectUserAttentionMap(ShowUserInfoParams.ID, user.ID)
	if err == gorm.ErrRecordNotFound {
		UserInfo.IsAttention = false
	} else {
		UserInfo.IsAttention = true
	}
	count = Controller.SocialDB.SelectUserPostCount(user_info.ID)
	UserInfo.PostCount = count

	var UserPostInfo []result.UserPostInfo
	UserPostInfo = make([]result.UserPostInfo, 0)
	post := Controller.SocialDB.SelectUserPost(user_info.ID, ShowUserInfoParams.Offset, ShowUserInfoParams.Limit)
	for _, tmp := range post {
		var ListAllPost result.UserPostInfo
		_, err := Controller.SocialDB.SelectPostStarMap(tmp.ID, user.ID)
		if err == gorm.ErrRecordNotFound {
			ListAllPost.IsStar = false
		} else {
			ListAllPost.IsStar = true
		}
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

		count = Controller.SocialDB.SelectCommentCount(tmp.ID)
		ListAllPost.CommentCount = count
		count = Controller.SocialDB.SelectQuotedCount(tmp.ID)
		ListAllPost.QuotedCount = count
		count = Controller.SocialDB.SelectStartCount(tmp.ID)
		ListAllPost.StarCount = count

		UserPostInfo = append(UserPostInfo, ListAllPost)
	}

	var UserInfoResult result.UserInfoResult
	UserInfoResult.UserInfo = UserInfo
	UserInfoResult.PostInfo = UserPostInfo
	JSONSuccess(ctx, http.StatusOK, UserInfoResult)
}

// 关注别人
func (Controller Controller) AddAttention(ctx *gin.Context, user tables.User) {

	var ShowUserInfoParams params.AddQuotedParams
	if err := ctx.ShouldBindBodyWith(&ShowUserInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	if ShowUserInfoParams.ID == user.ID {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "You can't pay attention to yourself", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": "You can't pay attention to yourself",
		})
		return
	}

	user_attention_map, err := Controller.SocialDB.SelectUserAttentionMap(ShowUserInfoParams.ID, user.ID)
	if err == nil && user_attention_map.ID > 0 {
		Controller.SocialDB.DeleteUserAttentionMap(user_attention_map)
	} else {
		user_attention_map.UserId = ShowUserInfoParams.ID
		user_attention_map.FollowerId = user.ID
		Controller.SocialDB.CreateUserAttentionMap(user_attention_map)
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 查看某帖子详细信息
func (Controller Controller) ShowPost(ctx *gin.Context, user tables.User) {

	var ShowUserInfoParams params.AddQuotedParams
	if err := ctx.ShouldBindBodyWith(&ShowUserInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	post, err := Controller.SocialDB.SelectPostInfo(ShowUserInfoParams.ID)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "select post error", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	var ShowPost result.ShowPost
	ShowPost.PictureUrl = make([]string, 0)
	ShowPost.CommentInfo = make([]result.CommentInfo, 0)

	ShowPost.ID = post.ID
	ShowPost.Title = post.Title
	ShowPost.Content = post.Content
	ShowPost.Type = post.Type
	ShowPost.FromId = post.FromId
	ShowPost.CreatedAt = post.CreatedAt.Format("2006-01-02 15:04:05")
	ShowPost.UpdatedAt = post.UpdatedAt.Format("2006-01-02 15:04:05")

	_, err = Controller.SocialDB.SelectPostStarMap(post.ID, user.ID)
	if err == gorm.ErrRecordNotFound {
		ShowPost.IsStar = false
	} else {
		ShowPost.IsStar = true
	}

	ShowPost.PictureUrl = make([]string, 0)
	post_picture_map := Controller.SocialDB.SelectPostPictureMap(post.ID)
	for _, val := range post_picture_map {
		ShowPost.PictureUrl = append(ShowPost.PictureUrl, val.PictureUrl)
	}

	people, _ := Controller.SocialDB.QueryUserById(post.UserId)
	ShowPost.UserInfo.ID = people.ID
	ShowPost.UserInfo.Nick = people.Nick
	ShowPost.UserInfo.Username = people.Username
	ShowPost.UserInfo.Email = people.Email
	ShowPost.UserInfo.Phone = people.Phone
	ShowPost.UserInfo.HeadImage = people.HeadImage
	ShowPost.UserInfo.Sex = people.Sex

	count := Controller.SocialDB.SelectCommentCount(post.ID)
	ShowPost.CommentCount = count
	count = Controller.SocialDB.SelectQuotedCount(post.ID)
	ShowPost.QuotedCount = count
	count = Controller.SocialDB.SelectStartCount(post.ID)
	ShowPost.StarCount = count

	post_comment_map := Controller.SocialDB.SelectPostCommentMap(post.ID)
	for _, tmp := range post_comment_map {
		var CommentInfo result.CommentInfo
		CommentInfo.ID = tmp.ID
		CommentInfo.Content = tmp.Content
		CommentInfo.CreatedAt = tmp.CreatedAt.Format("2006-01-02 15:04:05")

		people, _ := Controller.SocialDB.QueryUserById(tmp.UserId)
		CommentInfo.UserInfo.ID = people.ID
		CommentInfo.UserInfo.Nick = people.Nick
		CommentInfo.UserInfo.Username = people.Username
		CommentInfo.UserInfo.Email = people.Email
		CommentInfo.UserInfo.Phone = people.Phone
		CommentInfo.UserInfo.HeadImage = people.HeadImage
		CommentInfo.UserInfo.Sex = people.Sex

		ShowPost.CommentInfo = append(ShowPost.CommentInfo, CommentInfo)
	}

	JSONSuccess(ctx, http.StatusOK, ShowPost)
}

// 点赞帖子
func (Controller Controller) AddStar(ctx *gin.Context, user tables.User) {

	var ShowUserInfoParams params.AddQuotedParams
	if err := ctx.ShouldBindBodyWith(&ShowUserInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	post_star_map, err := Controller.SocialDB.SelectPostStarMap(ShowUserInfoParams.ID, user.ID)
	if err == nil && post_star_map.ID > 0 {
		Controller.SocialDB.DeletePostStarMap(post_star_map)
		JSONSuccess(ctx, http.StatusOK, "UnSuccess")
	} else {
		post_star_map.UserId = user.ID
		post_star_map.PostId = ShowUserInfoParams.ID
		Controller.SocialDB.CreatePostStarMap(post_star_map)
		JSONSuccess(ctx, http.StatusOK, "Success")
	}

}

// 点赞帖子
func (Controller Controller) AddComment(ctx *gin.Context, user tables.User) {

	var AddCommentParams params.AddCommentParams
	if err := ctx.ShouldBindBodyWith(&AddCommentParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	_, err := Controller.SocialDB.SelectPostInfo(AddCommentParams.ID)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "select post error", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	var post_comment_map tables.PostCommentMap
	post_comment_map.PostId = AddCommentParams.ID
	post_comment_map.Content = AddCommentParams.Content
	post_comment_map.UserId = user.ID
	Controller.SocialDB.CreatePostCommentMap(post_comment_map)

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 获取关注列表
func (Controller Controller) ShowUserAttention(ctx *gin.Context, user tables.User) {

	var ShowUserAttentionResult []result.ShowUserAttention
	ShowUserAttentionResult = make([]result.ShowUserAttention, 0)
	user_attention_map := Controller.SocialDB.SelectUserAttentionMapByFollowerId(user.ID)
	for _, tmp := range user_attention_map {
		var ShowUserAttention result.ShowUserAttention
		user, _ := Controller.SocialDB.QueryUserById(tmp.UserId)
		ShowUserAttention.ID = user.ID
		ShowUserAttention.Username = user.Username
		ShowUserAttention.Nick = user.Nick
		ShowUserAttention.Email = user.Email
		ShowUserAttention.Phone = user.Phone
		ShowUserAttention.Sex = user.Sex
		ShowUserAttention.HeadImage = user.HeadImage
		ShowUserAttentionResult = append(ShowUserAttentionResult, ShowUserAttention)
	}

	JSONSuccess(ctx, http.StatusOK, ShowUserAttentionResult)
}

// 转发帖子
func (Controller Controller) AddQuoted(ctx *gin.Context, user tables.User) {

	var AddQuotedParams params.AddQuotedParams
	if err := ctx.ShouldBindBodyWith(&AddQuotedParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IllegalRequestParameter,
			"Message": err.Error(),
		})
		return
	}

	post_info, err := Controller.SocialDB.SelectPostInfo(AddQuotedParams.ID)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "select post error", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	if post_info.UserId == user.ID {
		JSONFail(ctx, http.StatusOK, AccessDBError, "不能转发自己的帖子", gin.H{
			"Code":    AccessDBError,
			"Message": "不能转发自己的帖子",
		})
		return
	}

	var post tables.Post
	post.Title = post_info.Title
	post.Content = post_info.Content
	post.UserId = user.ID
	post.Type = 1
	post.FromId = post_info.UserId
	err = Controller.SocialDB.CreatePost(&post)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "create post error", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	post_picture_maps := Controller.SocialDB.SelectPostPictureMap(post_info.ID)
	if len(post_picture_maps) > 0 {
		var post_picture_map tables.PostPictureMap
		for _, tmp := range post_picture_maps {
			post_picture_map.PostId = post.ID
			post_picture_map.PictureUrl = tmp.PictureUrl
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
