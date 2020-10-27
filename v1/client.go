/*
  Package bitbank is a API Client of Bitbank.cc

  API document https://docs.bitbank.cc/

  Error codes https://docs.bitbank.cc/error_code/
*/
package bitbank

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/jjjjpppp/bitbank-go-client/v1/models"
)

const (
	// baseUrl is a public API URL
	baseUrl string = "https://public.bitbank.cc"
	// privateBaseUrl is a private API URL
	privateBaseUrl string = "https://api.bitbank.cc/v1"
	// version of this software
	version string = "0.0.1"
)

// Main client struct of this package
type Client struct {
	URL        *url.URL
	PrivateURL *url.URL
	ApiTokenID string
	ApiSecret  string
	HTTPClient *http.Client
	Logger     *log.Logger
	testServer *httptest.Server
}

// Generate Client with token ID and secret. you can find these infomation in bitbank site.
func NewClient(apiTokenID string, apiSecret string, logger *log.Logger) (*Client, error) {
	if len(apiTokenID) == 0 {
		return nil, fmt.Errorf("apiTokenID is not set")
	}

	if len(apiSecret) == 0 {
		return nil, fmt.Errorf("apiSecret is not set")
	}

	publicUrl, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, err
	}
	privateUrl, err := url.ParseRequestURI(privateBaseUrl)
	if err != nil {
		return nil, err
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	return &Client{URL: publicUrl, PrivateURL: privateUrl, ApiTokenID: apiTokenID, ApiSecret: apiSecret, HTTPClient: client, Logger: logger}, nil

}

// Get Ticker https://docs.bitbank.cc/#!/Ticker/ticker
func (c *Client) GetTicker(ctx context.Context, pair string) (*models.Ticker, error) {
	spath := fmt.Sprintf("/%s/ticker", pair)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var ticker models.Ticker
	if err := decodeBody(res, &ticker); err != nil {
		return nil, err
	}

	return &ticker, nil
}

// Get Depth https://docs.bitbank.cc/#!/Depth/depth
func (c *Client) GetDepth(ctx context.Context, pair string) (*models.Depth, error) {
	spath := fmt.Sprintf("/%s/depth", pair)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var depth models.Depth
	if err := decodeBody(res, &depth); err != nil {
		return nil, err
	}

	return &depth, nil
}

// Get Transactions https://docs.bitbank.cc/#!/Transactions/transactions
func (c *Client) GetTransactions(ctx context.Context, pair string) (*models.Transactions, error) {
	spath := fmt.Sprintf("/%s/transactions", pair)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var transaction models.Transactions
	if err := decodeBody(res, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Get Transactions by YYYYMMDD https://docs.bitbank.cc/#!/Transactions/transactions_YYYYMMDD
func (c *Client) GetTransactionsByYMD(ctx context.Context, pair, ymdString string) (*models.Transactions, error) {
	spath := fmt.Sprintf("/%s/transactions/%s", pair, ymdString)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var transaction models.Transactions
	if err := decodeBody(res, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Get Candlesticks https://docs.bitbank.cc/#!/Candlestick/candlestick
func (c *Client) GetCandlesticks(ctx context.Context, pair, candleType, ymdString string) (*models.Candlesticks, error) {
	spath := fmt.Sprintf("/%s/candlestick/%s/%s", pair, candleType, ymdString)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var candlesticks models.Candlesticks
	if err := decodeBody(res, &candlesticks); err != nil {
		return nil, err
	}

	return &candlesticks, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader, queryParam *map[string]string) (*http.Request, error) {

	var u url.URL
	// swith client url for unit test
	if c.testServer != nil {
		testUrl, _ := url.ParseRequestURI(c.testServer.URL)
		u = *testUrl
	} else if isRequestPublic(spath) {
		u = *c.URL
	} else {
		u = *c.PrivateURL
	}

	u.Path = path.Join(u.Path, spath)

	// build QueryParameter
	if queryParam != nil {
		q := u.Query()
		for k, v := range *queryParam {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	userAgent := fmt.Sprintf("GoClient/%s (%s)", version, runtime.Version())
	accessNonce := fmt.Sprintf("%d", time.Now().Unix())
	var tokenString string
	var b bytes.Buffer
	if method == "GET" {
		tokenString = makeHMAC(accessNonce+"/v1"+spath+u.RawQuery, c.ApiSecret)
	} else {
		io.Copy(&b, body)
		tokenString = makeHMAC(accessNonce+b.String(), c.ApiSecret)
	}

	req, err := http.NewRequest(method, u.String(), &b)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("ACCESS-KEY", c.ApiTokenID)
	req.Header.Set("ACCESS-NONCE", accessNonce)
	req.Header.Set("ACCESS-SIGNATURE", tokenString)

	return req, nil
}

func (c *Client) sendRequest(ctx context.Context, method, spath string, body io.Reader, queryParam *map[string]string) (*http.Response, error) {
	req, err := c.newRequest(ctx, method, spath, body, queryParam)
	c.Logger.Printf("Request:  %s \n", httpRequestLog(req))
	if err != nil {
		c.Logger.Printf("err: %#v \n", err)
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	c.Logger.Printf("Response: %s \n", httpResponseLog(res))
	if err != nil {
		c.Logger.Printf("err: %#v \n", err)
		return nil, err
	}

	if res.StatusCode != 200 {
		c.Logger.Printf("err: %#v \n", err)
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}
	return res, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func httpResponseLog(resp *http.Response) string {
	b, _ := httputil.DumpResponse(resp, true)
	return string(b)
}
func httpRequestLog(req *http.Request) string {
	b, _ := httputil.DumpRequest(req, true)
	return string(b)
}

func isRequestPublic(path string) bool {
	if strings.Contains(path, "/ticker") || strings.Contains(path, "/depth") || strings.Contains(path, "/transactions") || strings.Contains(path, "/candlestick") {
		return true
	}
	return false
}

func makeHMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
