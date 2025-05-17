package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/justindarmawan/payd-mini-project/backend/model"
)

type ShiftRepository struct {
	db *sql.DB
}

func NewShiftRepository(db *sql.DB) *ShiftRepository {
	return &ShiftRepository{db: db}
}

// worker
func (r *ShiftRepository) GetAvailableShifts() ([]model.Shift, error) {
	rows, err := r.db.Query(`
		SELECT s.id, s.date, s.start_time, s.end_time, s.role, s.location
		FROM shifts s
		LEFT JOIN shift_requests sr ON 
			s.id = sr.shift_id 
			AND sr.status = 'approved'
		WHERE 1=1
			AND date(s.date) >= date('now')
			AND sr.id IS NULL
		ORDER BY s.date, s.start_time
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []model.Shift
	for rows.Next() {
		var s model.Shift
		if err := rows.Scan(&s.ID, &s.Date, &s.StartTime, &s.EndTime, &s.Role, &s.Location); err != nil {
			return nil, err
		}
		shifts = append(shifts, s)
	}

	return shifts, nil
}

func (r *ShiftRepository) GetAssignedShifts(workerID int) ([]model.Shift, error) {
	rows, err := r.db.Query(`
		SELECT s.id, s.date, s.start_time, s.end_time, s.role, s.location
        FROM shifts s
        JOIN shift_requests sr ON 
			s.id = sr.shift_id
        WHERE  1=1
			AND sr.worker_id = ? 
			AND sr.status = 'approved'
			AND date(s.date) >= date('now')
        ORDER BY s.date, s.start_time
	`, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []model.Shift
	for rows.Next() {
		var s model.Shift
		if err := rows.Scan(&s.ID, &s.Date, &s.StartTime, &s.EndTime, &s.Role, &s.Location); err != nil {
			return nil, err
		}
		shifts = append(shifts, s)
	}

	return shifts, nil
}

func (r *ShiftRepository) CreateShiftRequest(input model.ShiftRequest) (int64, error) {
	query := `
		SELECT
			(
				SELECT COUNT(*) 
				FROM shift_requests 
				WHERE shift_id = ? 
					AND status = 'approved'
			) AS shift_assigned,
		
			(
				SELECT COUNT(*) 
				FROM shifts s
				JOIN shift_requests sr ON 
					s.id = sr.shift_id
				WHERE sr.worker_id = ? 
					AND sr.status IN ('pending', 'approved')
					AND s.date = (SELECT date FROM shifts WHERE id = ?)
					AND NOT (s.end_time <= (SELECT start_time FROM shifts WHERE id = ?)
					OR s.start_time >= (SELECT end_time FROM shifts WHERE id = ?))
			) AS overlapping,
		
			(
				SELECT COUNT(*) 
				FROM shifts s
				JOIN shift_requests sr ON 
					s.id = sr.shift_id
				WHERE sr.worker_id = ? 
					AND sr.status = 'approved'
					AND s.date = (SELECT date FROM shifts WHERE id = ?)
			) AS daily_count,
		
			(
				SELECT COUNT(*) 
				FROM shifts s
				JOIN shift_requests sr ON 
					s.id = sr.shift_id
				WHERE sr.worker_id = ? 
					AND sr.status = 'approved'
					AND s.date BETWEEN (SELECT date(date, 'weekday 1', '-7 days') FROM shifts WHERE id = ?)
					AND (SELECT date(date, 'weekday 1') FROM shifts WHERE id = ?)
			) AS weekly_count
	;
    `
	var shiftAssigned, overlapping, dailyCount, weeklyCount int

	err := r.db.QueryRow(query,
		input.ShiftID,
		input.WorkerID, input.ShiftID, input.ShiftID, input.ShiftID,
		input.WorkerID, input.ShiftID,
		input.WorkerID, input.ShiftID, input.ShiftID,
	).Scan(&shiftAssigned, &overlapping, &dailyCount, &weeklyCount)

	if err != nil {
		return 0, errors.New("database error")
	}
	if shiftAssigned > 0 {
		return 0, errors.New("shift already assigned")
	}
	if overlapping > 0 {
		return 0, errors.New("overlapping shift request exists")
	}
	if dailyCount >= 1 {
		return 0, errors.New("already assigned to a shift on this day")
	}
	if weeklyCount >= 5 {
		return 0, errors.New("shift limit of 5 per week reached")
	}

	result, err := r.db.Exec(`
        INSERT INTO shift_requests (shift_id, worker_id)
        VALUES (?, ?)
    `, input.ShiftID, input.WorkerID)

	if err != nil {
		return 0, errors.New("failed to create shift request")
	}

	return result.LastInsertId()
}

func (r *ShiftRepository) GetRequestsByWorker(workerID int) ([]model.Request, error) {
	rows, err := r.db.Query(`
		SELECT sr.id, sr.status, sr.requested_at
			,s.id, s.date, s.start_time, s.end_time, s.role, s.location
		FROM shift_requests sr
		JOIN shifts s ON 
			sr.shift_id = s.id
		WHERE sr.worker_id = ?
		ORDER BY s.date, s.start_time
	`, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []model.Request
	for rows.Next() {
		var r model.Request
		if err := rows.Scan(&r.RequestID, &r.Status, &r.RequestedAt, &r.ID, &r.Date, &r.StartTime, &r.EndTime, &r.Role, &r.Location); err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}

	return requests, nil
}

// admin
func (r *ShiftRepository) CreateShift(shift model.Shift) (int64, error) {
	query := `
        INSERT INTO shifts (date, start_time, end_time, role, location)
        VALUES (?, ?, ?, ?, ?)
    `
	result, err := r.db.Exec(query, shift.Date, shift.StartTime, shift.EndTime, shift.Role, shift.Location)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *ShiftRepository) EditShift(shift model.Shift) error {
	var fields []string
	var args []any

	if strings.TrimSpace(shift.Date) != "" {
		fields = append(fields, "date = ?")
		args = append(args, shift.Date)
	}
	if strings.TrimSpace(shift.StartTime) != "" {
		fields = append(fields, "start_time = ?")
		args = append(args, shift.StartTime)
	}
	if strings.TrimSpace(shift.EndTime) != "" {
		fields = append(fields, "end_time = ?")
		args = append(args, shift.EndTime)
	}
	if strings.TrimSpace(shift.Role) != "" {
		fields = append(fields, "role = ?")
		args = append(args, shift.Role)
	}
	if strings.TrimSpace(*shift.Location) != "" {
		fields = append(fields, "location = ?")
		args = append(args, shift.Location)
	}
	if len(fields) == 0 {
		return errors.New("no valid fields provided for update")
	}
	query := fmt.Sprintf("UPDATE shifts SET %s WHERE id = ?", strings.Join(fields, ", "))
	args = append(args, shift.ID)
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no shift found with given ID")
	}
	return nil
}

func (r *ShiftRepository) DeleteShift(shiftID int) error {
	row := r.db.QueryRow("SELECT COUNT(*) FROM shift_requests WHERE shift_id = ?", shiftID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete shift: shift has associated requests")
	}
	result, err := r.db.Exec("DELETE FROM shifts WHERE id = ?", shiftID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no shift found with given ID")
	}
	return nil
}

func (r *ShiftRepository) GetPendingRequests() ([]model.Request, error) {
	rows, err := r.db.Query(`
		SELECT sr.id, sr.status, sr.requested_at,
			s.id, s.date, s.start_time, s.end_time, s.role, s.location, u.username
		FROM shift_requests sr
		JOIN shifts s ON 
			sr.shift_id = s.id
		JOIN users u ON 
			sr.worker_id = u.id
		WHERE sr.status = 'pending'
			AND NOT EXISTS (
				SELECT 1 FROM shift_requests sr2
				WHERE sr2.shift_id = sr.shift_id 
					AND sr2.status = 'approved'
			)
		ORDER BY s.date, s.start_time, sr.requested_at
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reqs []model.Request
	for rows.Next() {
		var rqt model.Request
		if err := rows.Scan(&rqt.RequestID, &rqt.Status, &rqt.RequestedAt, &rqt.ID, &rqt.Date, &rqt.StartTime, &rqt.EndTime, &rqt.Role, &rqt.Location, &rqt.Username); err != nil {
			return nil, err
		}
		reqs = append(reqs, rqt)
	}
	return reqs, nil
}

func (r *ShiftRepository) ApproveShiftRequest(requestID int) error {
	var shiftID, workerID int
	err := r.db.QueryRow(`
        SELECT shift_id, worker_id FROM shift_requests WHERE id = ?
    `, requestID).Scan(&shiftID, &workerID)
	if err != nil {
		return err
	}
	var conflictCount int
	err = r.db.QueryRow(`
		SELECT COUNT(*) 
		FROM shift_requests sr
		JOIN shifts s ON 
			sr.shift_id = s.id
		WHERE sr.worker_id = ?
			AND sr.status = 'approved'
			AND (
				DATE(s.date) = (SELECT date FROM shifts WHERE id = ?)
				OR strftime('%W', s.date) = strftime('%W', (SELECT date FROM shifts WHERE id = ?))
			)
	`, workerID, shiftID, shiftID).Scan(&conflictCount)
	if err != nil {
		return err
	}
	if conflictCount >= 5 {
		return errors.New("worker has reached the max 5 approved shifts per week")
	}
	if conflictCount > 0 {
		return errors.New("worker already has a shift on this day or conflicting shift")
	}
	_, err = r.db.Exec("UPDATE shift_requests SET status = 'approved' WHERE id = ?", requestID)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(`
		UPDATE shift_requests SET status = 'rejected' 
		WHERE shift_id = ? AND id != ?
	`, shiftID, requestID)
	if err != nil {
		return err
	}
	return err
}

func (r *ShiftRepository) RejectShiftRequest(requestID int) error {
	result, err := r.db.Exec("UPDATE shift_requests SET status = 'rejected' WHERE id = ?", requestID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no shift request found with given ID")
	}
	return nil
}
