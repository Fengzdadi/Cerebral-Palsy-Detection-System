package Api

import (
	"Cerebral-Palsy-Detection-System/WS/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserRegisterService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {

		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}
