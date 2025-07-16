package models

type Order struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	StockID  int    `json:"stock_id"`
	Type     string `json:"type"` // BUY or SELL
	Quantity int    `json:"quantity"`
}
