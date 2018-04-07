package bitbank

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
)

func (c *Client) GetTrades(ctx context.Context, pair string, count, fromID, endID, since, end float64, order string) (*models.Trades, error) {
	spath := fmt.Sprintf("/user/spot/trade_history")
	queryParam := &map[string]string{
		"pair":    pair,
		"count":   fmt.Sprint(count),
		"from_id": fmt.Sprint(fromID),
		"end_id":  fmt.Sprint(endID),
		"since":   fmt.Sprint(since),
		"end":     fmt.Sprint(end),
		"order":   order}
	res, err := c.sendRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	var trades models.Trades
	if err := decodeBody(res, &trades); err != nil {
		return nil, err
	}

	return &trades, nil
}
