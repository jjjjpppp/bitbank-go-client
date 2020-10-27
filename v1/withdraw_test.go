package bitbank

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/jjjjpppp/bitbank-go-client/v1/request"
	"github.com/jjjjpppp/bitbank-go-client/v1/testutil"
)

func TestGetWithdrawalAccounts(t *testing.T) {
	type Param struct {
		asset        string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.WithdrawalAccounts
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{asset: "btc", jsonResponse: testutil.GetWithdrawalAccountsJsonResponse()},
			expect: Expect{path: "/user/withdrawal_account?asset=btc", method: "GET", body: "", e: testutil.ExpectedGetWithdrawalAccountsModel()},
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
		r, err := client.GetWithdrawalAccounts(ctx, c.param.asset)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestRequestWithdrawal(t *testing.T) {
	type Param struct {
		asset        string
		uuid         string
		amount       string
		otpToken     string
		smsToken     string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.RequestWithdrawal
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{asset: "btc", uuid: "12345", amount: "100", otpToken: "aaaa", smsToken: "bbbb", jsonResponse: testutil.RequestWithdrawalJsonResponse()},
			expect: Expect{path: "/user/request_withdrawal", method: "POST", body: testutil.ExpectedRequestWithdrawalBody(), e: testutil.ExpectedRequestWithdrawalModel()},
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
		params := request.RequestWithdrawalParams{c.param.asset, c.param.uuid, c.param.amount, c.param.otpToken, c.param.smsToken}
		r, err := client.RequestWithdrawal(ctx, params)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}
