package controller

import (
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetTestHistory(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("mySession")
	userid := session.Get("mySession")
	TH, err := model.GetTestHistory(userid.(int))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, TH)
}
