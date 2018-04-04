package testutil

import (
	"encoding/json"
	"fmt"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GenerateTestServer(t *testing.T, expectPath string, expectMethod string, expectBody string, jsonResponse string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.RequestURI() != expectPath {
				t.Errorf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), expectPath)
			}
			if r.Method != expectMethod {
				t.Errorf("worng Method. actual:%+v, expect:%+v", r.Method, expectMethod)
			}
			// Read body
			if expectBody != "" {
				b, err := ioutil.ReadAll(r.Body)
				s := string(b)
				defer r.Body.Close()
				if err != nil {
					t.Errorf("Worng body. err:%+v", err)
				}
				if s != expectBody {
					t.Errorf("Worng body. actual: %+v, expect:%+v", s, expectBody)
				}
			}

			// set expected json
			w.Header().Set("content-Type", "text")
			fmt.Fprintf(w, jsonResponse)
			return
		},
	))
}

func GetTickerJsonResponse() string {
	return `
{
  "success": 1,
  "data": {
    "sell": "string",
    "buy": "string",
    "high": "string",
    "low": "string",
    "last": "string",
    "vol": "string",
    "timestamp": 0
  }
}`
}

func ExpectedGetTickerModel() *models.Ticker {
	d := &models.TickerData{
		Sell:      "string",
		Buy:       "string",
		High:      "string",
		Low:       "string",
		Last:      "string",
		Vol:       "string",
		Timestamp: 0,
	}
	return &models.Ticker{Success: 1, Data: d}
}

func GetTransactionJsonResponse() string {
	return `
{
  "success": 1,
  "data": {
    "transactions": [
      {
        "transaction_id": 0,
        "side": "string",
        "price": "string",
        "amount": "string",
        "executed_at": 0
      }
    ]
  }
}`
}

func ExpectedGetTransactionModel() *models.Transactions {
	m1 := &models.Transaction{
		TransactionID: 0,
		Side:          "string",
		Price:         "string",
		Amount:        "string",
		ExecutedAt:    0,
	}
	data := &models.TransactionsData{Transactions: []*models.Transaction{m1}}
	return &models.Transactions{Success: 1, Data: data}
}

func GetCandlesticsJsonResponse() string {
	return `
{
  "success": 1,
  "data": {
    "candlestick": [
      {
        "type": "30min",
        "ohlcv": [
          [
            "100","200","50","150","1.11",12345678
          ]
        ]
      }
    ]
  }
}
`
}

func ExpectedGetCandlesticsModel() *models.Candlesticks {
	candlestick := &models.Candlestick{
		Type:  "30min",
		Ohlcv: [][]json.Number{{"100", "200", "50", "150", "1.11", "12345678"}},
	}
	data := &models.CandlesticksData{
		Candlesticks: []*models.Candlestick{candlestick},
	}
	return &models.Candlesticks{Success: 1, Data: data}
}
