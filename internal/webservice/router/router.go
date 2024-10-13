package router

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"
	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
	"github.com/ricky97gr/homeOnline/pkg/bininfo"
)

func Start() {
	gin.SetMode(bininfo.GinLogMode)
	bininfo.StartTime = time.Now().UnixMilli()
	engine := gin.New()
	engine.Use(middleware.Logger(), middleware.Recovery(), middleware.RouterManager())

	manager.RegisterRouter(engine)

	engine.Run(":8800")
}
