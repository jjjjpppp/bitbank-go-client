package models

type Order struct {
	Success int        `json:"success"`
	Data    *OrderData `json:"data"`
}

type OrderData struct {
	OrderID         int    `json:"order_id"`
	Pair            string `json:"pair"`
	Side            string `json:"side"`
	Type            string `json:"type"`
	StartAmount     string `json:"start_amount"`
	RemainingAmount string `json:"remaining_amount"`
	ExecutedAmount  string `json:"executed_amount"`
	Price           string `json:"price"`
	AveragePrice    string `json:"average_price"`
	OrderedAt       int    `json:"ordered_at"`
	Status          string `json:"status"`
}
