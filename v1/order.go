package bitbank

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
)

func (c *Client) GetOrder(ctx context.Context) (*models.Order, error) {
	spath := fmt.Sprintf("/user/spot/order")
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
