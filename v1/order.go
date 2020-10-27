package bitbank

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/jjjjpppp/bitbank-go-client/v1/request"
)

func (c *Client) GetOrder(ctx context.Context, params request.GetOrderParams) (*models.Order, error) {
	spath := fmt.Sprintf("/user/spot/order")
	queryParam := make(map[string]string)

	// set required param
	if params.Pair == "" {
		return nil, errors.New("pair parameter is required")
	}
	queryParam["pair"] = params.Pair
	if params.OrderID == "" {
		return nil, errors.New("order_id parameter is required")
	}
	queryParam["order_id"] = params.OrderID

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

func (c *Client) CreateOrder(ctx context.Context, params request.CreateOrderParams) (*models.Order, error) {
	spath := fmt.Sprintf("/user/spot/order")

	// check required param
	if params.Pair == "" {
		return nil, errors.New("pair parameter is required")
	}
	if params.Amount == "" {
		return nil, errors.New("amount parameter is required")
	}
	if params.Side == "" {
		return nil, errors.New("side parameter is required")
	}
	if params.Type == "" {
		return nil, errors.New("type parameter is required")
	}

	// build param
	bodyTemplate, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(bodyTemplate), "", "  ")
	if err != nil {
		return nil, err
	}
	body := buf.String()
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
	if params.Pair == "" {
		return nil, errors.New("pair parameter is required")
	}
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

func (c *Client) CancelOrder(ctx context.Context, params request.CancelOrderParams) (*models.Order, error) {
	spath := fmt.Sprintf("/user/spot/cancel_order")

	// check reuqired param
	if params.Pair == "" {
		return nil, errors.New("pair parameter is required")
	}
	if params.OrderID == 0.0 {
		return nil, errors.New("order_id parameter is required")
	}

	// build param
	bodyTemplate, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(bodyTemplate), "", "  ")
	if err != nil {
		return nil, err
	}
	body := buf.String()
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

func (c *Client) CancelOrders(ctx context.Context, params request.CancelOrdersParams) (*models.Orders, error) {
	spath := fmt.Sprintf("/user/spot/cancel_orders")

	// check required param
	if params.Pair == "" {
		return nil, errors.New("pair parameter is required")
	}
	if len(params.OrderIDs) == 0 {
		return nil, errors.New("order_ids parameter is required")
	}

	// build param
	bodyTemplate, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(bodyTemplate), "", "  ")
	if err != nil {
		return nil, err
	}
	body := buf.String()
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

func (c *Client) GetOrdersInfo(ctx context.Context, params request.GetOrdersInfoParams) (*models.Orders, error) {
	spath := fmt.Sprintf("/user/spot/orders_info")

	// check required param
	if params.Pair == "" {
		return nil, errors.New("pair parameter is required")
	}
	if len(params.OrderIDs) == 0 {
		return nil, errors.New("order_ids parameter is required")
	}

	// build param
	bodyTemplate, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(bodyTemplate), "", "  ")
	if err != nil {
		return nil, err
	}
	body := buf.String()
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
