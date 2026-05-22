package routes

import (
	"attendance-management-system/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/students", studentsHandler)
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		handlers.CreateStudent(w, r)
		return
	}

	if r.Method == http.MethodGet {
		handlers.GetStudents(w, r)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
