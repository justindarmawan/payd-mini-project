package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justindarmawan/payd-mini-project/backend/service"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{AuthService: service}
}

func (a *AuthController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, user, err := a.AuthService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "role": user.Role, "id": user.ID})
}
