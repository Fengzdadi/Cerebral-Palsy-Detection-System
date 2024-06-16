package controller

import (
	"Cerebral-Palsy-Detection-System/Serializer"
	"Cerebral-Palsy-Detection-System/model"
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
	value, _ := c.Get("user_name")
	userid := model.GetUserid(value.(string))
	th, res := model.GetTestHistory(userid)
	c.JSON(200, gin.H{
		"res": res,
		"th":  th,
	})
}

func GetTestHistoryYear(c *gin.Context) {
	value, _ := c.Get("user_name")
	userid := model.GetUserid(value.(string))
	thys, res := model.GetTestHisYear(userid)
	c.JSON(200, gin.H{
		"res":  res,
		"thys": thys,
	})
}
