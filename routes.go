package main

import (
	"Cerebral-Palsy-Detection-System/Utils"
	"Cerebral-Palsy-Detection-System/WS/service"
	"Cerebral-Palsy-Detection-System/controller"
	"Cerebral-Palsy-Detection-System/controller/GET"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// GET
	r.GET("/Hello", GET.Hello)
	// Return video which can show in the front-end
	r.GET("/VideoRes", GET.ReturnVideoRes)

	// Database part
	// For someone can get the history of result from database
	// The return is array
	// r.GET("/UserHisResult", GET.GetHisResult)

	//Ws part
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}) // For test
	r.GET("/ws", service.WsHandler)

	// POST
	// User part
	// r.POST("/UserBaseInfo", POST.UserBaseInfo)
	// r.POST("/UserLogin", POST.UserLogin)
	// r.POST("/UserRegister", controller.UserRegister)

	// Video part
	// r.POST("/UploadVideo", POST.VideoUpload)
	// r.POST("/StartDetection", POST.StartDetection)
	// r.POST("/StartDetectionTest", POST.StartDetectionTest)

	// Database part

	// New Request Part
	// User part
	r.POST("/UserRegister", controller.UserRegister)
	r.POST("/UserLogin", controller.UserLogin)
	r.POST("/UserUpdatePwd", Utils.AuthMiddleware(), controller.UserUpdatePwd)

	// BaseInfo part
	r.GET("/GetBaseInfoHis", Utils.AuthMiddleware(), controller.GetBaseInfoHis)
	r.POST("/AddBaseInfoHis", Utils.AuthMiddleware(), controller.AddBaseInfoHis)

	//ChildInfo part
	r.GET("/GetChildInfo", Utils.AuthMiddleware(), controller.GetChildInfo)
	r.POST("/AddChildInfo", Utils.AuthMiddleware(), controller.AddChildInfo)

	//Kinship part
	r.GET("/GetKinship", Utils.AuthMiddleware(), controller.GetKinship)

	//TestHistory part
	r.GET("/GetTestHistory", Utils.AuthMiddleware(), controller.GetTestHistory)

	return r
}
