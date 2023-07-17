package logger

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
)

func SetupLogger() {
	logger := hertzzap.NewLogger()
	// 在程序结束时调用 logger.Sync() 来确保日志缓冲区中的所有条目都被写入
	defer logger.Sync()
	//hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)
}
