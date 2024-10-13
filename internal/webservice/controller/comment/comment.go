package comment

import (
	"github.com/ricky97gr/homeOnline/pkg/paginate"
	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	commentService "github.com/ricky97gr/homeOnline/internal/webservice/service/comemnt"
	topicService "github.com/ricky97gr/homeOnline/internal/webservice/service/topic"
)

func UserCreateComment(ctx *gin.Context) {
	com := &commentService.UIComment{}
	err := ctx.ShouldBindJSON(com)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	if com.Context == "" {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	com.AuthorID = ctx.Request.Header.Get("userID")
	com.Address = ctx.Request.Header.Get("clientIP")
	com.User = ctx.Request.Header.Get("userName")
	err = commentService.UserCreateComment(com)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	err = topicService.AddUsedCountByTopicName(com.Topic)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "successfully", 1)
}

func UserGetComment(ctx *gin.Context) {
	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	comments, err := commentService.UserGetComment(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, comments, 1)
}

func UserGetCommentByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	comments, err := commentService.UserGetCommentByUserID(userID)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, comments, 1)
}

func UserGetFirstComment(ctx *gin.Context) {

	q, err := paginate.GetPageQuery(ctx)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	comments, err := commentService.UserGetFirstComment(q)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, comments, 1)
}

func UserGetChildComment(ctx *gin.Context) {
	type info struct {
		CommentID string `json:"commentID" form:"commentID"`
	}
	var id info
	err := ctx.ShouldBindQuery(&id)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	result, err := commentService.UserGetCommentByCommentID(id.CommentID)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)
}
