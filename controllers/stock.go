package controllers

import (
	"encoding/json"
	"net/http"

	"mini-trading-platform/database"
	"mini-trading-platform/models"
)

func GetStocks(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, symbol, name, price FROM stocks")
	if err != nil {
		http.Error(w, "Failed to fetch stocks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var s models.Stock
		err := rows.Scan(&s.ID, &s.Symbol, &s.Name, &s.Price)
		if err != nil {
			http.Error(w, "Error reading stock", http.StatusInternalServerError)
			return
		}
		stocks = append(stocks, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}
