// This document is the entry of the project.
package main

import (
	"Cerebral-Palsy-Detection-System/middleware"
	_ "Cerebral-Palsy-Detection-System/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	// Use cookie-based sessions
	store := cookie.NewStore([]byte("loginUser"))
	r.Use(sessions.Sessions("session", store))

	r = CollectRoutes(r)

	panic(r.Run(":8080"))
}
