package handlers

import (
	"attendance-management-system/db"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Record struct {
	StudentID string  `json:"student_id"`
	CheckIn   string  `json:"check_in"`
	CheckOut  *string `json:"check_out"` // 👈 IMPORTANT FIX
}

func GetDashboard(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	var totalStudents int
	var activeStudents int
	var completedSessions int

	db.DB.QueryRow("SELECT COUNT(*) FROM students").Scan(&totalStudents)

	db.DB.QueryRow(`
		SELECT COUNT(*)
		FROM attendance
		WHERE check_out IS NULL
	`).Scan(&activeStudents)

	db.DB.QueryRow(`
		SELECT COUNT(*)
		FROM attendance
		WHERE check_out IS NOT NULL
	`).Scan(&completedSessions)

	rows, err := db.DB.Query(`
		SELECT student_id, check_in, check_out
		FROM attendance
		ORDER BY check_in DESC
	`)

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var records []Record

	for rows.Next() {
		var r Record
		var checkOut sql.NullString

		rows.Scan(&r.StudentID, &r.CheckIn, &checkOut)

		// ✅ FIX: convert empty string → NULL properly
		if checkOut.Valid {
			r.CheckOut = &checkOut.String
		} else {
			r.CheckOut = nil
		}

		records = append(records, r)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"total_students":     totalStudents,
		"active_students":    activeStudents,
		"completed_sessions": completedSessions,
		"records":            records,
	})
}
