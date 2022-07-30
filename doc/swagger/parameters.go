package swagger

// swagger:parameters createUser updateUser
type user struct {
	// required: true
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

// swagger:parameters updateFinancialInformation
type financialInformation struct {
	// required: true
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

// swagger:parameters updateGoal
type goal struct {
	// required: true
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

// swagger:parameters updatePeriodicalDeposits
type periodicalDeposits struct {
	// required: true
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

// swagger:parameters updateRiskAssessment
type riskAssessment struct {
	// required: true
	// in: body
	Body struct {
		// Example: 5
		RiskAppetite *int `json:"riskAppetite"`
		// Example: 10.5
		MaxVolatility *float64 `json:"maxVolatility"`
	}
}

// swagger:parameters addExperience updateExperience
type experience struct {
	// required: true
	// in: body
	Body struct {
		// Example: 16
		Id *int `json:"id"`
		// Example: Stocks
		Type *string `json:"type"`
		// Example: 5
		Transactions *int `json:"transactions"`
		// Example: 3y
		TimePeriod *string `json:"timePeriod"`
	}
}

// swagger:parameters addDeposit updateDeposit
type deposit struct {
	// required: true
	// in: body
	Body struct {
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
}

// swagger:parameters addSecurity updateSecurity
type security struct {
	// required: true
	// in: body
	Body struct {
		// Example: 16
		Id *int `json:"id"`
		// Example: MSFT
		Symbol *string `json:"symbol"`
		// Example: 4.00
		Shares *float64 `json:"shares"`
		// Example: 280.00
		AverageCostPerShare *float64 `json:"averageCostPerShare"`
	}
}
