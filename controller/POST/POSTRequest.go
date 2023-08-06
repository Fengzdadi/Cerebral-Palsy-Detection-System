package POST

import (
	"Cerebral-Palsy-Detection-System/Algorithm"
	"Cerebral-Palsy-Detection-System/Database"
	"Cerebral-Palsy-Detection-System/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UserBaseInfo(c *gin.Context) {
	//var user model.User
	//fmt.Print(c.PostForm("Username"))
	//Database.GetUserinfo(c.PostForm("Username"), &user)
	//c.JSON(200, user)
	session := sessions.Default(c)
	Username := session.Get("mySession")
	var user model.User
	Database.GetUserinfo(Username.(string), &user)
	fmt.Print(user)
	c.JSON(200, user)
}

func StartDetection(c *gin.Context) {
	var res model.VideoResult
	Algorithm.StartAlgorithm(&res)
	c.JSON(200, res)
}

func StartDetectionTest(c *gin.Context) {
	c.JSON(200, model.VideoResult{
		VideoName:   11114,
		Userid:      1523,
		VideoPath:   "1",
		VideoRes:    "1",
		Probability: "0.97",
	})
}

func UserLogin(c *gin.Context) {
	Username := c.PostForm("Username")
	UserPassword := c.PostForm("Password")
	// fmt.Print(Username, UserPassword)
	switch Database.UserCheck(Username, UserPassword) {
	case 1:
		// login success with session
		var user model.User
		Database.GetUserinfo(Username, &user)
		session := sessions.Default(c)
		// cause the session can only store string, so we need to convert the struct to string
		//sessionValues := model.SessionData{
		//	Username: Username,
		//	Age:      user.Age,
		//	Gender:   user.Gender,
		//	Email:    user.Email,
		//	Phone:    user.Phone,
		//}
		// fmt.Print(sessionValues)
		session.Set("mySession", Username)
		err := session.Save()
		if err != nil {
			log.Print(err)
			return
		}
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
	// uploadDir = "../files/"
	uploadDir = "./VProcessing/"
	_, err = os.Stat(uploadDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			return
		}
	}

	fileId := strconv.FormatInt(time.Now().Unix(), 10)
	newVideoName := "input.mp4"
	// newVideoName := fileId + path.Ext(Video.Filename)
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
