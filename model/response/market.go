package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NoDataProviderDefined is a predefined error message to indicate that no data provider has been defined.
func NoDataProviderDefined(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "no data provider is defined"})
}

// NoDataFound is a predefined error message to indicate that no data has been found.
func NoDataFound(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "request contains not known symbols"})
}

// QueryBindingFailed is a predefined error message to indicate that the query binding has been failed.
func QueryBindingFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "query binding failed"})
}

// NoSymbolDefined is a predefined error message to indicate that no symbol has been defined.
func NoSymbolDefined(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "no symbol is defined"})
}

// PeriodNumberConversionFailed is a predefined error message to indicate that the period number conversion has been failed.
func PeriodNumberConversionFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "period number conversion failed"})
}
