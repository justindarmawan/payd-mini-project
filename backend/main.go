package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/justindarmawan/payd-mini-project/backend/cmd"
	"github.com/justindarmawan/payd-mini-project/backend/config"
	"github.com/justindarmawan/payd-mini-project/backend/database"

	_ "github.com/justindarmawan/payd-mini-project/backend/docs" // swagger generated files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.InitDB()
	config.InitConfig()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	cmd.RegisterRoutes(r)
	r.Run(":8080")
}
