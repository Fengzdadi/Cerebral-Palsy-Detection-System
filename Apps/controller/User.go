package controller

import (
	"Cerebral-Palsy-Detection-System/Algorithm"
	"Cerebral-Palsy-Detection-System/Serializer"
	"Cerebral-Palsy-Detection-System/model"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegister model.UserRegisterService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, gin.H{
			"res": res,
		})
	} else {
		c.JSON(400, gin.H{
			"res": Serializer.ErrorResponse(err),
		})
		logging.Info(err)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin model.UserLoginService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		// 从数据库查找用户的id
		if res.Msg == "ok" {
			var username = c.PostForm("user_name")

			//userid = model.GetUserid(username)
			//session := sessions.Default(c)
			//session.Options(sessions.Options{
			//	SameSite: http.SameSiteNoneMode,
			//	Secure:   true,
			//})
			//session.Set("mySession", userid)
			//err := session.Save()
			//if err != nil {
			//	return
			//}

			token, _ := Algorithm.GenerateToken(username)
			c.JSON(http.StatusOK, gin.H{
				"res":   res,
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"res": res,
			})
		}
	} else {
		c.JSON(200, gin.H{
			"res": Serializer.ErrorResponse(err),
		})
	}
}

func UserUpdatePwd(c *gin.Context) {
	username := c.PostForm("user_name")
	value, _ := c.Get("user_name")
	if username == value.(string) {
		var userUpdatePwd model.UserUpdatePwdService
		if err := c.ShouldBind(&userUpdatePwd); err == nil {
			res := userUpdatePwd.Update()
			c.JSON(200, gin.H{
				"res": res,
			})
		} else {
			c.JSON(400, gin.H{
				"res": Serializer.ErrorResponse(err),
			})
			logging.Info(err)
		}
	} else {
		c.JSON(400, gin.H{
			"res": Serializer.MyErrorResponse("You are not authorized to do this operation"),
		})
	}
}
