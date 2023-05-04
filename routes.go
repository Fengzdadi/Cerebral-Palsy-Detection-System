package main

import (
	"github.com/gin-gonic/gin"
	"github/Intelligent-scheduling-system/controller"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// GET
	r.GET("/Hello", controller.Hello)
	r.GET("/UserBaseInfo", controller.UserbaseInfor)
	// POST
	r.POST("/UserLogin", controller.UserLogin)
}
