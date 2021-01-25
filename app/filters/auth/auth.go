package auth

import (
	"github.com/gin-gonic/gin"
	"app/filters/auth/drivers"
	"net/http"
	"github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
)

const (
	CookieAuthDriverKey = "cookie"
)

var driverList = map[string]Auth{
	CookieAuthDriverKey: drivers.NewCookieAuthDriver(),
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{}
	Logout(http *http.Request, w http.ResponseWriter) bool
}

func RegisterGlobalAuthDriver(authKey string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, GenerateAuthDriver(authKey))
		c.Next()
	}
}

func Middleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !GenerateAuthDriver(authKey).Check(c) {
			c.HTML(http.StatusOK, "index.tpl", gin.H{
				"title": "login first",
			})
			c.Abort()
		}
		c.Next()
	}
}

func GenerateAuthDriver(string string) Auth {
	return driverList[string]
}

func GetCurUser(c *gin.Context, key string) map[string]interface{} {
	authDriver, _ := c.MustGet(key).(Auth)
	return authDriver.User(c).(map[string]interface{})
}

const KEY = "AEN233"

func EnableCookieSession() gin.HandlerFunc {
    store := cookie.NewStore([]byte(KEY))
    return sessions.Sessions("SAMPLE", store)
}
