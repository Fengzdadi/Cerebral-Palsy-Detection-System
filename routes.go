package main

import (
	"Cerebral-Palsy-Detection-System/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// GET
	r.GET("/Hello", controller.Hello)
	r.GET("/UserBaseInfo", controller.UserbaseInfor)
	// POST
	r.POST("/UserLogin", controller.UserLogin)

	return r
}
