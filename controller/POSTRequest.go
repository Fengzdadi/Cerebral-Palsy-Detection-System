package controller

import (
	"Cerebral-Palsy-Detection-System/Algorithm"
	"Cerebral-Palsy-Detection-System/Database"
	"Cerebral-Palsy-Detection-System/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func UserBaseInfo(c *gin.Context) {
	var user model.User
	fmt.Print(c.PostForm("Username"))
	Database.GetUserinfo(c.PostForm("Username"), &user)
	c.JSON(200, user)
}

func StartDetection(c *gin.Context) {
	var res model.Result
	Algorithm.StartAlgorithm(&res)
	c.JSON(200, res)
}

func UserLogin(c *gin.Context) {
	Username := c.PostForm("Username")
	UserPassword := c.PostForm("Password")
	switch Database.UserCheck(Username, UserPassword) {
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
