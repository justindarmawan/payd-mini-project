package model

type Shift struct {
	ID        int
	Date      string
	StartTime string
	EndTime   string
	Role      string
	Location  *string
}

type ShiftRequest struct {
	ID       int
	ShiftID  int
	WorkerID int
	Status   string
}

type Request struct {
	RequestID   int
	Status      string
	RequestedAt string
	Shift
	Username string
}
