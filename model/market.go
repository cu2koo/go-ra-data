package model

type Price struct {
	Date   *string  `json:"date"`
	Open   *float64 `json:"open"`
	High   *float64 `json:"high"`
	Low    *float64 `json:"low"`
	Close  *float64 `json:"close"`
	Volume *float64 `json:"volume"`
}

type PriceData struct {
	Ticker        *string  `json:"ticker"`
	Currency      *string  `json:"currency"`
	Prices        []*Price `json:"prices"`
	PriceHigh52Wk *float64 `json:"priceHigh52Wk"`
	PriceLow52Wk  *float64 `json:"priceLow52Wk"`
	Source        *string  `json:"source"`
}
