package controller

import (
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetKinship(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("mySession")
	userid := session.Get("mySession")
	kinship, err := model.GetKinship(userid.(int))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, kinship)
}
