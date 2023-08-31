package controller

import (
	"Cerebral-Palsy-Detection-System/Apps"
	"Cerebral-Palsy-Detection-System/Apps/WsApi"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	Apps.UserRegister(c)
}

func UserLogin(c *gin.Context) {
	Apps.UserLogin(c)
}

func UserUpdatePwd(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("mySession")
	if username == c.PostForm("Username") {
		Apps.UserUpdatePwd(c)
	} else {
		c.JSON(400, WsApi.MyErrorResponse("You are not authorized to do this operation"))
	}
}
