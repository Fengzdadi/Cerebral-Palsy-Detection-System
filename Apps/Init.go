package Apps

import (
	"Cerebral-Palsy-Detection-System/Apps/middleware"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitWebFrameWork() {
	r = gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	r.Use(middleware.Cors())

	// Use cookie-based sessions
	//store := cookie.NewStore([]byte("loginUser"))
	//r.Use(sessions.Sessions("session", store))

	CollectRoutes()

}

func StartServer() {
	r.Run("10.249.79.59:8080")
}
