package controllers

import (
	"encoding/json"
	"net/http"

	"mini-trading-platform/database"
	"mini-trading-platform/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err := database.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID)

	if err != nil {
		http.Error(w, "User already exists or DB error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
