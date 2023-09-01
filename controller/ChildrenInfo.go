package controller

import (
	"Cerebral-Palsy-Detection-System/Apps/WsApi"
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func AddChildInfo(c *gin.Context) {
	var child model.ChildrenInfo
	userid, _ := c.Get("userid")
	if floatUserID, ok := userid.(float64); ok {
		uintUserID := uint(floatUserID)
		if err := c.ShouldBind(&child); err == nil {
			res := child.AddChildInfo(uintUserID)
			c.JSON(200, res)
		} else {
			c.JSON(400, WsApi.ErrorResponse(err))
			logging.Info(err)
		}
	}

}

func GetChildInfo(c *gin.Context) {
	userid, _ := c.Get("userid")
	if floatUserID, ok := userid.(float64); ok {
		uintUserID := uint(floatUserID)
		ch, res := model.GetChildInfo(uintUserID)
		c.JSON(200, gin.H{
			"message": ch,
			"res":     res,
		})
	}
}
