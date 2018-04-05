package models

import (
	"encoding/json"
)

type Assets struct {
	Success int         `json:"success"`
	Data    *AssetsData `json:"data"`
}

type AssetsData struct {
	Assets []*Asset `json:"assets"`
}

type Asset struct {
	Asset           string `json:"asset"`
	AmountPrecision int    `json:"amount_precision"`
	OnhandAmount    string `json:"onhand_amount"`
	LockedAmount    string `json:"locked_amount"`
	FreeAmount      string `json:"free_amount"`
	StopDeposit     bool   `json:"stop_deposit"`
	StopWithdrawal  bool   `json:"stop_withdrawal"`
	// withdrawl_fee is either a string or object
	//      "withdrawal_fee": {
	//        "threshold": "30000.0000",
	//        "under": "540.0000",
	//        "over": "756.0000"
	//      }
	// or
	// "withdrawal_fee": "0.00100000"
	WithdrawalFee *WithdrawalFee `json:"withdrawal_fee"`
}

type WithdrawalFee struct {
	Fee       string
	Threshold string
	Under     string
	Over      string
}

func (w *WithdrawalFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.Fee)
}

func (w *WithdrawalFee) UnmarshalJSON(data []byte) error {
	var raw interface{}
	json.Unmarshal(data, &raw)
	switch raw := raw.(type) {
	case string:
		*w = WithdrawalFee{Fee: raw}
	case map[string]interface{}:
		*w = WithdrawalFee{
			Threshold: raw["threshold"].(string),
			Under:     raw["under"].(string),
			Over:      raw["over"].(string),
		}
	}
	return nil
}
