package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"app/config"
	"app/filters"
	"app/filters/auth"
	routeRegister "app/routes"
	"net/http"
	// proxy "github.com/chenhg5/gin-reverseproxy"
)

func initRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob(config.GetEnv().TemplatePath + "/*")

	if config.GetEnv().Debug {
		// performance tool
		pprof.Register(router)
	}

	router.Use(gin.Logger())

	router.Use(handleErrors())
	router.Use(filters.RegisterSession())
	router.Use(filters.RegisterCache())

	router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth"))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Route not found.",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Method not found.",
		})
	})

	routeRegister.RegisterApiRouter(router)

	// ReverseProxy
	// router.Use(proxy.ReverseProxy(map[string] string {
	// 	"localhost:10091" : "localhost:9000",
	// }))

	return router
}
