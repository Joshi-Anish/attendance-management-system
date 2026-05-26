package handlers

import (
	"attendance-management-system/db"
	"encoding/json"
	"net/http"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	// total students today
	var totalStudents int
	db.DB.QueryRow(`
		SELECT COUNT(DISTINCT student_id)
		FROM attendance
	`).Scan(&totalStudents)

	// active students (not checked out yet)
	var activeStudents int
	db.DB.QueryRow(`
		SELECT COUNT(*)
		FROM attendance
		WHERE check_out IS NULL
	`).Scan(&activeStudents)

	// completed sessions
	var completedSessions int
	db.DB.QueryRow(`
		SELECT COUNT(*)
		FROM attendance
		WHERE check_out IS NOT NULL
	`).Scan(&completedSessions)

	// full records
	rows, err := db.DB.Query(`SELECT student_id, check_in, check_out FROM attendance`)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Record struct {
		StudentID string `json:"student_id"`
		CheckIn   string `json:"check_in"`
		CheckOut  string `json:"check_out"`
	}

	var records []Record

	for rows.Next() {
		var r Record
		rows.Scan(&r.StudentID, &r.CheckIn, &r.CheckOut)
		records = append(records, r)
	}

	response := map[string]interface{}{
		"total_students":     totalStudents,
		"active_students":    activeStudents,
		"completed_sessions": completedSessions,
		"records":            records,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
