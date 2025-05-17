package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justindarmawan/payd-mini-project/backend/cmd"
	"github.com/justindarmawan/payd-mini-project/backend/config"
	"github.com/justindarmawan/payd-mini-project/backend/database"
)

func main() {
	database.InitDB()
	config.InitConfig()

	r := gin.Default()
	cmd.RegisterRoutes(r)

	r.Run(":8080")
}
