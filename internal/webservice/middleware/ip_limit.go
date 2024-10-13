package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type limiter struct {
	items  *sync.Map
	limit  float64
	burst  int
	length int
}

const maxIPCount = 10240

var iPLimiter = &limiter{
	items: &sync.Map{},
}

func IPLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		if ip == "127.0.0.1" {
			//不检查
			ctx.Next()
		}
		if iPLimiter.length >= maxIPCount {
			ctx.Abort()
		}
		l, ok := iPLimiter.items.LoadOrStore(ip, rate.NewLimiter(rate.Limit(iPLimiter.limit), iPLimiter.burst))
		if !ok {
			ctx.Abort()
			return
		}
		iPLimiter.length++
		c, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		if err := l.(*rate.Limiter).Wait(c); err != nil {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
