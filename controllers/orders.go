package controllers

import (
	"encoding/json"
	"net/http"

	"mini-trading-platform/database"
)

type OrderHistory struct {
	ID        int    `json:"id"`
	Symbol    string `json:"symbol"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	Timestamp string `json:"timestamp"`
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	query := `
        SELECT o.id, s.symbol, o.type, o.quantity, o.timestamp
        FROM orders o
        JOIN stocks s ON o.stock_id = s.id
        WHERE o.user_id = $1
        ORDER BY o.timestamp DESC;
    `

	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []OrderHistory
	for rows.Next() {
		var o OrderHistory
		err := rows.Scan(&o.ID, &o.Symbol, &o.Type, &o.Quantity, &o.Timestamp)
		if err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		orders = append(orders, o)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
