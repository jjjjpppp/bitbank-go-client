package request

type GetTradeParams struct {
	Pair string		`json:"pair,omitempty"`
	Count float64	`json:"count,omitempty"`
	FromID float64	`json:"from_id,omitempty"`
	EndID float64	`json:"end_id,omitempty"`
	Since float64	`json:"since,omitempty"`
	End float64		`json:"end,omitempty"`
	Order string	`json:"order,omitempty"`
}
