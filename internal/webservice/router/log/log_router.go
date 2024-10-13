package log

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
)

type LogPlugin struct {
	manager.BasePlugin
}

func init() {
	p := &LogPlugin{
		BasePlugin: manager.BasePlugin{
			PluginName:   "日志服务",
			Version:      "0.0.1_base",
			Author:       "forgocode",
			Description:  "用于通过grpc保存日志到mongo",
			ExecPath:     "",
			PluginStatus: manager.Stopped,
			ListenPort:   10002,
		},
	}
	manager.RegisterPlugin(p)
}

func (p *LogPlugin) Name() string {
	return p.PluginName
}

func (p *LogPlugin) Run() (*exec.Cmd, error) {
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

func (p *LogPlugin) Router() []manager.RouterInfo {
	return []manager.RouterInfo{}
}

func (p *LogPlugin) Uninstall() {

}

func (p *LogPlugin) Upgrade() {

}
