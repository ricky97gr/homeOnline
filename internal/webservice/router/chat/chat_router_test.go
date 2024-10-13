package chat

import (
	"testing"
	"time"

	"github.com/ricky97gr/homeOnline/internal/webservice/router/manager"
)

func TestPluginRun(t *testing.T) {
	p := &ChatPlugin{
		BasePlugin: manager.BasePlugin{
			ExecPath:   "/root/goWorkspace/family/bin/log_service",
			PluginName: "log_Service",
			ListenPort: 10002,
		},
	}
	p.Run()

	time.Sleep(50 * time.Second)
}
