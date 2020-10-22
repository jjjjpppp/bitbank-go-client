package testutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jjjjpppp/bitbank-go-client/v1/models"
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

func GetAssetsJsonResponse() string {
	return `
{
  "success": 1,
  "data": {
    "assets": [
      {
        "asset": "jpy",
        "amount_precision": 4,
        "onhand_amount": "0.8451",
        "locked_amount": "0.0000",
        "free_amount": "0.8451",
        "stop_deposit": false,
        "stop_withdrawal": false,
        "withdrawal_fee": {
          "threshold": "30000.0000",
          "under": "540.0000",
          "over": "756.0000"
        }
      },
      {
        "asset": "btc",
        "amount_precision": 8,
        "onhand_amount": "0.00000000",
        "locked_amount": "0.00000000",
        "free_amount": "0.00000000",
        "stop_deposit": false,
        "stop_withdrawal": false,
        "withdrawal_fee": "0.00100000"
      }
		]
  }
}
`
}

func ExpectedGetAssetsModel() *models.Assets {
	w1 := &models.WithdrawalFee{
		Threshold: "30000.0000",
		Under:     "540.0000",
		Over:      "756.0000",
	}
	w2 := &models.WithdrawalFee{
		Fee: "0.00100000",
	}

	a1 := &models.Asset{
		Asset:           "jpy",
		AmountPrecision: 4,
		OnhandAmount:    "0.8451",
		LockedAmount:    "0.0000",
		FreeAmount:      "0.8451",
		StopDeposit:     false,
		StopWithdrawal:  false,
		WithdrawalFee:   w1,
	}
	a2 := &models.Asset{
		Asset:           "btc",
		AmountPrecision: 8,
		OnhandAmount:    "0.00000000",
		LockedAmount:    "0.00000000",
		FreeAmount:      "0.00000000",
		StopDeposit:     false,
		StopWithdrawal:  false,
		WithdrawalFee:   w2,
	}
	ad := &models.AssetsData{Assets: []*models.Asset{a1, a2}}
	return &models.Assets{Success: 1, Data: ad}

}

func GetOrderJsonResponse() string {
	return `
{
  "success": 1,
  "data": {
    "order_id": 0,
    "pair": "string",
    "side": "string",
    "type": "string",
    "start_amount": "string",
    "remaining_amount": "string",
    "executed_amount": "string",
    "price": "string",
    "average_price": "string",
    "ordered_at": 0,
    "status": "string"
  }
}
`
}

func ExpectedGetOrderModel() *models.Order {
	data := &models.OrderData{
		OrderID:         0,
		Pair:            "string",
		Side:            "string",
		Type:            "string",
		StartAmount:     "string",
		RemainingAmount: "string",
		ExecutedAmount:  "string",
		Price:           "string",
		AveragePrice:    "string",
		OrderedAt:       0,
		Status:          "string",
	}
	return &models.Order{Success: 1, Data: data}
}

func GetOrdersJsonResponse() string {
	return `{
  "success": 1,
  "data": {
    "orders": [
      {
        "order_id": 0,
        "pair": "string",
        "side": "string",
        "type": "string",
        "start_amount": "string",
        "remaining_amount": "string",
        "executed_amount": "string",
        "price": "string",
        "average_price": "string",
        "ordered_at": 0,
        "status": "string"
      }
    ]
  }
}`
}

func ExpectedGetOrdersModel() *models.Orders {
	od := &models.OrderData{
		OrderID:         0,
		Pair:            "string",
		Side:            "string",
		Type:            "string",
		StartAmount:     "string",
		RemainingAmount: "string",
		ExecutedAmount:  "string",
		Price:           "string",
		AveragePrice:    "string",
		OrderedAt:       0,
		Status:          "string",
	}
	ods := &models.OrdersData{
		Orders: []*models.OrderData{od},
	}
	return &models.Orders{Success: 1, Data: ods}
}

func CreateOrderJsonResponse() string {
	return `{
  "success": 1,
  "data": {
    "order_id": 0,
    "pair": "string",
    "side": "string",
    "type": "string",
    "start_amount": "string",
    "remaining_amount": "string",
    "executed_amount": "string",
    "price": "string",
    "average_price": "string",
    "ordered_at": 0,
    "status": "string"
  }
}`
}

func ExpectedCreateOrderBody() string {
	return `{
  "pair": "btc_jpy",
  "amount": "0.1",
  "price": 100000,
  "side": "buy",
  "type": "limit"
}`
}

func ExpectedCreateOrderModel() *models.Order {
	data := &models.OrderData{
		OrderID:         0,
		Pair:            "string",
		Side:            "string",
		Type:            "string",
		StartAmount:     "string",
		RemainingAmount: "string",
		ExecutedAmount:  "string",
		Price:           "string",
		AveragePrice:    "string",
		OrderedAt:       0,
		Status:          "string",
	}
	return &models.Order{Success: 1, Data: data}
}

func CancelOrderJsonResponse() string {
	return `{
  "success": 1,
  "data": {
    "order_id": 0,
    "pair": "string",
    "side": "string",
    "type": "string",
    "start_amount": "string",
    "remaining_amount": "string",
    "executed_amount": "string",
    "price": "string",
    "average_price": "string",
    "ordered_at": 0,
    "status": "string"
  }
}`
}

func ExpectedCancelOrderBody() string {
	return `{
  "pair":"btc_jpy",
  "order_id":12345
}`
}

func ExpectedCancelOrderModel() *models.Order {
	data := &models.OrderData{
		OrderID:         0,
		Pair:            "string",
		Side:            "string",
		Type:            "string",
		StartAmount:     "string",
		RemainingAmount: "string",
		ExecutedAmount:  "string",
		Price:           "string",
		AveragePrice:    "string",
		OrderedAt:       0,
		Status:          "string",
	}
	return &models.Order{Success: 1, Data: data}
}

func CancelOrdersJsonResponse() string {
	return `{
  "success": 1,
  "data": {
    "orders": [
      {
        "order_id": 0,
        "pair": "string",
        "side": "string",
        "type": "string",
        "start_amount": "string",
        "remaining_amount": "string",
        "executed_amount": "string",
        "price": "string",
        "average_price": "string",
        "ordered_at": 0,
        "status": "string"
      }
    ]
  }
}`
}

func ExpectedCancelOrdersBody() string {
	return `{
  "pair":"btc_jpy",
  "order_ids":[1,2,3,4,5]
}`
}

func ExpectedCancelOrdersModel() *models.Orders {
	od := &models.OrderData{
		OrderID:         0,
		Pair:            "string",
		Side:            "string",
		Type:            "string",
		StartAmount:     "string",
		RemainingAmount: "string",
		ExecutedAmount:  "string",
		Price:           "string",
		AveragePrice:    "string",
		OrderedAt:       0,
		Status:          "string",
	}
	ods := &models.OrdersData{
		Orders: []*models.OrderData{od},
	}
	return &models.Orders{Success: 1, Data: ods}
}

func GetOrdersInfoJsonResponse() string {
	return `{
  "success": 1,
  "data": {
    "orders": [
      {
        "order_id": 0,
        "pair": "string",
        "side": "string",
        "type": "string",
        "start_amount": "string",
        "remaining_amount": "string",
        "executed_amount": "string",
        "price": "string",
        "average_price": "string",
        "ordered_at": 0,
        "status": "string"
      }
    ]
  }
}`
}

func ExpectedGetOrdersInfoBody() string {
	return `{
  "pair":"btc_jpy",
  "order_ids":[1,2,3,4,5]
}`
}

func ExpectedGetOrdersInfoModel() *models.Orders {
	od := &models.OrderData{
		OrderID:         0,
		Pair:            "string",
		Side:            "string",
		Type:            "string",
		StartAmount:     "string",
		RemainingAmount: "string",
		ExecutedAmount:  "string",
		Price:           "string",
		AveragePrice:    "string",
		OrderedAt:       0,
		Status:          "string",
	}
	ods := &models.OrdersData{
		Orders: []*models.OrderData{od},
	}
	return &models.Orders{Success: 1, Data: ods}
}

func GetTradesJsonResponse() string {
	return `{
  "success": 1,
  "data": {
    "trades": [
      {
        "trade_id": 0,
        "pair": "string",
        "order_id": 0,
        "side": "string",
        "type": "string",
        "amount": "string",
        "price": "string",
        "maker_taker": "string",
        "fee_amount_base": "string",
        "fee_amount_quote": "string",
        "executed_at": 0
      }
    ]
  }
}
`
}

func ExpectedGetTradesModel() *models.Trades {
	m1 := &models.Trade{
		TradeID:        0,
		Pair:           "string",
		OrderID:        0,
		Side:           "string",
		Type:           "string",
		Amount:         "string",
		Price:          "string",
		MakerTaker:     "string",
		FeeAmountBase:  "string",
		FeeAmountQuote: "string",
		ExecutedAt:     0,
	}

	d := &models.TradesData{
		Trades: []*models.Trade{m1},
	}

	return &models.Trades{Success: 1, Data: d}
}

func GetWithdrawalAccountsJsonResponse() string {
	return `{
  "success": 1,
  "data": {
    "accounts": [
      {
        "uuid": "string",
        "label": "string",
        "address": "string"
      }
    ]
  }
}
`
}

func ExpectedGetWithdrawalAccountsModel() *models.WithdrawalAccounts {
	a := &models.Account{
		UUID:    "string",
		Label:   "string",
		Address: "string",
	}
	w := &models.WithdrawalAccountsData{
		Accounts: []*models.Account{a},
	}
	return &models.WithdrawalAccounts{Success: 1, Data: w}
}

func RequestWithdrawalJsonResponse() string {
	return `{
  "success": 0,
  "data": {
    "uuid": "string",
    "asset": "string",
    "amount": 0,
    "account_uuid": "string",
    "fee": "string",
    "status": "string",
    "label": "string",
    "txid": "string",
    "address": "string"
  }
}`
}

func ExpectedRequestWithdrawalBody() string {
	return `{
  "asset":"btc",
  "uuid":"12345",
  "amount":"100",
  "otp_token":"aaaa",
  "sms_token":"bbbb"
}`
}

func ExpectedRequestWithdrawalModel() *models.RequestWithdrawal {

	m := &models.RequestWithdrawalData{
		UUID:        "string",
		Asset:       "string",
		Amount:      0,
		AccountUUID: "string",
		Fee:         "string",
		Status:      "string",
		Label:       "string",
		Txid:        "string",
		Address:     "string",
	}
	return &models.RequestWithdrawal{Success: 0, Data: m}
}
