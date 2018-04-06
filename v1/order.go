package bitbank

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
)

func (c *Client) GetOrder(ctx context.Context, pair, orderID string) (*models.Order, error) {
	spath := fmt.Sprintf("/user/spot/order")
	queryParam := &map[string]string{
		"pair":     pair,
		"order_id": orderID}
	res, err := c.sendRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
