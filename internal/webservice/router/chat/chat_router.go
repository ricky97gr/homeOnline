package chat

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/webservice/controller/web_im"
	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"
	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
	"github.com/gin-gonic/gin"
)

type ChatPlugin struct {
	manager.BasePlugin
}

func init() {
	p := &ChatPlugin{
		BasePlugin: manager.BasePlugin{
			PluginName:   "聊天服务",
			Version:      "0.0.1_base",
			Author:       "forgocode",
			Description:  "用于网络聊天室，即时通讯",
			ExecPath:     "",
			PluginStatus: manager.Stopped,
			ListenPort:   10002,
		},
	}
	manager.RegisterPlugin(p)
}

func (p *ChatPlugin) Name() string {
	return p.PluginName
}

func (p *ChatPlugin) Run() (*exec.Cmd, error) {
	if p.ExecPath == "" {
		return nil, nil
	}
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

func (p *ChatPlugin) Router() []manager.RouterInfo {
	return []manager.RouterInfo{
		{Group: "/normalUser", Path: "/ws", Method: "GET", Handles: []gin.HandlerFunc{web_im.ReceiveClientComm}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
	}
}

func (p *ChatPlugin) Uninstall() {

}

func (p *ChatPlugin) Upgrade() {

}
