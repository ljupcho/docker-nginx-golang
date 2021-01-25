package main

import (
	"github.com/gin-gonic/gin"
	_ "app/modules/log" // 日志
	// _ "app/modules/schedule" // 定时任务
	"runtime"
	"app/config"
	"app/modules/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := initRouter()

	server.Run(router)
}