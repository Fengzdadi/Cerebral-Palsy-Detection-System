package controller

import (
	"Cerebral-Palsy-Detection-System/Serializer"
	"Cerebral-Palsy-Detection-System/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func AddChildInfo(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("mySession")
	userid := session.Get("mySession")
	var child model.ChildrenInfo
	if err := c.ShouldBind(&child); err == nil {
		child.BelongTo = userid.(uint)
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
	session := sessions.Default(c)
	userid := session.Get("mySession")
	fmt.Print(userid)
	id := userid.(uint)
	ch, err := model.GetChildInfo(id)
	c.JSON(200, gin.H{
		"childInfo": ch,
		"res":       err,
	})
}
