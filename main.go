package main

import (
	"Cerebral-Palsy-Detection-System/middleware"
	_ "Cerebral-Palsy-Detection-System/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	Init()

	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	r.Use(middleware.Cors())

	r = CollectRoutes(r)

	panic(r.Run(":8080"))
}
