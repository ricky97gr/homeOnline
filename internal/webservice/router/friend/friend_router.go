package friend

import (
	"github.com/ricky97gr/homeOnline/internal/webservice/controller/friend"
	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"os/exec"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
)

type FriendPlugin struct {
	manager.BasePlugin
}

func init() {
	p := &FriendPlugin{
		BasePlugin: manager.BasePlugin{
			PluginName:   "好友服务",
			Version:      "0.0.1_base",
			Author:       "forgocode",
			Description:  "激活好友功能，可以添加好友",
			ExecPath:     "",
			PluginStatus: manager.Stopped,
			ListenPort:   10002,
		},
	}
	manager.RegisterPlugin(p)
}

func (p *FriendPlugin) Name() string {
	return p.PluginName
}

func (p *FriendPlugin) Run() (*exec.Cmd, error) {
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

func (p *FriendPlugin) Router() []manager.RouterInfo {
	return []manager.RouterInfo{
		{Group: "/normalUser", Path: "/friend", Method: "POST", Handles: []gin.HandlerFunc{friend.UserCreateFriendShip}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/normalUser", Path: "/friend", Method: "DELETE", Handles: []gin.HandlerFunc{friend.UserDeleteFriendShip}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/normalUser", Path: "/friendFollow/list", Method: "GET", Handles: []gin.HandlerFunc{friend.UserGetFriendFollowList}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/normalUser", Path: "/friendFollowed/list", Method: "GET", Handles: []gin.HandlerFunc{friend.UserGetFriendFollowedList}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
	}
}

func (p *FriendPlugin) Uninstall() {

}

func (p *FriendPlugin) Upgrade() {

}
