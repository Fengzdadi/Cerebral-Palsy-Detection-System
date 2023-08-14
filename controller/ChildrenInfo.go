package controller

import (
	"Cerebral-Palsy-Detection-System/Apps/WsApi"
	"Cerebral-Palsy-Detection-System/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func AddChildInfo(c *gin.Context) {
	var child model.ChildrenInfo
	if err := c.ShouldBind(&child); err == nil {
		res := child.AddChildInfo()
		c.JSON(200, res)
	} else {
		c.JSON(400, WsApi.ErrorResponse(err))
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
		"message": ch,
		"error":   err,
	})
}
