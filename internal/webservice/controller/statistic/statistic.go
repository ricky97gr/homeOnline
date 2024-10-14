package statistic

import (
	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	categoryService "github.com/ricky97gr/homeOnline/internal/webservice/service/category"
	commentService "github.com/ricky97gr/homeOnline/internal/webservice/service/comemnt"
	systemService "github.com/ricky97gr/homeOnline/internal/webservice/service/system"
	tagService "github.com/ricky97gr/homeOnline/internal/webservice/service/tag"
	topicService "github.com/ricky97gr/homeOnline/internal/webservice/service/topic"
)

type StaticCountInfo struct {
	ArticleTotal      int64 `json:"articleTotal"`
	TagTotal          int64 `json:"tagTotal"`
	CategoryTotal     int64 `json:"categoryTotal"`
	TopicTotal        int64 `json:"topicTotal"`
	ShortCommentTotal int64 `json:"shortCommentTotal"`
}

func Counts(ctx *gin.Context) {

	tagCount, err := tagService.AdminGetTagCount()
	if err != nil {
		newlog.Logger.Errorf("failed to statistic tag count, err: %+v\n", err)
	}
	categoryCount, err := categoryService.AdminGetCategoryCount()
	if err != nil {
		newlog.Logger.Errorf("failed to statistic category count, err: %+v\n", err)
	}
	topicCount, err := topicService.AdminGetTopicCount()
	if err != nil {
		newlog.Logger.Errorf("failed to statistic topic count, err: %+v\n", err)
	}
	articleCount, err := 0, nil
	if err != nil {
		newlog.Logger.Errorf("failed to statistic article count, err: %+v\n", err)
	}
	shortCommentCount, err := commentService.AdminGetShortCommentCount()
	if err != nil {
		newlog.Logger.Errorf("failed to statistic short comment count, err: %+v\n", err)
	}
	result := &StaticCountInfo{
		ArticleTotal:      int64(articleCount),
		TagTotal:          tagCount,
		CategoryTotal:     categoryCount,
		TopicTotal:        topicCount,
		ShortCommentTotal: shortCommentCount,
	}
	response.Success(ctx, *result, 1)
}

func UserAddTrend(ctx *gin.Context) {
	result, err := systemService.UserAddTrend()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)
}

func ArticleAddTrend(ctx *gin.Context) {
	result, err := commentService.CommentAddTrend()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)
}

func TopicTOP5(ctx *gin.Context) {
	result, err := topicService.AdminGetTopicTOP(5)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)

}
func TagTOP5(ctx *gin.Context) {
	result, err := tagService.AdminGetTagTOP(5)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)
}
func CategoryTOP5(ctx *gin.Context) {
	result, err := categoryService.AdminGetCategoryTOP(5)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)
}

func ScoreTop10(ctx *gin.Context) {
	users, err := systemService.UserScoreTop()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, users, int64(len(users)))
}

func UserActive30(ctx *gin.Context) {
	result, err := systemService.UserActiveCount()
	if err != nil {
		newlog.Logger.Errorf("failed to get active user count: %+v\n", err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)
}

func UserRoleCount(ctx *gin.Context) {
	result, err := systemService.UserRoleCount()
	if err != nil {
		newlog.Logger.Errorf("failed to get active user count: %+v\n", err)
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, result, 1)
}
