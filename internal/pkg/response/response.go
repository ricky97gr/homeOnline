package response

import (
	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/pkg/response"
)

func Success(ctx *gin.Context, result interface{}, total int64, detail ...interface{}) {
	response.Success(ctx, result, total, detail)
}

func Failed(ctx *gin.Context, errCode int32, detail ...interface{}) {
	response.Failed(ctx, errCode, errCodeMap[errCode].msgCn, detail)
}
