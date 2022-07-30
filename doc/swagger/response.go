package swagger

// swagger:response userResponses
type userResponses []*struct {
	// Example: 16
	UserId *int `json:"userId"`
	// Example: max.mustermann@domain.com
	Mail *string `json:"mail"`
	// Example: encrypted_secret_password
	Password *string `json:"password"`
}

// swagger:response userResponse
type userResponse struct {
	// in: body
	Body struct {
		// Example: 16
		UserId *int `json:"userId"`
		// Example: max.mustermann@domain.com
		Mail *string `json:"mail"`
		// Example: encrypted_secret_password
		Password *string `json:"password"`
	}
}

// swagger:response financialInformationResponse
type financialInformationResponse struct {
	// in: body
	Body struct {
		Assets *struct {
			// Example: 10000.00
			Liquid *float64 `json:"liquid"`
			// Example: 50000.00
			Illiquid *float64 `json:"illiquid"`
		} `json:"assets"`
		// Example: 4000.00
		Income *float64 `json:"income"`
		// Example: 2500.00
		Expenses *float64 `json:"expenses"`
		// Example: 40000.00
		Liabilities *float64 `json:"liabilities"`
	}
}

// swagger:response goalResponse
type goalResponse struct {
	// in: body
	Body struct {
		// Example: 2022-07-01
		StartDate *string `json:"startDate"`
		// Example: 2032-07-01
		EndDate *string `json:"endDate"`
		// Example: build wealth
		Goal *string `json:"goal"`
	}
}

// swagger:response periodicalDepositsResponse
type periodicalDepositsResponse struct {
	// in: body
	Body struct {
		// Example: 25
		Amount *float64 `json:"amount"`
		// Example: USD
		Currency *string `json:"currency"`
		// Example: 1m
		Period *string `json:"period"`
	}
}

// swagger:response riskAssessmentResponse
type riskAssessmentResponse struct {
	// in: body
	Body struct {
		// Example: 5
		RiskAppetite *int `json:"riskAppetite"`
		// Example: 10.5
		MaxVolatility *float64 `json:"maxVolatility"`
	}
}

// swagger:response experienceResponses
type experienceResponses []*struct {
	// Example: 16
	Id *int `json:"id"`
	// Example: Stocks
	Type *string `json:"type"`
	// Example: 5
	Transactions *int `json:"transactions"`
	// Example: 3y
	TimePeriod *string `json:"timePeriod"`
}

// swagger:response depositResponses
type depositResponses []*struct {
	// Example: 16
	Id *int `json:"id"`
	// Example: 2022-07-01
	Date *string `json:"date"`
	// Example: 2000.00
	Amount *float64 `json:"amount"`
	// Example: USD
	Currency *string `json:"currency"`
	// Example: liquidity
	Target *string `json:"target"`
}

// swagger:response securityResponses
type securityResponses []*struct {
	// Example: 16
	Id *int `json:"id"`
	// Example: MSFT
	Symbol *string `json:"symbol"`
	// Example: 4.00
	Shares *float64 `json:"shares"`
	// Example: 280.00
	AverageCostPerShare *float64 `json:"averageCostPerShare"`
}

// swagger:response priceDataResponses
type priceDataResponses []*struct {
	// Example: AAPL
	Ticker *string `json:"ticker"`
	// Example: USD
	Currency *string `json:"currency"`
	Prices   []*struct {
		// Example: 2020-01-10
		Date *string `json:"date"`
		// Example: 120.06
		Open *float64 `json:"open"`
		// Example: 140.21
		High *float64 `json:"high"`
		// Example: 110.92
		Low *float64 `json:"low"`
		// Example: 132.75
		Close *float64 `json:"close"`
		// Example: 218000
		Volume *float64 `json:"volume"`
	} `json:"prices"`
	// Example: 150.01
	PriceHigh52Wk *float64 `json:"priceHigh52Wk"`
	// Example: 99.99
	PriceLow52Wk *float64 `json:"priceLow52Wk"`
	// Example: domain.com
	Source *string `json:"source"`
}

// swagger:response successResponse
type successResponse struct {
	// in: body
	Body struct {
		// Example: action successful
		Message *string `json:"message"`
	}
}

// swagger:response errorResponse
type errorResponse struct {
	// in: body
	Body struct {
		// Example: something went wrong
		Message *string `json:"message"`
	}
}
