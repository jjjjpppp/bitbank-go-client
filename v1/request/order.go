package request

type GetOrderParams struct {
	Pair *string
	OrderID *string
}

type GetActiveOrdersParams struct {
	Pair   *string
	Count  *float64
	FromID *float64
	EndID  *float64
	Since  *float64
	End    *float64
}