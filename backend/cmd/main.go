package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justindarmawan/payd-mini-project/backend/controller"
)

func main() {
	r := gin.Default()
	controller.RegisterRoutes(r)

	r.Run(":8080")
}
