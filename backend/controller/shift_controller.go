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
// @Summary Get Available Shifts
// @Description Returns all available shifts
// @Tags worker
// @Accept json
// @Produce json
// @Success 200 {array} viewmodel.ShiftRes
// @Failure 500 {object} map[string]string
// @Router /api/shifts/available [get]
func (controller *ShiftController) GetAvailableShifts(c *gin.Context) {
	shifts, err := controller.service.GetAvailableShifts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shifts)
}

// @Summary Get Assigned Shifts
// @Description Returns shifts assigned to a specific worker
// @Tags worker
// @Accept json
// @Produce json
// @Param worker_id path int true "Worker ID"
// @Success 200 {array} viewmodel.ShiftRes
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/shifts/assigned/{worker_id} [get]
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

// @Summary Create Shift Request
// @Description Submit a shift request
// @Tags worker
// @Accept json
// @Produce json
// @Param request body viewmodel.ShiftRequestReq true "Shift Request Payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/request [post]
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

// @Summary Get Worker Requests
// @Description Returns all shift requests made by a specific worker
// @Tags worker
// @Accept json
// @Produce json
// @Param worker_id path int true "Worker ID"
// @Success 200 {array} viewmodel.RequestRes
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/requests/{worker_id} [get]
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
// @Summary Create Shift
// @Description Create a new shift
// @Tags admin
// @Accept json
// @Produce json
// @Param shift body viewmodel.ShiftReq true "Shift Payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/shifts [post]
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

// @Summary Edit Shift
// @Description Edit an existing shift
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Shift ID"
// @Param shift body viewmodel.ShiftReq true "Shift Payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/shifts/{id} [put]
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

// @Summary Delete Shift
// @Description Delete an existing shift
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Shift ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/shifts/{id} [delete]
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

// @Summary Get Pending Requests
// @Description Returns all pending shift requests
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {array} viewmodel.RequestRes
// @Failure 500 {object} map[string]string
// @Router /api/requests/pending [get]
func (controller *ShiftController) GetPendingRequests(c *gin.Context) {
	reqs, err := controller.service.GetPendingRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reqs)
}

// @Summary Approve Shift Request
// @Description Approve a shift request by ID
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Request ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /api/requests/{id}/approve [post]
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

// @Summary Reject Shift Request
// @Description Reject a shift request by ID
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Request ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/requests/{id}/reject [post]
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
