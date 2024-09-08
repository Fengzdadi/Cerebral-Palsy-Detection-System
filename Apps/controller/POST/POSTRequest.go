package POST

import (
	"Cerebral-Palsy-Detection-System/Algorithm"
	"Cerebral-Palsy-Detection-System/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func StartFiveSecond(c *gin.Context) {
	fmt.Println("start five second")
	res, err := Algorithm.PullVideo()
	if err != nil {
		log.Println("err", err)
	}
	c.JSON(200, res)
}

func StartDetection(c *gin.Context) {
	var res model.VideoResult
	Algorithm.StartAlgorithm(&res)
	c.JSON(200, res)
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

	//fileId := strconv.FormatInt(time.Now().Unix(), 10)
	newVideoName := "input.mp4"
	// newVideoName := fileId + path.Ext(Video.Filename)
	dst := uploadDir + newVideoName
	uploadErr := c.SaveUploadedFile(Video, dst)
	if uploadErr != nil {
		c.String(http.StatusInternalServerError, "上传失败："+uploadErr.Error())
		log.Fatal(uploadErr)
		return
	}

	var res model.VideoResult
	Algorithm.StartAlgorithm(&res)
	//resJson := gin.H{"state": "successful", "filePath": dst, "videoName": Video.Filename, "fileId": fileId}
	c.JSON(200, res)

	return
}
