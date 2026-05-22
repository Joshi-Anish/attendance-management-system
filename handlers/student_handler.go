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

	if student.Name == "" {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}

	//validate by id
	for _, s := range storage.Students {
		if s.ID == student.ID {
			http.Error(w, "Student ID already exists", http.StatusBadRequest)
			return
		}
	}

	storage.Students = append(storage.Students, student)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Students)
}
