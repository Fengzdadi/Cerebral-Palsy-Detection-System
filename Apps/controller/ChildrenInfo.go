package controller

import (
	"Cerebral-Palsy-Detection-System/Serializer"
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func AddChildInfo(c *gin.Context) {
	value, _ := c.Get("user_name")
	userid := model.GetUserid(value.(string))
	var child model.ChildrenInfo
	if err := c.ShouldBind(&child); err == nil {
		child.BelongTo = userid
		res := child.AddChildInfo()
		c.JSON(200, gin.H{
			"res": res,
		})
	} else {
		c.JSON(400, gin.H{
			"res": Serializer.ErrorResponse(err),
		})
		logging.Info(err)
	}
}

func GetChildInfo(c *gin.Context) {
	value, _ := c.Get("user_name")
	userid := model.GetUserid(value.(string))
	ch, err := model.GetChildInfo(userid)
	c.JSON(200, gin.H{
		"childInfo": ch,
		"res":       err,
	})
}
