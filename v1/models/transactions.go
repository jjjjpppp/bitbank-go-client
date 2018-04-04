package models

type Transactions struct {
	Success int               `json:"success"`
	Data    *TransactionsData `json:"data"`
}

type TransactionsData struct {
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionID int    `json:"transaction_id"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Amount        string `json:"amount"`
	ExecutedAt    int    `json:"executed_at"`
}
