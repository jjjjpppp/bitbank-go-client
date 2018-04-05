package bitbank

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
	"github.com/jjjjpppp/bitbank-go-client/v1/testutil"
	"testing"
	"time"
)

func TestGetAssets(t *testing.T) {
	type Param struct {
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		e      *models.Assets
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetAssetsJsonResponse()},
			expect: Expect{path: "/user/assets", method: "GET", body: "", e: testutil.ExpectedGetAssetsModel()},
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
		r, err := client.GetAssets(ctx)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}
