package handlers

import (
	"attendance-management-system/models"
	"attendance-management-system/storage"
	"encoding/json"
	"net/http"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	storage.Students = append(storage.Students, student)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Students)
}
