package models

import (
	"encoding/json"
	"sort"
)

type Depth struct {
	Success int       `json:"success"`
	Data    DepthData `json:"data"`
}

type DepthData struct {
	Asks [][]json.Number `json:"asks"`
	Bids [][]json.Number `json:"bids"`
}

func (p *Depth) GetAsksFloat64() [][]float64 {
	var asksFloat64 [][]float64
	for _, ask := range p.Data.Asks {
		a, _ := ask[0].Float64()
		b, _ := ask[1].Float64()
		asksFloat64 = append(asksFloat64, []float64{a, b})
	}
	return asksFloat64
}

func (p *Depth) GetBidsFloat64() [][]float64 {
	var bidsFloat64 [][]float64
	for _, bid := range p.Data.Bids {
		a, _ := bid[0].Float64()
		b, _ := bid[1].Float64()
		bidsFloat64 = append(bidsFloat64, []float64{a, b})
	}
	return bidsFloat64
}

func (p *Depth) SortAsksByQuontity() [][]float64 {
	sortAsks := p.GetAsksFloat64()
	sort.Slice(sortAsks, func(i, j int) bool {
		return sortAsks[i][1] > sortAsks[j][1]
	})
	return sortAsks
}

func (p *Depth) SortBidsByQuontity() [][]float64 {
	sortBids := p.GetBidsFloat64()
	sort.Slice(sortBids, func(i, j int) bool {
		return sortBids[i][1] > sortBids[j][1]
	})
	return sortBids
}

func (p *Depth) SortAsksByPrice(order string) [][]float64 {
	sortAsks := p.GetAsksFloat64()
	if order == "asc" {
		sort.Slice(sortAsks, func(i, j int) bool {
			return sortAsks[i][0] < sortAsks[j][0]
		})
	} else {
		sort.Slice(sortAsks, func(i, j int) bool {
			return sortAsks[i][0] > sortAsks[j][0]
		})
	}
	return sortAsks
}

func (p *Depth) SortBidsByPrice(order string) [][]float64 {
	sortBids := p.GetBidsFloat64()
	if order == "asc" {
		sort.Slice(sortBids, func(i, j int) bool {
			return sortBids[i][0] < sortBids[j][0]
		})
	} else {
		sort.Slice(sortBids, func(i, j int) bool {
			return sortBids[i][0] > sortBids[j][0]
		})
	}
	return sortBids
}
