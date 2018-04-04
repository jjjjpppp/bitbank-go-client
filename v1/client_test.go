package bitbank

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/jjjjpppp/bitbank-go-client/v1/testutil"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	type Param struct {
		apiToken  string
		apiSecret string
		logger    *log.Logger
	}
	type Expect struct {
		client *Client
		err    error
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{apiToken: "", apiSecret: "secret", logger: nil},
			expect: Expect{client: nil, err: fmt.Errorf("apiTokenID is not set")},
		},
		// test case 2
		{
			param:  Param{apiToken: "apiToken", apiSecret: "", logger: nil},
			expect: Expect{client: nil, err: fmt.Errorf("apiSecret is not set")},
		},
		// test case 3
		{
			param: Param{apiToken: "apiToken", apiSecret: "apiSecret", logger: nil},
			expect: Expect{client: &Client{
				ApiTokenID: "apiToken",
				ApiSecret:  "apiSecret",
				HTTPClient: &http.Client{Timeout: time.Duration(10) * time.Second},
				Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
			}, err: nil},
		},
	}
	for _, c := range cases {
		client, e := NewClient(c.param.apiToken, c.param.apiSecret, c.param.logger)
		if client == nil && e.Error() != c.expect.err.Error() {
			t.Errorf("Worng err. test set is %+v", c)
		}
		if client == nil {
			t.Logf("client is nil. skip this case.: test set: %+v", c)
			continue
		}
		if client.ApiTokenID != c.expect.client.ApiTokenID {
			t.Errorf("Worng apiToken. test set: %+v", c)
		}
		if client.ApiSecret != c.expect.client.ApiSecret {
			t.Errorf("Worng ApiSecret. test set: %+v", c)
		}
		if reflect.TypeOf(client.HTTPClient) != reflect.TypeOf(c.expect.client.HTTPClient) {
			t.Errorf("Worng HTTPClient. test set: %+v", c)
		}
		if reflect.TypeOf(client.Logger) != reflect.TypeOf(c.expect.client.Logger) {
			t.Errorf("Worng Logger. test set: %+v", c)
		}
	}
}
func TestNewRequest(t *testing.T) {
	type Param struct {
		method     string
		spath      string
		queryParam *map[string]string
	}
	type Expect struct {
		method string
		url    string
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{method: "GET", spath: "btc_jpy/ticker", queryParam: nil},
			expect: Expect{method: "GET", url: "https://public.bitbank.cc/btc_jpy/ticker"},
		},
		// test case 2
		{
			param:  Param{method: "GET", spath: "btc_jpy/ticker", queryParam: &map[string]string{"product_id": "1", "limit": "1", "page": "1"}},
			expect: Expect{method: "GET", url: "https://public.bitbank.cc/btc_jpy/ticker?limit=1&page=1&product_id=1"},
		},
	}

	for _, c := range cases {
		client, _ := NewClient("apiTokenID", "secret", nil)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		req, _ := client.newRequest(ctx, c.param.method, c.param.spath, nil, c.param.queryParam)
		if req.Method != c.expect.method {
			t.Errorf("Worng method. case: %+v", c)
		}
		if len(req.Header["Access-Key"]) < 1 {
			t.Errorf("Worng Header No ACCESS-KEY. case: %+v", c)
		}
		if len(req.Header["Access-Nonce"]) < 1 {
			t.Errorf("Worng Header No ACCESS-NONCE. case: %+v", c)
		}
		if len(req.Header["Access-Signature"]) < 1 {
			t.Errorf("Worng Header no ACCWSS-SIGNATURE. case: %+v", c)
		}
		if req.URL.String() != c.expect.url {
			t.Errorf("Worng URL case: %+v, actual: %+v", c, req.URL.String())
		}
	}
}

func TestGetTicker(t *testing.T) {
	type Param struct {
		pair         string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Ticker
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", jsonResponse: testutil.GetTickerJsonResponse()},
			expect: Expect{path: "/btc_jpy/ticker", method: "GET", body: "", e: testutil.ExpectedGetTickerModel()},
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
		r, err := client.GetTicker(ctx, c.param.pair)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestGetTransactions(t *testing.T) {
	type Param struct {
		pair         string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Transactions
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", jsonResponse: testutil.GetTransactionJsonResponse()},
			expect: Expect{path: "/btc_jpy/transactions", method: "GET", body: "", e: testutil.ExpectedGetTransactionModel()},
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
		r, err := client.GetTransactions(ctx, c.param.pair)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestGetTransactionsByYMD(t *testing.T) {
	type Param struct {
		pair         string
		ymd          string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Transactions
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", ymd: "20180402", jsonResponse: testutil.GetTransactionJsonResponse()},
			expect: Expect{path: "/btc_jpy/transactions/20180402", method: "GET", body: "", e: testutil.ExpectedGetTransactionModel()},
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
		r, err := client.GetTransactionsByYMD(ctx, c.param.pair, c.param.ymd)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestGetCandlesticks(t *testing.T) {
	type Param struct {
		pair         string
		candleType   string
		ymd          string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Candlesticks
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{pair: "btc_jpy", candleType: "30min", ymd: "20180402", jsonResponse: testutil.GetCandlesticsJsonResponse()},
			expect: Expect{path: "/btc_jpy/candlestick/30min/20180402", method: "GET", body: "", e: testutil.ExpectedGetCandlesticsModel()},
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
		r, err := client.GetCandlesticks(ctx, c.param.pair, c.param.candleType, c.param.ymd)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}
