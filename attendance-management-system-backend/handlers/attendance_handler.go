package handlers

import (
	"attendance-management-system/db"
	"encoding/json"
	"net/http"
	"time"
)

// ================= CHECK-IN =================
func CheckIn(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		StudentID string `json:"student_id"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	var count int
	db.DB.QueryRow(
		"SELECT COUNT(*) FROM attendance WHERE student_id = ? AND check_out IS NULL",
		input.StudentID,
	).Scan(&count)

	if count > 0 {
		http.Error(w, "Student already checked in", http.StatusBadRequest)
		return
	}

	db.DB.Exec(
		"INSERT INTO attendance (student_id, check_in) VALUES (?, NOW())",
		input.StudentID,
	)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Check-in successful",
	})
}

// ================= CHECK-OUT =================
func CheckOut(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		StudentID string `json:"student_id"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	res, err := db.DB.Exec(
		"UPDATE attendance SET check_out = NOW() WHERE student_id = ? AND check_out IS NULL",
		input.StudentID,
	)

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	rows, _ := res.RowsAffected()

	if rows == 0 {
		http.Error(w, "No active check-in found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Check-out successful",
	})
}

// ================= TOTAL TIME =================
func GetTotalTime(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	studentID := r.URL.Query().Get("student_id")

	var totalSeconds int

	db.DB.QueryRow(`
		SELECT IFNULL(SUM(TIMESTAMPDIFF(SECOND, check_in, check_out)), 0)
		FROM attendance
		WHERE student_id = ? AND check_out IS NOT NULL
	`, studentID).Scan(&totalSeconds)

	duration := time.Duration(totalSeconds) * time.Second

	json.NewEncoder(w).Encode(map[string]string{
		"student_id": studentID,
		"total_time": duration.String(),
	})
}
