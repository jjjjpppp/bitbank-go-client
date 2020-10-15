package bitbank

import (
	"bitbank-go-client/v1/request"
	"context"
	"fmt"
	"strings"

	"github.com/jjjjpppp/bitbank-go-client/v1/models"
)

func (c *Client) GetOrder(ctx context.Context, params request.GetOrderParams) (*models.Order, error) {
	spath := fmt.Sprintf("/user/spot/order")
	queryParam := make(map[string]string)
	
	// set required param
	queryParam["pair"] = *params.Pair
	queryParam["order_id"] = *params.OrderID

	res, err := c.sendRequest(ctx, "GET", spath, nil, &queryParam)
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

func (c *Client) GetActiveOrders(ctx context.Context, params request.GetActiveOrdersParams) (*models.Orders, error) {
	spath := fmt.Sprintf("/user/spot/active_orders")
	queryParam := make(map[string]string)

	// set required param
	queryParam["pair"] = *params.Pair

	// set optional param
	if params.Count != nil {
		queryParam["count"] = fmt.Sprint(*params.Count)
	}
	if params.FromID != nil {
		queryParam["from_id"] = fmt.Sprint(*params.FromID)
	}
	if params.EndID != nil {
		queryParam["end_id"] = fmt.Sprint(*params.EndID)
	}
	if params.Since != nil {
		queryParam["since"] = fmt.Sprint(*params.Since)
	}
	if params.End != nil {
		queryParam["end"] = fmt.Sprint(*params.End)
	}

	res, err := c.sendRequest(ctx, "GET", spath, nil, &queryParam)
	if err != nil {
		return nil, err
	}

	var orders models.Orders
	if err := decodeBody(res, &orders); err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *Client) CancelOrder(ctx context.Context, pair string, orderID int) (*models.Order, error) {
	spath := fmt.Sprintf("/user/spot/cancel_order")
	bodyTemplate :=
		`{
  "pair":"%s",
  "order_id":%d
}`
	body := fmt.Sprintf(bodyTemplate, pair, orderID)
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

func (c *Client) CancelOrders(ctx context.Context, pair string, orderIDs []int) (*models.Orders, error) {
	spath := fmt.Sprintf("/user/spot/cancel_orders")
	bodyTemplate :=
		`{
  "pair":"%s",
  "order_ids":[%s]
}`
	body := fmt.Sprintf(bodyTemplate, pair, arrayToString(orderIDs, ","))
	res, err := c.sendRequest(ctx, "POST", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var orders models.Orders
	if err := decodeBody(res, &orders); err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *Client) GetOrdersInfo(ctx context.Context, pair string, orderIDs []int) (*models.Orders, error) {
	spath := fmt.Sprintf("/user/spot/orders_info")
	bodyTemplate :=
		`{
  "pair":"%s",
  "order_ids":[%s]
}`
	body := fmt.Sprintf(bodyTemplate, pair, arrayToString(orderIDs, ","))
	res, err := c.sendRequest(ctx, "POST", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var orders models.Orders
	if err := decodeBody(res, &orders); err != nil {
		return nil, err
	}

	return &orders, nil
}
