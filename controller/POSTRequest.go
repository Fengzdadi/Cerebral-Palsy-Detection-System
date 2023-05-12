package controller

import (
	"Cerebral-Palsy-Detection-System/Database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func UserLogin(c *gin.Context) {
	UserID := c.PostForm("UserID")
	UserPassword := c.PostForm("UserPassword")
	switch Database.UserCheck(UserID, UserPassword) {
	case 1:
		c.JSON(200, gin.H{
			"message": "UserLogin, Success!",
		})
	case 0:
		c.JSON(http.StatusForbidden, gin.H{
			"message": "UserLogin, Failed!",
		})
	case -1:
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "UserLogin, No user match!",
		})
	}
}

func VideoUpload(c *gin.Context) {
	Video, err := c.FormFile("Video")
	if err != nil {
		c.String(http.StatusInternalServerError, "读取失败："+err.Error())
	}
	var uploadDir string
	uploadDir = "../files/"
	_, err = os.Stat(uploadDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			return
		}
	}

	fileId := strconv.FormatInt(time.Now().Unix(), 10)
	newVideoName := fileId + path.Ext(Video.Filename)
	dst := uploadDir + newVideoName
	uploadErr := c.SaveUploadedFile(Video, dst)
	if uploadErr != nil {
		c.String(http.StatusInternalServerError, "上传失败："+uploadErr.Error())
		log.Fatal(uploadErr)
		return
	}

	resJson := gin.H{"state": "successful", "filePath": dst, "videoName": Video.Filename, "fileId": fileId}
	c.JSON(200, resJson)

	return
}
