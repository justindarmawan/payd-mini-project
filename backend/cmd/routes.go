package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/justindarmawan/payd-mini-project/backend/controller"
	"github.com/justindarmawan/payd-mini-project/backend/database"
	"github.com/justindarmawan/payd-mini-project/backend/repository"
	"github.com/justindarmawan/payd-mini-project/backend/service"
	"github.com/justindarmawan/payd-mini-project/backend/utils"
)

func RegisterRoutes(r *gin.Engine) {
	db := database.GetDB()

	healthController := controller.NewHealthController()
	authController := controller.NewAuthController(service.NewAuthService(repository.NewUserRepository(db)))
	shiftController := controller.NewShiftController(service.NewShiftService(repository.NewShiftRepository(db)))

	health := r.Group("/api")
	{
		health.GET("/health", healthController.Health)
		health.POST("/health", healthController.Health)
	}

	r.POST("/login", authController.Login)

	admin := r.Group("/api")
	{
		admin.Use(utils.AuthMiddleware("admin"))
		admin.POST("/shifts", shiftController.CreateShift)
		admin.PUT("/shifts/:id", shiftController.EditShift)
		admin.DELETE("/shifts/:id", shiftController.DeleteShift)
		admin.GET("/requests/pending", shiftController.GetPendingRequests)
		admin.POST("/requests/:id/approve", shiftController.ApproveShiftRequest)
		admin.POST("/requests/:id/reject", shiftController.RejectShiftRequest)
	}

	worker := r.Group("/api")
	{
		worker.Use(utils.AuthMiddleware("worker"))
		worker.GET("/shifts/available", shiftController.GetAvailableShifts)
		worker.GET("/shifts/assigned/:worker_id", shiftController.GetAssignedShifts)
		worker.POST("/requests", shiftController.CreateShiftRequest)
		worker.GET("/requests/:worker_id", shiftController.GetRequestsByWorker)
	}
}
