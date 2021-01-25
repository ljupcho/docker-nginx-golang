package auth

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    m "app/models"
)

func SaveAuthSession(c *gin.Context, id uint) {
    session := sessions.Default(c)
    session.Set("userId", id)
    session.Save()
}

func ClearAuthSession(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
}

func HasSession(c *gin.Context) bool {
    session := sessions.Default(c)
    if sessionValue := session.Get("userId"); sessionValue == nil {
        return false
    }
    return true
}

func GetSessionUserId(c *gin.Context) uint {
    session := sessions.Default(c)
    sessionValue := session.Get("userId")
    if sessionValue == nil {
        return 0
    }
    return sessionValue.(uint)
}

func GetEmailSession(c *gin.Context) map[string]interface{} {
    hasSession := HasSession(c)
    email := ""
    if hasSession {
        userId := GetSessionUserId(c)
        email = m.UserById(userId).Email
    }
    data := make(map[string]interface{})
    data["hasSession"] = hasSession
    data["email"] = email
    return data
}

func GetUserSession(c *gin.Context) m.User {
    userId := GetSessionUserId(c)
    return m.UserById(userId)
}


func EncryptPassword(source string) (string, error) {
    hashPwd, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
    return string(hashPwd), err
}

func ComparePassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

