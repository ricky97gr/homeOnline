package link

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/webservice/controller/like"
	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"
	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
	"github.com/gin-gonic/gin"
)

// 点赞，收藏，转发，
type LinkPlugin struct {
	manager.BasePlugin
}

func init() {
	p := &LinkPlugin{
		BasePlugin: manager.BasePlugin{
			PluginName:   "互动服务",
			Version:      "0.0.1_base",
			Author:       "forgocode",
			Description:  "开启文章，圈子的点赞，收藏，转发功能",
			ExecPath:     "",
			PluginStatus: manager.Stopped,
			ListenPort:   10002,
		},
	}
	manager.RegisterPlugin(p)
}

func (p *LinkPlugin) Name() string {
	return p.PluginName
}

func (p *LinkPlugin) Run() (*exec.Cmd, error) {
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

func (p *LinkPlugin) Router() []manager.RouterInfo {
	return []manager.RouterInfo{
		{Group: "/normalUser", Path: "/like", Method: "POST", Handles: []gin.HandlerFunc{like.GiveLike}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/normalUser", Path: "/unlike", Method: "POST", Handles: []gin.HandlerFunc{like.GiveLike}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
	}
}

func (p *LinkPlugin) Uninstall() {

}

func (p *LinkPlugin) Upgrade() {

}
