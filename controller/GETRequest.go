package controller

import (
	"Cerebral-Palsy-Detection-System/Database"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

func UserBaseInfor(c *gin.Context) {
	j := Database.GetUserinfo(c.PostForm("name"))
	c.JSON(200, j)
}
