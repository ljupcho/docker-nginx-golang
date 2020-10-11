package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"morningo/config"
	"morningo/filters"
	"morningo/filters/auth"
	routeRegister "morningo/routes"
	"net/http"
	// proxy "github.com/chenhg5/gin-reverseproxy"
)

func initRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob(config.GetEnv().TemplatePath + "/*") // html模板

	if config.GetEnv().Debug {
		pprof.Register(router) // 性能分析工具
	}

	router.Use(gin.Logger())

	router.Use(handleErrors())            // 错误处理
	router.Use(filters.RegisterSession()) // 全局session
	router.Use(filters.RegisterCache())   // 全局cache

	router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth")) // 全局auth cookie
	router.Use(auth.RegisterGlobalAuthDriver("jwt", "jwt_auth"))    // 全局auth jwt

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该方法",
		})
	})

	routeRegister.RegisterApiRouter(router)

	// ReverseProxy
	// router.Use(proxy.ReverseProxy(map[string] string {
	// 	"localhost:4000" : "localhost:9090",
	// }))

	return router
}
