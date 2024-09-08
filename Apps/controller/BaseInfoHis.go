package controller

import (
	"Cerebral-Palsy-Detection-System/Serializer"
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-gonic/gin"
)

func GetBaseInfoHis(c *gin.Context) {
	//session := sessions.Default(c)
	value, _ := c.Get("user_name")
	userid := model.GetUserid(value.(string))
	//session.Get("mySession")
	//userid := session.Get("mySession")
	b, err := model.GetBaseInfoHis(userid)
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
	value, _ := c.Get("user_name")
	userid := model.GetUserid(value.(string))
	base.BelongToChildID = userid
	err := base.AddBaseInfoHis()
	c.JSON(200, gin.H{
		"res": err,
	})
}
