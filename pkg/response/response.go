package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, result interface{}, total int64, detail ...interface{}) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code": 200,
			"msg":  "handle successfully",
			//"detail": detail[0],
			"result": result,
			"total":  total,
		})
}

func Failed(ctx *gin.Context, errCode int32, msg string, detail ...interface{}) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code": errCode,
			"msg":  msg,
			//"detail": detail[0],
		})
}
