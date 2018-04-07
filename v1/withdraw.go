package bitbank

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"strings"
)

func (c *Client) GetWithdrawalAccounts(ctx context.Context, asset string) (*models.WithdrawalAccounts, error) {
	spath := fmt.Sprintf("/user/withdrawal_account")
	queryParam := &map[string]string{
		"asset": asset,
	}
	res, err := c.sendRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	var withdrawalAcdounts models.WithdrawalAccounts
	if err := decodeBody(res, &withdrawalAcdounts); err != nil {
		return nil, err
	}

	return &withdrawalAcdounts, nil
}

func (c *Client) RequestWithdrawal(ctx context.Context, asset, uuid, amount, otpToken, smsToken string) (*models.RequestWithdrawal, error) {
	spath := fmt.Sprintf("/user/request_withdrawal")
	bodyTemplate :=
		`{
  "asset":"%s",
  "uuid":"%s",
  "amount":"%s",
  "otp_token":"%s",
  "sms_token":"%s"
}`
	body := fmt.Sprintf(bodyTemplate, asset, uuid, amount, otpToken, smsToken)
	res, err := c.sendRequest(ctx, "POST", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var requestWithdrawal models.RequestWithdrawal
	if err := decodeBody(res, &requestWithdrawal); err != nil {
		return nil, err
	}

	return &requestWithdrawal, nil
}
