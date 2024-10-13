package middleware

import (
	"github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
	"github.com/gin-gonic/gin"
)

func RouterManager() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if manager.ISRouterPass(ctx.FullPath()) {
			ctx.Next()
			return
		}
		response.Failed(ctx, response.ErrRight)
		ctx.Abort()
	}
}
