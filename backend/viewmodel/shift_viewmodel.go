package viewmodel

type ShiftRes struct {
	ID        int     `json:"id"`
	Date      string  `json:"date"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	Role      string  `json:"role"`
	Location  *string `json:"location,omitempty"`
}

type RequestRes struct {
	RequestID   int     `json:"request_id"`
	Status      string  `json:"status"`
	RequestedAt string  `json:"requested_at"`
	ID          int     `json:"shift_id"`
	Date        string  `json:"date"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Role        string  `json:"role"`
	Location    *string `json:"location,omitempty"`
	Username    string  `json:"username,omitempty"`
}

type ShiftReq struct {
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Role      string `json:"role"`
	Location  string `json:"location"`
}

type ShiftRequestReq struct {
	ShiftID  int `json:"shift_id" binding:"required"`
	WorkerID int `json:"worker_id" binding:"required"`
}
