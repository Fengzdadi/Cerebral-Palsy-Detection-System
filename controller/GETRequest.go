package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

func ReturnVideoRes(c *gin.Context) {
	videoFile, err := os.Open(".\\Vprocessing\\output.mp4")
	if err != nil {
		c.JSON(500, gin.H{"message": "{err}"})
		return
	}
	defer videoFile.Close()

	c.Header("Content-Type", "video/mp4")

	_, err = io.Copy(c.Writer, videoFile)
	if err != nil {
		c.JSON(500, gin.H{"message": "{err}"})
		return
	}

}
