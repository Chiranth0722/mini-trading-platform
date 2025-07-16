package controllers

import (
	"encoding/json"
	"net/http"

	"mini-trading-platform/database"
	"mini-trading-platform/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// 👇 Handle preflight request from browser
		w.WriteHeader(http.StatusOK)
		return
	}

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err := database.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID)

	if err != nil {
		http.Error(w, "Error inserting user: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
