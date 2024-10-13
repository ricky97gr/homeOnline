package newlog

import (
	"go.uber.org/zap/zapcore"

	"github.com/ricky97gr/homeOnline/pkg/log"
)

var Logger *log.NewLogger

func InitLogger(fileName string, w zapcore.WriteSyncer) {
	opt := log.Options{
		FileName:   fileName,
		Level:      "info",
		ModuleName: "",
		W:          w,
	}
	Logger = log.New(opt)
}
