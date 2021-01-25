package routes

import (
	"github.com/gin-gonic/gin"
	"app/controllers"
	"app/filters/auth"
	"app/middlewares"
)

func RegisterApiRouter(router *gin.Engine) {
	sr := router.Group("/", auth.EnableCookieSession())
    {
        sr.GET("/login", controllers.Login)
        sr.POST("/register", controllers.Register)
        sr.GET("/logout", controllers.Logout)

        authorized := sr.Group("/auth", middlewares.AuthSessionMiddle())
        {
            authorized.GET("/me", controllers.Me)
        }
    }


	api := router.Group("/api")
	api.GET("/test", controllers.Test)
	api.GET("/login", controllers.Login)
	// cookie auth middleware
	api.Use(auth.Middleware(auth.CookieAuthDriverKey))
	{
		
		// api.GET("/orm", controllers.OrmExample)
		// api.GET("/store", controllers.StoreExample)
		// api.GET("/db", controllers.DBExample)
		// api.GET("/cookie/get", controllers.CookieGetExample)
	}


}
