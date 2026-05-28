package handlers

import (
	"attendance-management-system/db"
	"encoding/json"
	"net/http"
)

// register
func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	json.NewDecoder(r.Body).Decode(&user)

	_, err := db.DB.Exec(
		"INSERT INTO users (username, password, role) VALUES (?, ?, ?)",
		user.Username, user.Password, user.Role,
	)

	if err != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

// login
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	var dbPassword string
	var role string

	err := db.DB.QueryRow(
		"SELECT password, role FROM users WHERE username = ?",
		input.Username,
	).Scan(&dbPassword, &role)

	if err != nil || dbPassword != input.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Login successful",
		"role":     role,
		"username": input.Username,
	})
}
