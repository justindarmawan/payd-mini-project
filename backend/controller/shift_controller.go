package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	shift := r.Group("/api")
	{
		shift.GET("/health", Health)
	}
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "health")
}
