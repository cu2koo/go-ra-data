package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoDataProviderDefined(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "no data provider is defined"})
}

func NoDataFound(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "request contains not known symbols"})
}

func QueryBindingFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "query binding failed"})
}

func NoSymbolDefined(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "no symbol is defined"})
}

func PeriodNumberConversionFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "period number conversion failed"})
}
