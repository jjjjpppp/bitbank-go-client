package models

type WithdrawalAccounts struct {
	Success int                     `json:"success"`
	Data    *WithdrawalAccountsData `json:"data"`
}

type WithdrawalAccountsData struct {
	Accounts []*Account `json:"accounts"`
}

type Account struct {
	UUID    string `json:"uuid"`
	Label   string `json:"label"`
	Address string `json:"address"`
}

type RequestWithdrawal struct {
	Success int                    `json:"success"`
	Data    *RequestWithdrawalData `json:"data"`
}

type RequestWithdrawalData struct {
	UUID        string `json:"uuid"`
	Asset       string `json:"asset"`
	Amount      int    `json:"amount"`
	AccountUUID string `json:"account_uuid"`
	Fee         string `json:"fee"`
	Status      string `json:"status"`
	Label       string `json:"label"`
	Txid        string `json:"txid"`
	Address     string `json:"address"`
}
