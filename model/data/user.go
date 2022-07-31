package data

import (
	"github.com/cu2koo/go-ra-data/model"
)

// User is a data connection interface for the user database.
type User interface {
	// CreateUser handles user creation.
	CreateUser(user *model.User) error
	// GetUser handles user retrieval.
	GetUser(userId *int) (*model.User, error)
	// GetUsers handles retrieval of users.
	GetUsers() (*model.Users, error)
	// UpdateUser handles user update.
	UpdateUser(userId *int, user *model.User) error
	// DeleteUser handles user deletion.
	DeleteUser(userId *int) error
	// GetFinancialInformation handles financial information retrieval.
	GetFinancialInformation(userId *int) (*model.FinancialInformation, error)
	// UpdateFinancialInformation handles financial information update.
	UpdateFinancialInformation(userId *int, financialInformation *model.FinancialInformation) error
	// GetGoal handles goal information retrieval.
	GetGoal(userId *int) (*model.Goal, error)
	// UpdateGoal handles goal information update.
	UpdateGoal(userId *int, goal *model.Goal) error
	// GetPeriodicalDeposits handles periodical deposits information retrieval.
	GetPeriodicalDeposits(userId *int) (*model.PeriodicalDeposits, error)
	// UpdatePeriodicalDeposits handles periodical deposits update.
	UpdatePeriodicalDeposits(userId *int, periodicalDeposits *model.PeriodicalDeposits) error
	// GetRiskAssessment handles risk assessment retrieval.
	GetRiskAssessment(userId *int) (*model.RiskAssessments, error)
	// UpdateRiskAssessment handles risk assessment update.
	UpdateRiskAssessment(userId *int, riskAssessments *model.RiskAssessments) error
	// AddExperience handles experience entry creation.
	AddExperience(userId *int, experience *model.Experience) error
	// UpdateExperience handles experience entry update.
	UpdateExperience(id *int, experience *model.Experience) error
	// DeleteExperience handles experience entry deletion.
	DeleteExperience(id *int) error
	// GetExperiences handles retrieval of experience entries.
	GetExperiences(userId *int) (*model.Experiences, error)
	// AddDeposit handles deposit entry creation.
	AddDeposit(userId *int, deposit *model.Deposit) error
	// UpdateDeposit handles deposit entry update.
	UpdateDeposit(id *int, deposit *model.Deposit) error
	// DeleteDeposit handles  deposit entry deletion.
	DeleteDeposit(id *int) error
	// GetDeposits handles retrieval of deposit entries.
	GetDeposits(userId *int) (*model.Deposits, error)
	// AddSecurity handles security entry creation.
	AddSecurity(userId *int, security *model.Security) error
	// UpdateSecurity handles security entry update.
	UpdateSecurity(id *int, security *model.Security) error
	// DeleteSecurity handles security entry deletion.
	DeleteSecurity(id *int) error
	// GetSecurities handles retrieval of security entries.
	GetSecurities(userId *int) (*model.Securities, error)
}
