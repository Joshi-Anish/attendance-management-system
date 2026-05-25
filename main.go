package main

import (
	"fmt"
	"net/http"

	"attendance-management-system/db"
	"attendance-management-system/routes"
)

func main() {
	db.Connect()
	routes.SetupRoutes()

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
