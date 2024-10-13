package main

import (
	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/webservice/database"
	"github.com/ricky97gr/homeOnline/internal/webservice/im_server/message"
	"github.com/ricky97gr/homeOnline/internal/webservice/router"
	"github.com/ricky97gr/homeOnline/pkg/bininfo"
	_ "net/http/pprof"
	"os"
)

func main() {
	newlog.InitLogger("", os.Stdout)
	newlog.Logger.Infof("%s", bininfo.String())
	newlog.Logger.Infof("Family Community System is started!\n")
	database.Start()
	go message.ReceiveMessage()
	router.Start()

}
