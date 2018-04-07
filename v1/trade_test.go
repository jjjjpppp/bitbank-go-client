package bitbank

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/jjjjpppp/bitbank-go-client/v1/testutil"
	"testing"
	"time"
)

func TestGetTrades(t *testing.T) {
	type Param struct {
		pair         string
		count        float64
		fromID       float64
		endID        float64
		since        float64
		end          float64
		order        string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Trades
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", count: 1, fromID: 1, endID: 10, since: 12345, end: 12355, order: "asc", jsonResponse: testutil.GetTradesJsonResponse()},
			expect: Expect{path: "/user/spot/trade_history?count=1&end=12355&end_id=10&from_id=1&order=asc&pair=btc_jpy&since=12345", method: "GET", body: "", e: testutil.ExpectedGetTradesModel()},
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
		r, err := client.GetTrades(ctx, c.param.pair, c.param.count, c.param.fromID, c.param.endID, c.param.since, c.param.end, c.param.order)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}
