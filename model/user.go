package model

type User struct {
	UserId   *int    `json:"userId"`
	Mail     *string `json:"mail"`
	Password *string `json:"password"`
}

type Users []*User

type Assets struct {
	Liquid   *float64 `json:"liquid"`
	Illiquid *float64 `json:"illiquid"`
}

type FinancialInformation struct {
	Assets      *Assets  `json:"assets"`
	Income      *float64 `json:"income"`
	Expenses    *float64 `json:"expenses"`
	Liabilities *float64 `json:"liabilities"`
}

type Goal struct {
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
	Goal      *string `json:"goal"`
}

type PeriodicalDeposits struct {
	Amount   *float64 `json:"amount"`
	Currency *string  `json:"currency"`
	Period   *string  `json:"period"`
}

type RiskAssessments struct {
	RiskAppetite  *int     `json:"riskAppetite"`
	MaxVolatility *float64 `json:"maxVolatility"`
}

type Experience struct {
	Id           *int    `json:"id"`
	Type         *string `json:"type"`
	Transactions *int    `json:"transactions"`
	TimePeriod   *string `json:"timePeriod"`
}

type Experiences []*Experience

type Deposit struct {
	Id       *int     `json:"id"`
	Date     *string  `json:"date"`
	Amount   *float64 `json:"amount"`
	Currency *string  `json:"currency"`
	Target   *string  `json:"target"`
}

type Deposits []*Deposit

type Security struct {
	Id                  *int     `json:"id"`
	Symbol              *string  `json:"symbol"`
	Shares              *float64 `json:"shares"`
	AverageCostPerShare *float64 `json:"averageCostPerShare"`
}

type Securities []*Security
