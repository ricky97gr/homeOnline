package comment

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/webservice/controller/comment"
	"github.com/ricky97gr/homeOnline/internal/webservice/controller/topic"
	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"
	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
	"github.com/gin-gonic/gin"
)

type CommentPlugin struct {
	manager.BasePlugin
}

func init() {
	p := &CommentPlugin{
		BasePlugin: manager.BasePlugin{
			PluginName:   "评论服务",
			Version:      "0.0.1_base",
			Author:       "forgocode",
			Description:  "用于评论文章，兴趣圈子",
			ExecPath:     "",
			PluginStatus: manager.Stopped,
			ListenPort:   10002,
		},
	}
	manager.RegisterPlugin(p)
}

func (p *CommentPlugin) Name() string {
	return p.PluginName
}

func (p *CommentPlugin) Run() (*exec.Cmd, error) {
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

func (p *CommentPlugin) Router() []manager.RouterInfo {
	return []manager.RouterInfo{
		{Group: "/normalUser", Path: "/comment", Method: "POST", Handles: []gin.HandlerFunc{comment.UserCreateComment}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "", Path: "/comment", Method: "GET", Handles: []gin.HandlerFunc{comment.UserGetComment}},
		{Group: "", Path: "/firstcomment", Method: "GET", Handles: []gin.HandlerFunc{comment.UserGetFirstComment}},
		{Group: "", Path: "/comment/child", Method: "GET", Handles: []gin.HandlerFunc{comment.UserGetChildComment}},
		{Group: "/normalUser", Path: "/commentList/:userID", Method: "GET", Handles: []gin.HandlerFunc{comment.UserGetCommentByUserID}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},

		{Group: "/normalUser", Path: "/topic", Method: "GET", Handles: []gin.HandlerFunc{topic.NormalGetAllTopic}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/admin", Path: "/topic", Method: "GET", Handles: []gin.HandlerFunc{topic.AdminGetAllTopic}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/topic", Method: "POST", Handles: []gin.HandlerFunc{topic.AdminCreateTopic}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/topic", Method: "PUT", Handles: []gin.HandlerFunc{topic.AdminUpdateTopic}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/topic", Method: "DELETE", Handles: []gin.HandlerFunc{topic.AdminDeleteTopic}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
	}
}

func (p *CommentPlugin) Uninstall() {

}

func (p *CommentPlugin) Upgrade() {

}
