package models

type Depth struct {
	Success int `json:"success"`
	Data    struct {
		Asks [][]string `json:"asks"`
		Bids [][]string `json:"bids"`
	} `json:"data"`
}
