package data

import (
	"github.com/cu2koo/go-ra-data/model"
	"time"
)

type Market interface {
	GetPrices(symbols *[]string, period *time.Duration) ([]*model.PriceData, error)
}
