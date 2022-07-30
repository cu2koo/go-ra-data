package data

import (
	"github.com/cu2koo/go-ra-data/model"
)

type User interface {
	CreateUser(user *model.User) error
	GetUser(userId *int) (*model.User, error)
	GetUsers() (*model.Users, error)
	UpdateUser(userId *int, user *model.User) error
	DeleteUser(userId *int) error

	GetFinancialInformation(userId *int) (*model.FinancialInformation, error)
	UpdateFinancialInformation(userId *int, financialInformation *model.FinancialInformation) error

	GetGoal(userId *int) (*model.Goal, error)
	UpdateGoal(userId *int, goal *model.Goal) error

	GetPeriodicalDeposits(userId *int) (*model.PeriodicalDeposits, error)
	UpdatePeriodicalDeposits(userId *int, periodicalDeposits *model.PeriodicalDeposits) error

	GetRiskAssessment(userId *int) (*model.RiskAssessments, error)
	UpdateRiskAssessment(userId *int, riskAssessments *model.RiskAssessments) error

	AddExperience(userId *int, experience *model.Experience) error
	UpdateExperience(id *int, experience *model.Experience) error
	DeleteExperience(id *int) error
	GetExperiences(userId *int) (*model.Experiences, error)

	AddDeposit(userId *int, deposit *model.Deposit) error
	UpdateDeposit(id *int, deposit *model.Deposit) error
	DeleteDeposit(id *int) error
	GetDeposits(userId *int) (*model.Deposits, error)

	AddSecurity(userId *int, security *model.Security) error
	UpdateSecurity(id *int, security *model.Security) error
	DeleteSecurity(id *int) error
	GetSecurities(userId *int) (*model.Securities, error)
}
