package Apps

import (
	"Cerebral-Palsy-Detection-System/Apps/controller"
	"Cerebral-Palsy-Detection-System/Apps/controller/GET"
	"Cerebral-Palsy-Detection-System/Apps/controller/WS/service"
	"Cerebral-Palsy-Detection-System/Apps/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes() {
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
	r.POST("/register", controller.UserRegister)
	r.POST("/login", controller.UserLogin)

	user := r.Group("user", middleware.AuthMiddleware())
	user.POST("/updatePwd", controller.UserUpdatePwd)

	// BaseInfo part
	user.GET("/getBaseInfoHis", controller.GetBaseInfoHis)
	user.POST("/addBaseInfoHis", controller.AddBaseInfoHis)

	//ChildInfo part
	user.GET("/getChildInfo", controller.GetChildInfo)
	user.POST("/addChildInfo", controller.AddChildInfo)

	//Kinship part
	user.GET("/getKinship", controller.GetKinship)

	//TestHistory part
	user.GET("/getTestHistory", controller.GetTestHistory)

}
