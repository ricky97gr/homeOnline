package main

import (
	"flag"
	"github.com/ricky97gr/homeOnline/internal/consumer"
	"os"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
)

var port uint64

func main() {
	flag.Uint64Var(&port, "port", 10000, "the port for cmd listen")
	flag.Parse()
	newlog.InitLogger("", os.Stdout)
	newlog.Logger.Infof("grpc server is start succfully, listen port: %d!\n", port)

	//数据库
	//从kafka读消息，或者作为一个grpc server接受其他模块的消息
	//启动

	consumer.Start()
	//logserver.Start(port)
	select {}
}

type Handler interface {
	AddLog()
	UpdateLog()
}
