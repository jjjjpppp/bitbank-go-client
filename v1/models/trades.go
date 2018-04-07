package models

type Trades struct {
	Success int         `json:"success"`
	Data    *TradesData `json:"data"`
}

type TradesData struct {
	Trades []*Trade `json:"trades"`
}

type Trade struct {
	TradeID        int    `json:"trade_id"`
	Pair           string `json:"pair"`
	OrderID        int    `json:"order_id"`
	Side           string `json:"side"`
	Type           string `json:"type"`
	Amount         string `json:"amount"`
	Price          string `json:"price"`
	MakerTaker     string `json:"maker_taker"`
	FeeAmountBase  string `json:"fee_amount_base"`
	FeeAmountQuote string `json:"fee_amount_quote"`
	ExecutedAt     int    `json:"executed_at"`
}
