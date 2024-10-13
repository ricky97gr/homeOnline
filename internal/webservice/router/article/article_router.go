package article

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/webservice/controller/article"
	"github.com/ricky97gr/homeOnline/internal/webservice/controller/category"
	"github.com/ricky97gr/homeOnline/internal/webservice/controller/tag"
	"github.com/ricky97gr/homeOnline/internal/webservice/middleware"

	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
	"github.com/gin-gonic/gin"
)

type ArticlePlugin struct {
	manager.BasePlugin
}

func init() {
	p := &ArticlePlugin{
		BasePlugin: manager.BasePlugin{
			PluginName:   "文章服务",
			Version:      "0.0.1_base",
			Author:       "forgocode",
			Description:  "用于发布文章",
			ExecPath:     "",
			PluginStatus: manager.Stopped,
			ListenPort:   10002,
		},
	}
	manager.RegisterPlugin(p)
}

func (p *ArticlePlugin) Name() string {
	return p.PluginName
}

func (p *ArticlePlugin) Run() (*exec.Cmd, error) {
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

func (p *ArticlePlugin) Router() []manager.RouterInfo {
	return []manager.RouterInfo{
		{Group: "/admin", Path: "/article/publish", Method: "PUT", Handles: []gin.HandlerFunc{article.AdminPublishArticle}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/article/ban", Method: "PUT", Handles: []gin.HandlerFunc{article.AdminBanArticle}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/article/sendback", Method: "PUT", Handles: []gin.HandlerFunc{article.AdminSendBackArticle}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "", Path: "/article", Method: "GET", Handles: []gin.HandlerFunc{article.NormalGetArticle}},
		{Group: "", Path: "/article/:id", Method: "GET", Handles: []gin.HandlerFunc{article.NormalGetArticleInfo}},
		{Group: "/admin", Path: "/article", Method: "GET", Handles: []gin.HandlerFunc{article.AdminGetArticle}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/normalUser", Path: "/article", Method: "POST", Handles: []gin.HandlerFunc{article.CreateNewArticle}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/normalUser", Path: "/article", Method: "GET", Handles: []gin.HandlerFunc{article.NormalGetArticle}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/normalUser", Path: "/article/:id", Method: "GET", Handles: []gin.HandlerFunc{article.NormalGetArticleInfo}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/normalUser/articleList/:userID", Method: "GET", Handles: []gin.HandlerFunc{article.NormalGetArticleList}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},

		{Group: "/normalUser", Path: "/tags", Method: "GET", Handles: []gin.HandlerFunc{tag.NormalGetAllTag}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/admin", Path: "/tags", Method: "GET", Handles: []gin.HandlerFunc{tag.AdminGetAllTag}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/tags", Method: "POST", Handles: []gin.HandlerFunc{tag.AdminCreateTag}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/tags", Method: "DELETE", Handles: []gin.HandlerFunc{tag.AdminDeleteTag}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/tags", Method: "PUT", Handles: []gin.HandlerFunc{tag.AdminUpdateTag}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},

		{Group: "/normalUser", Path: "/category", Method: "GET", Handles: []gin.HandlerFunc{category.NormalGetAllCategory}, Middleware: []gin.HandlerFunc{middleware.AuthNormal()}},
		{Group: "/admin", Path: "/category", Method: "GET", Handles: []gin.HandlerFunc{category.AdminGetAllCategory}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/category", Method: "POST", Handles: []gin.HandlerFunc{category.AdminCreateCategory}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/category", Method: "DELETE", Handles: []gin.HandlerFunc{category.AdminDeleteCategory}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "/admin", Path: "/category", Method: "PUT", Handles: []gin.HandlerFunc{category.AdminUpdateCategory}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
	}
}

func (p *ArticlePlugin) Uninstall() {

}

func (p *ArticlePlugin) Upgrade() {

}
