package main

import (
	"Cerebral-Palsy-Detection-System/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// GET
	r.GET("/Hello", controller.Hello)
	// Return video which can show in the front-end
	r.GET("/VideoRes", controller.ReturnVideoRes)

	// Database part
	// For someone can get the history of result from database
	// The return is array
	r.GET("/UserHisResult", controller.GetHisResult)

	// POST
	// User part
	r.POST("/UserBaseInfo", controller.UserBaseInfo)
	r.POST("/UserLogin", controller.UserLogin)
	// r.POST("/UserRegister", controller.UserRegister)

	// Video part
	r.POST("/UploadVideo", controller.VideoUpload)
	r.POST("/StartDetection", controller.StartDetection)

	// Database part
	return r
}
