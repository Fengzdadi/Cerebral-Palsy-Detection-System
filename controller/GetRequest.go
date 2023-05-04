package controller

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

func UserbaseInfor(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "UserbaseInfor, Success!",
	})
}

func UserLogin(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "UserLogin, Success!",
	})

}
