package controllers

import (
	"github.com/gin-gonic/gin"
	"app/filters/auth"
	m "app/models"
    mlog "app/modules/log"
    "net/http"
)

func Test(c *gin.Context) {
    c.String(200, "!!!testing!!!")
}

func Login(c *gin.Context) {
    mlog.Info(mlog.E{Info: mlog.M{"data": "in login",},})
    email := c.Request.FormValue("email")
    password := c.Request.FormValue("password")
    mlog.Info(mlog.E{Info: mlog.M{"email": email,},})
    if hasSession := auth.HasSession(c); hasSession == true {
        c.String(200, "User logged in.")
        return
    }

    mlog.Info(mlog.E{Info: mlog.M{"here0": "in login",},})
    user := m.UserByEmail(email)
    
    if err := auth.ComparePassword(user.Password, password); err != nil {
        c.String(401, "Unauthenticated.")
        return
    }
    mlog.Info(mlog.E{Info: mlog.M{"here1": "in login",},})
    auth.SaveAuthSession(c, user.ID)

    c.String(200, "User logged in.")
}

func Logout(c *gin.Context) {
    if hasSession := auth.HasSession(c); hasSession == false {
        c.String(401, "Unauthenticated.")
        return
    }
    auth.ClearAuthSession(c)
    c.String(200, "User logged out successfully.")
}

func Register(c *gin.Context) {
    mlog.Info(mlog.E{Info: mlog.M{"data": "in register",},})
    var user m.User
    user.FirstName = c.Request.FormValue("first_name")
    user.LastName = c.Request.FormValue("last_name")
    user.Email = c.Request.FormValue("email")

    if hasSession := auth.HasSession(c); hasSession == true {
        c.String(200, "User logged in.")
        return
    }

    if existUser := m.UserByEmail(user.Email); existUser.ID != 0 {
        c.String(200, "Email already exists.")
        return
    }

    if c.Request.FormValue("password") != c.Request.FormValue("password_confirmation") {
        c.String(200, "Mismatch password!")
        return
    }

    if pwd, err := auth.EncryptPassword(c.Request.FormValue("password")); err == nil {
        user.Password = pwd
    }

    m.AddUser(&user)

    auth.SaveAuthSession(c, user.ID)

    c.String(200, "Registration successful.")
}


func Me(c *gin.Context) {
    currentUser := c.MustGet("userId").(uint)
    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "data": currentUser,
    })
}