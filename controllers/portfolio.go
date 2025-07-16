package controllers

import (
	"encoding/json"
	"net/http"

	"mini-trading-platform/database"
)

type Holding struct {
	Symbol      string  `json:"symbol"`
	NetQuantity int     `json:"quantity"`
	LatestPrice float64 `json:"latest_price"`
	Value       float64 `json:"value"`
}

func GetPortfolio(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	query := `
        SELECT s.symbol, SUM(
            CASE 
                WHEN o.type = 'BUY' THEN o.quantity
                WHEN o.type = 'SELL' THEN -o.quantity
                ELSE 0
            END
        ) as net_quantity,
        s.price
        FROM orders o
        JOIN stocks s ON o.stock_id = s.id
        WHERE o.user_id = $1
        GROUP BY s.symbol, s.price
        HAVING SUM(
            CASE 
                WHEN o.type = 'BUY' THEN o.quantity
                WHEN o.type = 'SELL' THEN -o.quantity
                ELSE 0
            END
        ) > 0;
    `

	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, "Failed to fetch portfolio", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var holdings []Holding
	for rows.Next() {
		var h Holding
		err := rows.Scan(&h.Symbol, &h.NetQuantity, &h.LatestPrice)
		if err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		h.Value = float64(h.NetQuantity) * h.LatestPrice
		holdings = append(holdings, h)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(holdings)
}
