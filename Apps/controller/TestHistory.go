package controller

import (
	"Cerebral-Palsy-Detection-System/Serializer"
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AddTestHistory function for add message of test history
func AddTestHistory(c *gin.Context) {
	//session := sessions.Default(c)
	//userid := session.Get("mySession")
	var th model.TestHistory
	if err := c.Bind(&th); err == nil {
		res := th.AddTestHistory()
		c.JSON(200, gin.H{
			"res": res,
		})
	} else {
		c.JSON(400, gin.H{
			"res": Serializer.ErrorResponse(err),
		})
	}
}

func GetTestHistory(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get("mySession")
	th, res := model.GetTestHistory(userid.(uint))
	c.JSON(200, gin.H{
		"res": res,
		"th":  th,
	})
}

func GetTestHistoryYear(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("mySession")
	userid := session.Get("mySession")
	thys := model.GetTestHisYear(userid.(uint))
	c.JSON(200, gin.H{
		"thys": thys,
	})
}
