package bitbank

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/jjjjpppp/bitbank-go-client/v1/testutil"
	"testing"
	"time"
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
