package service

import "github.com/gin-gonic/gin"

// Market is a service interface for the market service.
type Market interface {
	// GetPrices handles the service for the retrieval of price data.
	GetPrices(c *gin.Context)
}
