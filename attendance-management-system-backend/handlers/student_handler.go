package handlers

import (
	"attendance-management-system/db"
	"encoding/json"
	"net/http"
)

// ================= CREATE STUDENT =================
func CreateStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var student struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Grade   string `json:"grade"`
	}

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if student.Name == "" {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}

	// check duplicate ID in DB
	var count int
	err := db.DB.QueryRow(
		"SELECT COUNT(*) FROM students WHERE id = ?",
		student.ID,
	).Scan(&count)

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	if count > 0 {
		http.Error(w, "Student ID already exists", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"INSERT INTO students (id, name, address, grade) VALUES (?, ?, ?, ?)",
		student.ID, student.Name, student.Address, student.Grade,
	)

	if err != nil {
		http.Error(w, "Failed to create student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// ================= GET STUDENTS =================
func GetStudents(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query("SELECT id, name, address, grade FROM students")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Student struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Grade   string `json:"grade"`
	}

	var students []Student

	for rows.Next() {
		var s Student
		rows.Scan(&s.ID, &s.Name, &s.Address, &s.Grade)
		students = append(students, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
