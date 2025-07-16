package controllers

import (
	"encoding/json"
	"net/http"

	"mini-trading-platform/database"
	"mini-trading-platform/models"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	query := `INSERT INTO orders (user_id, stock_id, type, quantity) VALUES ($1, $2, $3, $4) RETURNING id`
	err := database.DB.QueryRow(query, order.UserID, order.StockID, order.Type, order.Quantity).Scan(&order.ID)

	if err != nil {
		http.Error(w, "Failed to place order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
