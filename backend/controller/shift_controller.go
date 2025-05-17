package controller

import (
	"net/http"
	"strconv"

	"github.com/justindarmawan/payd-mini-project/backend/service"
	"github.com/justindarmawan/payd-mini-project/backend/utils"
	"github.com/justindarmawan/payd-mini-project/backend/viewmodel"

	"github.com/gin-gonic/gin"
)

type ShiftController struct {
	service *service.ShiftService
}

func NewShiftController(service *service.ShiftService) *ShiftController {
	return &ShiftController{service: service}
}

// worker
func (controller *ShiftController) GetAvailableShifts(c *gin.Context) {
	shifts, err := controller.service.GetAvailableShifts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shifts)
}

func (controller *ShiftController) GetAssignedShifts(c *gin.Context) {
	workerID, err := strconv.Atoi(c.Param("worker_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shifts, err := controller.service.GetAssignedShifts(workerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shifts)
}

func (controller *ShiftController) CreateShiftRequest(c *gin.Context) {
	var input viewmodel.ShiftRequestReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := controller.service.CreateShiftRequest(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (controller *ShiftController) GetRequestsByWorker(c *gin.Context) {
	workerID, err := strconv.Atoi(c.Param("worker_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requests, err := controller.service.GetRequestsByWorker(workerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, requests)
}

// admin
func (controller *ShiftController) CreateShift(c *gin.Context) {
	var input viewmodel.ShiftReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateCreateShiftReq(input); len(err) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "errors": err})
		return
	}
	if _, err := controller.service.CreateShift(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create shift"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (controller *ShiftController) EditShift(c *gin.Context) {
	var input viewmodel.ShiftReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shiftID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateEditShiftReq(input); len(err) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "errors": err})
		return
	}
	if err := controller.service.EditShift(input, shiftID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (controller *ShiftController) DeleteShift(c *gin.Context) {
	shiftID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.service.DeleteShift(shiftID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (controller *ShiftController) GetPendingRequests(c *gin.Context) {
	reqs, err := controller.service.GetPendingRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reqs)
}

func (controller *ShiftController) ApproveShiftRequest(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = controller.service.ApproveShiftRequest(requestID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (controller *ShiftController) RejectShiftRequest(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = controller.service.RejectShiftRequest(requestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
