package request

type GetWithdrawalAccountsParams struct {
	Asset string		`json:"asset,omitempty"`
}

type RequestWithdrawalParams struct {
	Asset string		`json:"asset,omitempty"`
	UuID string			`json:"uuid,omitempty"`
	Amount string		`json:"amount,omitempty"`
	OtpToken string		`json:"otp_token,omitempty"`
	SmsToken string		`json:"sms_token,omitempty"`
}
