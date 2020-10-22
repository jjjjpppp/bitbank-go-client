package request

type GetOrderParams struct {
	Pair	string	`json:"pair,omitempty"`
	OrderID	string	`json:"order_id,omitempty"`
}

type CreateOrderParams struct {
	Pair	string	`json:"pair,omitempty"`
	Amount	string	`json:"amount,omitempty"`
	Price	int		`json:"price,omitempty"`
	Side	string	`json:"side,omitempty"`
	Type	string	`json:"type,omitempty"`
}

type GetActiveOrdersParams struct {
	Pair   string	`json:"pair,omitempty"`
	Count  float64	`json:"count,omitempty"`
	FromID float64	`json:"from_id,omitempty"`
	EndID  float64	`json:"end_id,omitempty"`
	Since  float64	`json:"since,omitempty"`
	End    float64	`json:"end,omitempty"`
}