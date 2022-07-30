package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActionFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "action failed"})
}

func ActionSuccessful(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "action successful"})
}

func FormBindingFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "form binding failed"})
}

func IdConversionFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "id conversion failed"})
}
