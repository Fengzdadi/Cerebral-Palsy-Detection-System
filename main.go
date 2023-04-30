package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r = CollectRoute(r)

	panic(r.Run(":8080"))
}
