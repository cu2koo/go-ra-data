package data

import (
	"github.com/cu2koo/go-ra-data/model"
	"time"
)

// Market is a data connection interface for the market data provider.
type Market interface {
	// GetPrices handles price data retrieval.
	GetPrices(symbols *[]string, period *time.Duration) ([]*model.PriceData, error)
}
