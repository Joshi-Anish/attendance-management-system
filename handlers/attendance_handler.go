package handlers

import (
	"attendance-management-system/models"
	"attendance-management-system/storage"
	"encoding/json"
	"net/http"
	"time"
)

// CHECK-IN
func CheckIn(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		StudentID string `json:"student_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//  validate empty ID
	if input.StudentID == "" {
		http.Error(w, "StudentID is required", http.StatusBadRequest)
		return
	}

	//  PREVENT DOUBLE CHECK-IN
	for _, record := range storage.Attendance {
		if record.StudentID == input.StudentID && record.CheckOut.IsZero() {
			http.Error(w, "Student already checked in", http.StatusBadRequest)
			return
		}
	}

	record := models.Attendance{
		StudentID: input.StudentID,
		CheckIn:   time.Now(),
	}

	storage.Attendance = append(storage.Attendance, record)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

// CHECK-OUT
func CheckOut(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		StudentID string `json:"student_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//  validate empty ID
	if input.StudentID == "" {
		http.Error(w, "StudentID is required", http.StatusBadRequest)
		return
	}

	// find latest active record (no checkout yet)
	for i := len(storage.Attendance) - 1; i >= 0; i-- {
		if storage.Attendance[i].StudentID == input.StudentID &&
			storage.Attendance[i].CheckOut.IsZero() {

			//  update checkout time
			storage.Attendance[i].CheckOut = time.Now()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(storage.Attendance[i])
			return
		}
	}

	http.Error(w, "No active check-in found", http.StatusNotFound)
}

// time calculation
func GetTotalTime(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	// get student_id from URL
	studentID := r.URL.Query().Get("student_id")

	if studentID == "" {
		http.Error(w, "student_id is required", http.StatusBadRequest)
		return
	}

	var totalTime time.Duration

	for _, record := range storage.Attendance {

		if record.StudentID == studentID && !record.CheckOut.IsZero() {

			duration := record.CheckOut.Sub(record.CheckIn)
			totalTime += duration
		}
	}

	response := map[string]string{
		"student_id": studentID,
		"total_time": totalTime.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
