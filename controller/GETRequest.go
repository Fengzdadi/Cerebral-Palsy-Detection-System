package controller

import (
	"Cerebral-Palsy-Detection-System/Database"
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

func UserBaseInfo(c *gin.Context) {
	var user model.User
	Database.GetUserinfo(c.PostForm("Username"), &user)
	c.JSON(200, user)
}
