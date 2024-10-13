package friend

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/webservice/controller/plugin"
	"github.com/ricky97gr/homeOnline/internal/webservice/controller/statistic"
	"github.com/ricky97gr/homeOnline/internal/webservice/controller/system"
	"github.com/ricky97gr/homeOnline/internal/webservice/controller/user"
	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"
	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
	"github.com/gin-gonic/gin"
)

type BasePlugin struct {
	manager.BasePlugin
}

func init() {
	p := &BasePlugin{
		BasePlugin: manager.BasePlugin{
			PluginName:   "基础服务",
			Version:      "0.0.1_base",
			Author:       "forgocode",
			Description:  "基础设施，提供基础的系统设施",
			ExecPath:     "",
			PluginStatus: manager.Running,
			ListenPort:   10002,
		},
	}
	manager.RegisterPlugin(p)
}

func (p *BasePlugin) Name() string {
	return p.PluginName
}

func (p *BasePlugin) Run() (*exec.Cmd, error) {
	cmd := exec.Command(p.ExecPath, "-port", strconv.Itoa(int(p.ListenPort)))
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Printf("failed to wait plugin: %+v, err: %+v\n", p, err)
		}
	}()
	return cmd, nil
}

func (p *BasePlugin) Router() []manager.RouterInfo {
	return []manager.RouterInfo{
		{Group: "", Path: "/register", Method: "POST", Handles: []gin.HandlerFunc{system.Register}},
		{Group: "", Path: "/login", Method: "POST", Handles: []gin.HandlerFunc{system.Login}},

		{Group: "/admin", Path: "/user", Method: "GET", Handles: []gin.HandlerFunc{user.NormalGetAllUser}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/user", Method: "PUT", Handles: []gin.HandlerFunc{user.AdminUpdateUser}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/user", Method: "POST", Handles: []gin.HandlerFunc{user.AdminCreateUser}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/user", Method: "DELETE", Handles: []gin.HandlerFunc{user.AdminDeleteUser}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/normalUser", Path: "/info", Method: "GET", Handles: []gin.HandlerFunc{user.AdminGetUserInfo}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},

		{Group: "/admin", Path: "/operationLog", Method: "GET", Handles: []gin.HandlerFunc{system.AdminGetOperationLog}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},

		{Group: "/admin", Path: "/statistic/counts", Method: "GET", Handles: []gin.HandlerFunc{statistic.Counts}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/statistic/usertrend", Method: "GET", Handles: []gin.HandlerFunc{statistic.UserAddTrend}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/statistic/articletrend", Method: "GET", Handles: []gin.HandlerFunc{statistic.ArticleAddTrend}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/statistic/topictop5", Method: "GET", Handles: []gin.HandlerFunc{statistic.TopicTOP5}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/statistic/tagtop5", Method: "GET", Handles: []gin.HandlerFunc{statistic.TagTOP5}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/statistic/categorytop5", Method: "GET", Handles: []gin.HandlerFunc{statistic.CategoryTOP5}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/statistic/scoretop10", Method: "GET", Handles: []gin.HandlerFunc{statistic.ScoreTop10}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/statistic/userActive30", Method: "GET", Handles: []gin.HandlerFunc{statistic.UserActive30}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},

		{Group: "/admin", Path: "/version", Method: "GET", Handles: []gin.HandlerFunc{system.GetVersion}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/monitor", Method: "GET", Handles: []gin.HandlerFunc{system.GetMonitor}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},

		{Group: "/admin", Path: "/plugin", Method: "GET", Handles: []gin.HandlerFunc{plugin.AdminGetAllPlugin}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/plugin", Method: "PUT", Handles: []gin.HandlerFunc{plugin.AdminUpdatePlugin}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/plugin", Method: "POST", Handles: []gin.HandlerFunc{plugin.AdminCreatePlugin}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/plugin", Method: "DELETE", Handles: []gin.HandlerFunc{plugin.AdminDeletePlugin}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
	}
}

func (p *BasePlugin) Uninstall() {

}

func (p *BasePlugin) Upgrade() {

}
