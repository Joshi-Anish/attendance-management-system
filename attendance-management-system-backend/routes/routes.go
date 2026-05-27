package routes

import (
	"attendance-management-system/handlers"
	"net/http"
)

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		// handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func SetupRoutes() {

	http.HandleFunc("/students", cors(studentsHandler))

	// DELETE STUDENT ROUTE
	http.HandleFunc("/students/delete", cors(handlers.DeleteStudent))

	http.HandleFunc("/checkin", cors(handlers.CheckIn))
	http.HandleFunc("/checkout", cors(handlers.CheckOut))

	http.HandleFunc("/dashboard", cors(handlers.GetDashboard))

	http.HandleFunc("/attendance", cors(handlers.GetAttendanceHistory))
	http.HandleFunc("/attendance/clear", cors(handlers.ClearAttendance))
	http.HandleFunc("/attendance/download", cors(handlers.DownloadAttendance))
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handlers.GetStudents(w, r)
	case http.MethodPost:
		handlers.CreateStudent(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
