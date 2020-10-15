package bitbank

import (
	"bitbank-go-client/v1/request"
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/jjjjpppp/bitbank-go-client/v1/testutil"
)

func TestGetOrder(t *testing.T) {
	type Param struct {
		pair         string
		orderID      string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", orderID: "12345", jsonResponse: testutil.GetOrderJsonResponse()},
			expect: Expect{path: "/user/spot/order?order_id=12345&pair=btc_jpy", method: "GET", body: "", e: testutil.ExpectedGetOrderModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		r, err := client.GetOrder(ctx, c.param.pair, c.param.orderID)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestGetActiveOrders(t *testing.T) {
	type Param struct {
		pair         string
		count        float64
		fromID       float64
		endID        float64
		since        float64
		end          float64
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Orders
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", count: 1, fromID: 1, endID: 10, since: 12345, end: 12355, jsonResponse: testutil.GetOrdersJsonResponse()},
			expect: Expect{path: "/user/spot/active_orders?count=1&end=12355&end_id=10&from_id=1&pair=btc_jpy&since=12345", method: "GET", body: "", e: testutil.ExpectedGetOrdersModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		params := request.GetActiveOrdersParams{&c.param.pair, &c.param.count, &c.param.fromID, &c.param.endID, &c.param.since, &c.param.end}
		r, err := client.GetActiveOrders(ctx, params)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestCreateOrder(t *testing.T) {
	type Param struct {
		pair         string
		amount       string
		price        int
		side         string
		_type        string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", amount: "0.1", price: 100000, side: "buy", _type: "limit", jsonResponse: testutil.CreateOrderJsonResponse()},
			expect: Expect{path: "/user/spot/order", method: "POST", body: testutil.ExpectedCreateOrderBody(), e: testutil.ExpectedCreateOrderModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		r, err := client.CreateOrder(ctx, c.param.pair, c.param.amount, c.param.price, c.param.side, c.param._type)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestCancelOrder(t *testing.T) {
	type Param struct {
		pair         string
		orderID      int
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", orderID: 12345, jsonResponse: testutil.CancelOrderJsonResponse()},
			expect: Expect{path: "/user/spot/cancel_order", method: "POST", body: testutil.ExpectedCancelOrderBody(), e: testutil.ExpectedCancelOrderModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		r, err := client.CancelOrder(ctx, c.param.pair, c.param.orderID)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestCancelOrders(t *testing.T) {
	type Param struct {
		pair         string
		orderIDs     []int
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Orders
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", orderIDs: []int{1, 2, 3, 4, 5}, jsonResponse: testutil.CancelOrdersJsonResponse()},
			expect: Expect{path: "/user/spot/cancel_orders", method: "POST", body: testutil.ExpectedCancelOrdersBody(), e: testutil.ExpectedCancelOrdersModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		r, err := client.CancelOrders(ctx, c.param.pair, c.param.orderIDs)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestGetOrdersInfo(t *testing.T) {
	type Param struct {
		pair         string
		orderIDs     []int
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Orders
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", orderIDs: []int{1, 2, 3, 4, 5}, jsonResponse: testutil.GetOrdersInfoJsonResponse()},
			expect: Expect{path: "/user/spot/orders_info", method: "POST", body: testutil.ExpectedGetOrdersInfoBody(), e: testutil.ExpectedGetOrdersInfoModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		r, err := client.GetOrdersInfo(ctx, c.param.pair, c.param.orderIDs)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}
