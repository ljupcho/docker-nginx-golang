package routes

import (
	"github.com/gin-gonic/gin"
	"morningo/controllers"
	"morningo/filters/auth"
)

func RegisterApiRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		apiRouter.GET("/test/index", controllers.IndexApi)
	}

	api := router.Group("/api")
	api.GET("/index", controllers.IndexApi)
	api.GET("/user/:userId", controllers.GetUser)
	api.GET("/getUsers", controllers.GetUsers)
	api.GET("/groups/bulk", controllers.CreateGroups)
	api.GET("/users/bulk", controllers.CreateUsers)
	api.GET("/users/run-chan", controllers.CreateUsersWithChannels)
	api.GET("/users/run-gor", controllers.CreateUserGoroutines)
	api.POST("/user/create", controllers.CreateUser)	
	api.PUT("/user/:userId/update", controllers.UpdateUser)
	api.GET("/cookie/set/:userid", controllers.CookieSetExample)

	// cookie auth middleware
	api.Use(auth.Middleware(auth.CookieAuthDriverKey))
	{
		api.GET("/orm", controllers.OrmExample)
		api.GET("/store", controllers.StoreExample)
		api.GET("/db", controllers.DBExample)
		api.GET("/cookie/get", controllers.CookieGetExample)
	}

	jwtApi := router.Group("/api")
	jwtApi.GET("/jwt/set/:userid", controllers.JwtSetExample)

	// jwt auth middleware
	jwtApi.Use(auth.Middleware(auth.JwtAuthDriverKey))
	{
		jwtApi.GET("/jwt/get", controllers.JwtGetExample)
	}
}
