package controller

import (
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetBaseInfoHis(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("mySession")
	userid := session.Get("mySession")
	b, err := model.GetBaseInfoHis(userid.(int))
	c.JSON(200, gin.H{
		"message": b,
		"error":   err,
	})
}

func AddBaseInfoHis(c *gin.Context) {
	var base model.BaseInfoHis
	if err := c.ShouldBindJSON(&base); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Get("mySession")
	userid := session.Get("mySession")
	base.BelongToChildID = userid.(int)
	err := base.AddBaseInfoHis()
	c.JSON(200, err)
}
