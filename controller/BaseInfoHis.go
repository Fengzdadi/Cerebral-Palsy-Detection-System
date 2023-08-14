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
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, b)
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
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
}
