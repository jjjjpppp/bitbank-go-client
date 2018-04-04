package models

type Ticker struct {
	Success int         `json:"success"`
	Data    *TickerData `json:"data"`
}

type TickerData struct {
	Sell      string `json:"sell"`
	Buy       string `json:"buy"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Last      string `json:"last"`
	Vol       string `json:"vol"`
	Timestamp int    `json:"timestamp"`
}
