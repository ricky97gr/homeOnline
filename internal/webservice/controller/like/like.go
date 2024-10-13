package like

import (
	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	like_service "github.com/ricky97gr/homeOnline/internal/webservice/service/like"
)

func GiveLike(ctx *gin.Context) {
	info := &like_service.UILike{}
	err := ctx.ShouldBindJSON(info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	uid := ctx.Request.Header.Get("userID")

	err = like_service.GiveLike(uid, info)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	if info.IsCancel {
		response.Success(ctx, "取消赞成功", 1)
		return
	}
	response.Success(ctx, "点赞成功", 1)
}
