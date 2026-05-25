package handlers

import (
	"attendance-management-system/db"
	"encoding/json"
	"net/http"
	"time"
)

// checkin

func CheckIn(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		StudentID string `json:"student_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.StudentID == "" {
		http.Error(w, "StudentID is required", http.StatusBadRequest)
		return
	}

	// double checkin prevent garnaae
	var count int
	err := db.DB.QueryRow(
		"SELECT COUNT(*) FROM attendance WHERE student_id = ? AND check_out IS NULL",
		input.StudentID,
	).Scan(&count)

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	if count > 0 {
		http.Error(w, "Student already checked in", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"INSERT INTO attendance (student_id, check_in) VALUES (?, NOW())",
		input.StudentID,
	)

	if err != nil {
		http.Error(w, "Failed to check in", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Check-in successful",
	})
}

// check out kura haru
func CheckOut(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		StudentID string `json:"student_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.StudentID == "" {
		http.Error(w, "StudentID is required", http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec(
		"UPDATE attendance SET check_out = NOW() WHERE student_id = ? AND check_out IS NULL",
		input.StudentID,
	)

	if err != nil {
		http.Error(w, "Failed to check out", http.StatusInternalServerError)
		return
	}

	rows, _ := res.RowsAffected()

	if rows == 0 {
		http.Error(w, "No active check-in found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Check-out successful",
	})
}

// total time
func GetTotalTime(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	studentID := r.URL.Query().Get("student_id")

	if studentID == "" {
		http.Error(w, "student_id is required", http.StatusBadRequest)
		return
	}

	var totalSeconds int

	err := db.DB.QueryRow(`
		SELECT IFNULL(SUM(TIMESTAMPDIFF(SECOND, check_in, check_out)), 0)
		FROM attendance
		WHERE student_id = ? AND check_out IS NOT NULL
	`, studentID).Scan(&totalSeconds)

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	duration := time.Duration(totalSeconds) * time.Second

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"student_id": studentID,
		"total_time": duration.String(),
	})
}
