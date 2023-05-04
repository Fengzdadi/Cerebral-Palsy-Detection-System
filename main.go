package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	Init()

	r := gin.Default()

	r = CollectRoutes(r)

	panic(r.Run(":8080"))
}
