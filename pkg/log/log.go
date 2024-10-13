package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type NewLogger struct {
	fileName string
	Level    string
	*zap.SugaredLogger
	zapLevel   zap.AtomicLevel
	moduleName string
	writer     zapcore.WriteSyncer
}

type Options struct {
	FileName   string
	Level      string
	ModuleName string
	W          zapcore.WriteSyncer
}

func New(opts ...Options) *NewLogger {
	logger := &NewLogger{}
	for _, opt := range opts {
		logger.fileName = opt.FileName
		logger.Level = opt.Level
		logger.moduleName = opt.ModuleName
		logger.writer = opt.W
	}
	if len(opts) == 0 {
		logger.defaultOptions()
	}
	logger.zapLevel = getLevelEnabler()
	logger.SetLevel(logger.Level)
	logger.SugaredLogger = initLogger(logger.fileName, logger.zapLevel, logger.moduleName, logger.writer)
	return logger
}

func (l *NewLogger) SetLevel(level string) {
	l.Level = level
	switch level {
	case "info":
		l.zapLevel.SetLevel(zapcore.InfoLevel)
	case "debug":
		l.zapLevel.SetLevel(zapcore.DebugLevel)
	}
}

func (l *NewLogger) defaultOptions() {
	l.fileName = "running.log"
	l.Level = "info"
}

func initLogger(fileName string, level zapcore.LevelEnabler, moduleName string, w zapcore.WriteSyncer) *zap.SugaredLogger {
	writerSyncer := getLogWriter(fileName)
	if w != nil {
		writerSyncer = w
	}
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, level)
	loger := zap.New(core, zap.AddCaller())
	return loger.Sugar().Named(moduleName)
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	logConf := zap.NewProductionEncoderConfig()
	logConf.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	logConf.EncodeLevel = zapcore.CapitalLevelEncoder
	//logConf.NameKey = "module"
	logConf.SkipLineEnding = true
	return zapcore.NewConsoleEncoder(logConf)
}

func getLevelEnabler() zap.AtomicLevel {
	return zap.NewAtomicLevel()
}
