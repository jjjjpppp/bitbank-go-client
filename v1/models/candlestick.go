package models

import (
	"encoding/json"
)

type Candlesticks struct {
	Success int               `json:"success"`
	Data    *CandlesticksData `json:"data"`
}

type CandlesticksData struct {
	Candlesticks []*Candlestick `json:"candlestick"`
}

type Candlestick struct {
	Type  string          `json:"type"`
	Ohlcv [][]json.Number `json:"ohlcv"`
}
