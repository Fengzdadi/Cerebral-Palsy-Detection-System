package controller

import (
	"Cerebral-Palsy-Detection-System/Serializer"
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetBaseInfoHis(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("mySession")
	userid := session.Get("mySession")
	b, err := model.GetBaseInfoHis(userid.(uint))
	c.JSON(200, gin.H{
		"baseInfo": b,
		"res":      err,
	})
}

func AddBaseInfoHis(c *gin.Context) {
	var base model.BaseInfoHis
	if err := c.ShouldBind(&base); err != nil {
		c.JSON(400, gin.H{
			"res": Serializer.ErrorResponse(err),
		})
		return
	}
	session := sessions.Default(c)
	session.Get("mySession")
	userid := session.Get("mySession")
	base.BelongToChildID = userid.(uint)
	err := base.AddBaseInfoHis()
	c.JSON(200, gin.H{
		"res": err,
	})
}
