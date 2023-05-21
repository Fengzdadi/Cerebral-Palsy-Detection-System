package main

import (
	"Cerebral-Palsy-Detection-System/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// GET
	r.GET("/Hello", controller.Hello)
	// POST
	r.POST("/UserBaseInfo", controller.UserBaseInfo)
	r.POST("/UserLogin", controller.UserLogin)
	// r.POST("/UserRegister", controller.UserRegister)
	r.POST("/UploadVideo", controller.VideoUpload)
	r.POST("/StartDetection", controller.StartDetection)
	return r
}
