package logger

import (
	"fmt"
	"go.uber.org/zap"
	"hertz-starter-kit/pkg/config"
)

func SetupLogger() {
	var logger *zap.Logger
	var err error
	if config.Config.Server.RunMode == "debug" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(fmt.Errorf("failed to init logger: %w", err))
	}
	//logger := hertzzap.NewLogger()
	// 在程序结束时调用 logger.Sync() 来确保日志缓冲区中的所有条目都被写入
	defer logger.Sync()
	//hlog.SetLogger(logger)
	//hlog.SetLevel(hlog.LevelInfo)
	zap.ReplaceGlobals(logger)
}
