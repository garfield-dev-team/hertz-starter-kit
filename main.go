// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"hertz-starter-kit/biz/middleware"
	"hertz-starter-kit/pkg/setting"
)

func main() {
	if err := setting.SetupSetting(); err != nil {
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}

	logger := hertzzap.NewLogger()
	// 在程序结束时调用 logger.Sync() 来确保日志缓冲区中的所有条目都被写入
	defer logger.Sync()
	//hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)

	h := server.Default(
		server.WithHostPorts(":8080"),
		server.WithBasePath("/api/"),
		server.WithHandleMethodNotAllowed(true),
	)
	h.Use(middleware.GlobalErrorHandler)
	register(h)
	h.Spin()
}
