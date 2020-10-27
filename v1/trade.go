package bitbank

import (
	"context"
	"fmt"

	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/jjjjpppp/bitbank-go-client/v1/request"
)

func (c *Client) GetTrades(ctx context.Context, params request.GetTradeParams) (*models.Trades, error) {
	spath := fmt.Sprintf("/user/spot/trade_history")
	queryParam := make(map[string]string)

	// set required param
	queryParam["pair"] = params.Pair

	// set optional param
	if params.Count != 0.0 {
		queryParam["count"] = fmt.Sprint(params.Count)
	}
	if params.FromID != 0.0 {
		queryParam["from_id"] = fmt.Sprint(params.FromID)
	}
	if params.EndID != 0.0 {
		queryParam["end_id"] = fmt.Sprint(params.EndID)
	}
	if params.Since != 0.0 {
		queryParam["since"] = fmt.Sprint(params.Since)
	}
	if params.End != 0.0 {
		queryParam["end"] = fmt.Sprint(params.End)
	}
	if params.Order != "" {
		queryParam["order"] = fmt.Sprint(params.Order)
	}

	res, err := c.sendRequest(ctx, "GET", spath, nil, &queryParam)
	if err != nil {
		return nil, err
	}

	var trades models.Trades
	if err := decodeBody(res, &trades); err != nil {
		return nil, err
	}

	return &trades, nil
}
