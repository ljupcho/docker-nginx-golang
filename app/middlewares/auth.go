package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "net/http"
)

func AuthSessionMiddle() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        sessionValue := session.Get("userId")
        if sessionValue == nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Unauthorized",
            })
            c.Abort()
            return
        }

        c.Set("userId", sessionValue.(uint))

        c.Next()
        return
    }
}