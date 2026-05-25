package handlers

import (
	"attendance-management-system/storage"
	"encoding/json"
	"net/http"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	totalStudents := make(map[string]bool)
	activeCount := 0
	completedSessions := 0

	for _, record := range storage.Attendance {

		totalStudents[record.StudentID] = true

		if record.CheckOut.IsZero() {
			activeCount++
		} else {
			completedSessions++
		}
	}

	response := map[string]interface{}{
		"total_students_today": len(totalStudents),
		"active_students":      activeCount,
		"completed_sessions":   completedSessions,
		"records":              storage.Attendance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
