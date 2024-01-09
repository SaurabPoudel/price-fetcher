package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return s.MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC":  20_000.0,
	"ETH":  1_000.0,
	"DOGE": 0.01,
	"SP":   1000_000.0,
}

func (s *priceFetcher) MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}
