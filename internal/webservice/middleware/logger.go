package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		host := ctx.ClientIP()
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		ctx.Next()
		raw := ctx.Request.URL.RawQuery
		status := ctx.Writer.Status()
		newlog.Logger.Infof("| %d | %s | '%s' | %s | %+v | \t%s\t \n", status, host, path, method, time.Since(start), raw)
	}
}
