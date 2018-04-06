package bitbank

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"strings"
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

func (c *Client) CreateOrder(ctx context.Context, pair, amount string, price int, side, _type string) (*models.Order, error) {
	spath := fmt.Sprintf("/user/spot/order")
	bodyTemplate :=
		`{
  "pair":"%s",
  "amount":"%s",
  "price":%d,
  "side":"%s",
  "type":"%s"
}`
	body := fmt.Sprintf(bodyTemplate, pair, amount, price, side, _type)
	res, err := c.sendRequest(ctx, "POST", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
