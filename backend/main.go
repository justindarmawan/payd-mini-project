package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/justindarmawan/payd-mini-project/backend/cmd"
	"github.com/justindarmawan/payd-mini-project/backend/config"
	"github.com/justindarmawan/payd-mini-project/backend/database"

	_ "github.com/justindarmawan/payd-mini-project/backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.InitDB()
	config.InitConfig()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	cmd.RegisterRoutes(r)
	r.Run(":8080")
}
