package manager

import (
	"os/exec"
	"sync"

	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/plugin"
	"github.com/gin-gonic/gin"
)

type Plugin interface {
	Run() (*exec.Cmd, error)
	Uninstall()
	Upgrade()
	Name() string
	Router() []RouterInfo
	Convert() model.Plugin
	Status() int
	SetStatus(int)
}

type BasePlugin struct {
	PluginName   string `json:"pluginName" gorm:"column:pluginName"`
	Md5          string `json:"md5" gorm:"column:md5"`
	Version      string `json:"version" gorm:"column:version"`
	Author       string `json:"author" gorm:"author"`
	Description  string `json:"description" gorm:"description"`
	PluginStatus int    `json:"pluginStatus" gorm:"column:pluginStatus"`
	ExecPath     string `json:"execPath"`
	ListenPort   int32
}

func (p *BasePlugin) Convert() model.Plugin {
	return model.Plugin{
		Name:        p.PluginName,
		Version:     p.Version,
		Author:      p.Author,
		Md5:         p.Md5,
		Description: p.Description,
		Status:      p.PluginStatus,
	}
}

func (p *BasePlugin) Status() int {
	return p.PluginStatus
}

func (p *BasePlugin) SetStatus(status int) {
	p.PluginStatus = status
}

type RouterInfo struct {
	Group      string
	Path       string
	Method     string
	Handles    []gin.HandlerFunc
	Middleware []gin.HandlerFunc
}

const (
	Running = iota + 1
	Stopped
	Upgrading
	Checking
)

func RegisterPlugin(p Plugin) {
	pluginManagerCenter.RegisterPlugin(p)

}

func StartPlugin(name string) {
	pluginManagerCenter.Load(name)
}

func StopPlugin(name string) {
	pluginManagerCenter.UnLoad(name)
}

func RegisterRouter(engine *gin.Engine) {
	// 注册插件管理的路由
	pluginManagerCenter.RegisterRouter(engine)

}

func (m *pluginManager) RegisterRouter(engine *gin.Engine) {
	pluginManagerCenter.mu.Lock()
	defer pluginManagerCenter.mu.Unlock()
	plugins := plugin.ListAllPlugins()
	for _, p := range plugins {
		for k, v := range pluginManagerCenter.items {
			if k.Name() == p.Name {
				v.status = p.Status
				pluginManagerCenter.items[k] = v
			}
		}
	}

	for p, s := range pluginManagerCenter.items {
		plugin.CreatePlugin(p.Convert())

		for _, r := range p.Router() {
			switch r.Method {
			case "GET":
				engine.Group(r.Group).Use(r.Middleware...).GET(r.Path, r.Handles...)
			case "POST":
				engine.Group(r.Group).Use(r.Middleware...).POST(r.Path, r.Handles...)
			case "PUT":
				engine.Group(r.Group).Use(r.Middleware...).PUT(r.Path, r.Handles...)
			case "DELETE":
				engine.Group(r.Group).Use(r.Middleware...).DELETE(r.Path, r.Handles...)
			}
			if s.status != Running {
				routerManagerCenter.items.Store(r.Group+r.Path, false)
				continue
			}
			routerManagerCenter.items.Store(r.Group+r.Path, true)
		}
	}
}

func ISRouterPass(path string) bool {
	v, ok := routerManagerCenter.items.Load(path)
	if !ok {
		return false
	}
	return v.(bool)
}

type pluginManager struct {
	mu    sync.RWMutex
	items map[Plugin]PluginInfo
}

type RouterManager struct {
	items sync.Map
}

type PluginInfo struct {
	cmdInfo *exec.Cmd
	status  int
}

var pluginManagerCenter = &pluginManager{
	mu:    sync.RWMutex{},
	items: make(map[Plugin]PluginInfo),
}

var routerManagerCenter = &RouterManager{
	items: sync.Map{},
}

func (m *pluginManager) RegisterPlugin(p Plugin) {
	m.mu.Lock()
	defer m.mu.Unlock()
	info := PluginInfo{
		cmdInfo: nil,
		status:  p.Status(),
	}

	m.items[p] = info
}

func (m *pluginManager) Load(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.items {
		if k.Name() != name {
			continue
		}
		cmd, err := k.Run()
		if err != nil {
			//return err
		}

		v.cmdInfo = cmd
		v.status = Running

		for _, r := range k.Router() {
			routerManagerCenter.items.Store(r.Group+r.Path, true)
		}
		m.items[k] = v
	}
	return nil
}

func (m *pluginManager) UnLoad(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.items {
		if k.Name() != name {
			continue
		}
		if v.status != Running {
			continue
		}
		for _, r := range k.Router() {
			routerManagerCenter.items.Store(r.Group+r.Path, false)
		}
		continue
		//TODO: ?可以取消吗
		return v.cmdInfo.Cancel()
	}
	return nil
}
