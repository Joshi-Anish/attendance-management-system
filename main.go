package main

import (
	"attendance-management-system/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.SetupRoutes()

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
