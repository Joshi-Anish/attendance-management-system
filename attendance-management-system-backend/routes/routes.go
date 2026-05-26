package routes

import (
	"attendance-management-system/handlers"
	"net/http"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}

func SetupRoutes() {
	http.HandleFunc("/students", studentsHandler)
	http.HandleFunc("/checkin", handlers.CheckIn)
	http.HandleFunc("/checkout", handlers.CheckOut)
	http.HandleFunc("/attendance/time", handlers.GetTotalTime)
	http.HandleFunc("/dashboard", handlers.GetDashboard)
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {

	enableCORS(w)

	if r.Method == http.MethodOptions {
		return
	}

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
