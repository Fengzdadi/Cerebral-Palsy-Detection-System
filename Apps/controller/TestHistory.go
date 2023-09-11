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
	c.JSON(200, gin.H{
		"message": TH,
		"error":   err,
	})
}

// AddTestHistory function for add message of test history
func AddTestHistory() {}
