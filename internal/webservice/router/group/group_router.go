package group

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/webservice/controller/group"
	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"
	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
	"github.com/gin-gonic/gin"
)

type GroupPlugin struct {
	manager.BasePlugin
}

func init() {
	p := &GroupPlugin{
		BasePlugin: manager.BasePlugin{
			PluginName:   "群组服务",
			Version:      "0.0.1_base",
			Author:       "forgocode",
			Description:  "激活群组功能，进行群聊",
			ExecPath:     "",
			PluginStatus: manager.Stopped,
			ListenPort:   10002,
		},
	}
	manager.RegisterPlugin(p)
}

func (p *GroupPlugin) Name() string {
	return p.PluginName
}

func (p *GroupPlugin) Run() (*exec.Cmd, error) {
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

func (p *GroupPlugin) Router() []manager.RouterInfo {
	return []manager.RouterInfo{
		{Group: "/normalUser", Path: "/grouplist", Method: "GET", Handles: []gin.HandlerFunc{group.GetAllGroupByUserUID}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/normalUser", Path: "/groupmember/:id", Method: "GET", Handles: []gin.HandlerFunc{group.GetMemberByGroupUID}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
	}
}

func (p *GroupPlugin) Uninstall() {

}

func (p *GroupPlugin) Upgrade() {

}
