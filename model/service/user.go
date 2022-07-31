package service

import "github.com/gin-gonic/gin"

// User is a service interface for the user service.
type User interface {
	// CreateUser handles the service for creating a user.
	CreateUser(c *gin.Context)
	// GetUser handles the service for getting a user.
	GetUser(c *gin.Context)
	// GetUsers handles the service for getting all users.
	GetUsers(c *gin.Context)
	// UpdateUser handles the service for updating a user.
	UpdateUser(c *gin.Context)
	// DeleteUser handles the service for deleting a user.
	DeleteUser(c *gin.Context)
	// GetFinancialInformation handles the service for getting financial information about a user.
	GetFinancialInformation(c *gin.Context)
	// UpdateFinancialInformation handles the service for updating financial information about a user.
	UpdateFinancialInformation(c *gin.Context)
	// GetGoal handles the service for getting goal information about a user.
	GetGoal(c *gin.Context)
	// UpdateGoal handles the service for updating goal information about a user.
	UpdateGoal(c *gin.Context)
	// GetPeriodicalDeposits handles the service for getting periodical deposits information about a user.
	GetPeriodicalDeposits(c *gin.Context)
	// UpdatePeriodicalDeposits handles the service for updating periodical deposits information about a user.
	UpdatePeriodicalDeposits(c *gin.Context)
	// GetRiskAssessment handles the service for getting risk assessment information about a user.
	GetRiskAssessment(c *gin.Context)
	// UpdateRiskAssessment handles the service for updating risk assessment information about a user.
	UpdateRiskAssessment(c *gin.Context)
	// AddExperience handles the service for creating an experience entry for a user.
	AddExperience(c *gin.Context)
	// UpdateExperience handles the service for updating an experience entry for a user.
	UpdateExperience(c *gin.Context)
	// DeleteExperience handles the service for deleting an experience entry for a user.
	DeleteExperience(c *gin.Context)
	// GetExperiences handles the service for getting all experience entries for a user.
	GetExperiences(c *gin.Context)
	// AddDeposit handles the service for creating a deposit entry for a user.
	AddDeposit(c *gin.Context)
	// UpdateDeposit handles the service for updating a deposit entry for a user.
	UpdateDeposit(c *gin.Context)
	// DeleteDeposit handles the service for deleting a deposit entry for a user.
	DeleteDeposit(c *gin.Context)
	// GetDeposits handles the service for getting all deposit entries for a user.
	GetDeposits(c *gin.Context)
	// AddSecurity handles the service for creating a security entry for a user.
	AddSecurity(c *gin.Context)
	// UpdateSecurity handles the service for updating a security entry for a user.
	UpdateSecurity(c *gin.Context)
	// DeleteSecurity handles the service for deleting a security entry for a user.
	DeleteSecurity(c *gin.Context)
	// GetSecurities handles the service for getting all security entries for a user.
	GetSecurities(c *gin.Context)
}
