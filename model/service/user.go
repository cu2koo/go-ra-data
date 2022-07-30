package service

import "github.com/gin-gonic/gin"

type User interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	GetFinancialInformation(c *gin.Context)
	UpdateFinancialInformation(c *gin.Context)

	GetGoal(c *gin.Context)
	UpdateGoal(c *gin.Context)

	GetPeriodicalDeposits(c *gin.Context)
	UpdatePeriodicalDeposits(c *gin.Context)

	GetRiskAssessment(c *gin.Context)
	UpdateRiskAssessment(c *gin.Context)

	AddExperience(c *gin.Context)
	UpdateExperience(c *gin.Context)
	DeleteExperience(c *gin.Context)
	GetExperiences(c *gin.Context)

	AddDeposit(c *gin.Context)
	UpdateDeposit(c *gin.Context)
	DeleteDeposit(c *gin.Context)
	GetDeposits(c *gin.Context)

	AddSecurity(c *gin.Context)
	UpdateSecurity(c *gin.Context)
	DeleteSecurity(c *gin.Context)
	GetSecurities(c *gin.Context)
}
