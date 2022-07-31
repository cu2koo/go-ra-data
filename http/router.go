package http

import (
	"github.com/cu2koo/go-ra-data/model/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// serviceNotImplemented is a predefined error message to indicate that the service has not been implemented.
func serviceNotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "service not implemented"})
}

// SetupRouter routes http requests to a service endpoint.
func SetupRouter(accounts *gin.Accounts, c service.User, ms service.Market) *gin.Engine {
	r := gin.Default()

	// Group for the authorized users is created.
	authorized := r.Group("", gin.BasicAuth(*accounts))

	// Routing of the user service.
	if c != nil {
		authorized.POST("/user", c.CreateUser)
		authorized.GET("/user/:user", c.GetUsers)
		authorized.GET("/user/:userId", c.GetUser)
		authorized.PUT("/user/:userId", c.UpdateUser)
		authorized.DELETE("/user/:userId", c.DeleteUser)
		authorized.GET("/user/:userId/fi", c.GetFinancialInformation)
		authorized.PUT("/user/:userId/fi", c.UpdateFinancialInformation)
		authorized.GET("/user/:userId/goal", c.GetGoal)
		authorized.PUT("/user/:userId/goal", c.UpdateGoal)
		authorized.GET("/user/:userId/pd", c.GetPeriodicalDeposits)
		authorized.PUT("/user/:userId/pd", c.UpdatePeriodicalDeposits)
		authorized.GET("/user/:userId/ra", c.GetRiskAssessment)
		authorized.PUT("/user/:userId/ra", c.UpdateRiskAssessment)
		authorized.POST("/user/:userId/experience", c.AddExperience)
		authorized.GET("/user/:userId/experience", c.GetExperiences)
		authorized.PUT("/user/:userId/experience/:id", c.UpdateExperience)
		authorized.DELETE("/user/:userId/experience/:id", c.DeleteExperience)
		authorized.POST("/user/:userId/deposit", c.AddDeposit)
		authorized.GET("/user/:userId/deposit", c.GetDeposits)
		authorized.PUT("/user/:userId/deposit/:id", c.UpdateDeposit)
		authorized.DELETE("/user/:userId/deposit/:id", c.DeleteDeposit)
		authorized.POST("/user/:userId/security", c.AddSecurity)
		authorized.GET("/user/:userId/security", c.GetSecurities)
		authorized.PUT("/user/:userId/security/:id", c.UpdateSecurity)
		authorized.DELETE("/user/:userId/security/:id", c.DeleteSecurity)
	} else {
		r.POST("/user", serviceNotImplemented)
		r.GET("/user", serviceNotImplemented)
		r.GET("/user/:userId", serviceNotImplemented)
		r.PUT("/user/:userId", serviceNotImplemented)
		r.DELETE("/user/:userId", serviceNotImplemented)
		r.GET("/user/:userId/fi", serviceNotImplemented)
		r.PUT("/user/:userId/fi", serviceNotImplemented)
		r.GET("/user/:userId/goal", serviceNotImplemented)
		r.PUT("/user/:userId/goal", serviceNotImplemented)
		r.GET("/user/:userId/pd", serviceNotImplemented)
		r.PUT("/user/:userId/pd", serviceNotImplemented)
		r.GET("/user/:userId/ra", serviceNotImplemented)
		r.PUT("/user/:userId/ra", serviceNotImplemented)
		r.POST("/user/:userId/experience", serviceNotImplemented)
		r.GET("/user/:userId/experience", serviceNotImplemented)
		r.PUT("/user/:userId/experience/:id", serviceNotImplemented)
		r.DELETE("/user/:userId/experience/:id", serviceNotImplemented)
		r.POST("/user/:userId/deposit", serviceNotImplemented)
		r.GET("/user/:userId/deposit", serviceNotImplemented)
		r.PUT("/user/:userId/deposit/:id", serviceNotImplemented)
		r.DELETE("/user/:userId/deposit/:id", serviceNotImplemented)
		r.POST("/user/:userId/security", serviceNotImplemented)
		r.GET("/user/:userId/security", serviceNotImplemented)
		r.PUT("/user/:userId/security/:id", serviceNotImplemented)
		r.DELETE("/user/:userId/security/:id", serviceNotImplemented)
	}

	// Routing of the market service.
	if ms != nil {
		authorized.GET("/market/prices", ms.GetPrices)
	} else {
		r.GET("/market/prices", serviceNotImplemented)
	}

	return r
}
