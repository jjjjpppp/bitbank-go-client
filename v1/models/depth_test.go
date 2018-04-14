package models

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetAsksFloat64(t *testing.T) {
	type Param struct {
		depth Depth
	}
	type Expect struct {
		e [][]float64
	}
	type Case struct {
		param  Param
		expect Expect
	}
	// test case 1
	c1 := Case{
		param: Param{
			depth: Depth{
				Success: 1,
				Data: DepthData{
					Asks: [][]json.Number{{"0.1", "0.2"}, {"0.3", "0.4"}},
					Bids: [][]json.Number{{"0.5", "0.6"}, {"0.7", "0.8"}},
				},
			},
		},
		expect: Expect{
			e: [][]float64{{0.1, 0.2}, {0.3, 0.4}},
		},
	}
	cases := []Case{c1}

	for _, c := range cases {
		r := c.param.depth.GetAsksFloat64()

		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestGetBidsFloat64(t *testing.T) {
	type Param struct {
		depth Depth
	}
	type Expect struct {
		e [][]float64
	}
	type Case struct {
		param  Param
		expect Expect
	}
	// test case 1
	c1 := Case{
		param: Param{
			depth: Depth{
				Success: 1,
				Data: DepthData{
					Asks: [][]json.Number{{"0.1", "0.2"}, {"0.3", "0.4"}},
					Bids: [][]json.Number{{"0.5", "0.6"}, {"0.7", "0.8"}},
				},
			},
		},
		expect: Expect{
			e: [][]float64{{0.5, 0.6}, {0.7, 0.8}},
		},
	}
	cases := []Case{c1}

	for _, c := range cases {
		r := c.param.depth.GetBidsFloat64()

		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestSortAsksByPrice(t *testing.T) {
	type Param struct {
		depth Depth
		order string
	}
	type Expect struct {
		e [][]float64
	}
	type Case struct {
		param  Param
		expect Expect
	}
	// test case 1
	c1 := Case{
		param: Param{
			depth: Depth{
				Success: 1,
				Data: DepthData{
					Asks: [][]json.Number{{"0.1", "0.2"}, {"0.8", "0.9"}, {"0.3", "0.4"}},
					Bids: [][]json.Number{{"0.5", "0.6"}, {"0.7", "0.8"}},
				},
			},
			order: "asc",
		},
		expect: Expect{
			e: [][]float64{{0.1, 0.2}, {0.3, 0.4}, {0.8, 0.9}},
		},
	}
	c2 := Case{
		param: Param{
			depth: Depth{
				Success: 1,
				Data: DepthData{
					Asks: [][]json.Number{{"0.1", "0.2"}, {"0.8", "0.9"}, {"0.3", "0.4"}},
					Bids: [][]json.Number{{"0.5", "0.6"}, {"0.7", "0.8"}},
				},
			},
			order: "desc",
		},
		expect: Expect{
			e: [][]float64{{0.8, 0.9}, {0.3, 0.4}, {0.1, 0.2}},
		},
	}
	cases := []Case{c1, c2}

	for _, c := range cases {
		r := c.param.depth.SortAsksByPrice(c.param.order)

		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestSortBidsByPrice(t *testing.T) {
	type Param struct {
		depth Depth
		order string
	}
	type Expect struct {
		e [][]float64
	}
	type Case struct {
		param  Param
		expect Expect
	}
	// test case 1
	c1 := Case{
		param: Param{
			depth: Depth{
				Success: 1,
				Data: DepthData{
					Asks: [][]json.Number{{"0.1", "0.2"}, {"0.3", "0.4"}},
					Bids: [][]json.Number{{"0.5", "0.6"}, {"0.75", "0.85"}, {"0.7", "0.8"}},
				},
			},
			order: "asc",
		},
		expect: Expect{
			e: [][]float64{{0.5, 0.6}, {0.7, 0.8}, {0.75, 0.85}},
		},
	}
	c2 := Case{
		param: Param{
			depth: Depth{
				Success: 1,
				Data: DepthData{
					Asks: [][]json.Number{{"0.1", "0.2"}, {"0.3", "0.4"}},
					Bids: [][]json.Number{{"0.5", "0.6"}, {"0.75", "0.85"}, {"0.7", "0.8"}},
				},
			},
			order: "desc",
		},
		expect: Expect{
			e: [][]float64{{0.75, 0.85}, {0.7, 0.8}, {0.5, 0.6}},
		},
	}
	cases := []Case{c1, c2}

	for _, c := range cases {
		r := c.param.depth.SortBidsByPrice(c.param.order)

		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}
