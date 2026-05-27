package handlers

import (
	"attendance-management-system/db"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"net/http"
)

func CheckIn(w http.ResponseWriter, r *http.Request) {

	var input struct {
		StudentID string `json:"student_id"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	var count int
	db.DB.QueryRow(`
		SELECT COUNT(*)
		FROM attendance
		WHERE student_id = ? AND check_out IS NULL
	`, input.StudentID).Scan(&count)

	if count > 0 {
		http.Error(w, "Already checked in", http.StatusBadRequest)
		return
	}

	db.DB.Exec(`
		INSERT INTO attendance (student_id, check_in)
		VALUES (?, NOW())
	`, input.StudentID)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Check-in successful",
	})
}

func CheckOut(w http.ResponseWriter, r *http.Request) {

	var input struct {
		StudentID string `json:"student_id"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	res, err := db.DB.Exec(`
		UPDATE attendance
		SET check_out = NOW()
		WHERE student_id = ? AND check_out IS NULL
	`, input.StudentID)

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	rows, _ := res.RowsAffected()

	if rows == 0 {
		http.Error(w, "No active session", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Check-out successful",
	})
}

func GetAttendanceHistory(w http.ResponseWriter, r *http.Request) {

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

	type Record struct {
		StudentID string  `json:"student_id"`
		CheckIn   string  `json:"check_in"`
		CheckOut  *string `json:"check_out"`
	}

	var records []Record

	for rows.Next() {
		var r Record
		var checkOut sql.NullString

		rows.Scan(&r.StudentID, &r.CheckIn, &checkOut)

		if checkOut.Valid {
			r.CheckOut = &checkOut.String
		} else {
			r.CheckOut = nil
		}

		records = append(records, r)
	}

	json.NewEncoder(w).Encode(records)
}

// clear attendance history
func ClearAttendance(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := db.DB.Exec("DELETE FROM attendance")
	if err != nil {
		http.Error(w, "Failed to clear data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "All attendance data cleared",
	})
}

// download attendance history

func DownloadAttendance(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT student_id, check_in, check_out FROM attendance")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Disposition", "attachment;filename=attendance.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// header
	writer.Write([]string{"Student ID", "Check In", "Check Out"})

	for rows.Next() {
		var id, checkIn, checkOut string
		rows.Scan(&id, &checkIn, &checkOut)

		writer.Write([]string{id, checkIn, checkOut})
	}
}
