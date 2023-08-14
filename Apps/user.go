package Apps

import (
	"Cerebral-Palsy-Detection-System/Apps/WsApi"
	"Cerebral-Palsy-Detection-System/WS/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserRegisterService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, WsApi.ErrorResponse(err))
		logging.Info(err)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserLoginService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		// 从数据库查找用户的id
		var userid uint
		userid = service.GetUserid(c.PostForm("user_name"))
		session := sessions.Default(c)
		session.Set("mySession", userid)
		err := session.Save()
		if err != nil {
			return
		}
		c.JSON(200, res)
	} else {
		c.JSON(400, WsApi.ErrorResponse(err))
		logging.Info(err)
	}
}

func UserUpdatePwd(c *gin.Context) {
	var userUpdatePwd service.UserUpdatePwdService
	if err := c.ShouldBind(&userUpdatePwd); err == nil {
		res := userUpdatePwd.Update()
		c.JSON(200, res)
	} else {
		c.JSON(400, WsApi.ErrorResponse(err))
		logging.Info(err)
	}
}
