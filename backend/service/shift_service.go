package service

import (
	"github.com/jinzhu/copier"

	"github.com/justindarmawan/payd-mini-project/backend/model"
	"github.com/justindarmawan/payd-mini-project/backend/repository"
	"github.com/justindarmawan/payd-mini-project/backend/viewmodel"
)

type ShiftService struct {
	repo *repository.ShiftRepository
}

func NewShiftService(repo *repository.ShiftRepository) *ShiftService {
	return &ShiftService{repo: repo}
}

// worker
func (s *ShiftService) GetAvailableShifts() ([]viewmodel.ShiftRes, error) {
	var vm []viewmodel.ShiftRes
	shifts, err := s.repo.GetAvailableShifts()
	copier.Copy(&vm, &shifts)
	return vm, err
}

func (s *ShiftService) GetAssignedShifts(workerID int) ([]viewmodel.ShiftRes, error) {
	var vm []viewmodel.ShiftRes
	shifts, err := s.repo.GetAssignedShifts(workerID)
	copier.Copy(&vm, &shifts)
	return vm, err
}

func (s *ShiftService) CreateShiftRequest(input viewmodel.ShiftRequestReq) (int64, error) {
	var shiftRequest model.ShiftRequest
	copier.Copy(&shiftRequest, &input)
	return s.repo.CreateShiftRequest(shiftRequest)
}

func (s *ShiftService) GetRequestsByWorker(workerID int) ([]viewmodel.RequestRes, error) {
	var vm []viewmodel.RequestRes
	requests, err := s.repo.GetRequestsByWorker(workerID)
	copier.Copy(&vm, &requests)
	return vm, err
}

// admin
func (s *ShiftService) CreateShift(input viewmodel.ShiftReq) (int64, error) {
	var shift model.Shift
	copier.Copy(&shift, &input)
	return s.repo.CreateShift(shift)
}

func (s *ShiftService) EditShift(input viewmodel.ShiftReq, shiftID int) error {
	var shift model.Shift
	copier.Copy(&shift, &input)
	shift.ID = shiftID
	return s.repo.EditShift(shift)
}

func (s *ShiftService) DeleteShift(shiftID int) error {
	return s.repo.DeleteShift(shiftID)
}

func (s *ShiftService) GetPendingRequests() ([]viewmodel.RequestRes, error) {
	var vm []viewmodel.RequestRes
	requests, err := s.repo.GetPendingRequests()
	copier.Copy(&vm, &requests)
	return vm, err
}

func (s *ShiftService) ApproveShiftRequest(requestID int) error {
	return s.repo.ApproveShiftRequest(requestID)
}

func (s *ShiftService) RejectShiftRequest(requestID int) error {
	return s.repo.RejectShiftRequest(requestID)
}
