package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ActionFailed is a predefined error message to indicate that an operation has been failed.
func ActionFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "action failed"})
}

// ActionSuccessful is a predefined success message to indicate that an operation has been successful.
func ActionSuccessful(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "action successful"})
}

// FormBindingFailed is a predefined error message to indicate that the form binding has been failed.
func FormBindingFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "form binding failed"})
}

// IdConversionFailed is a predefined error message to indicate that the id conversion has been failed.
func IdConversionFailed(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "id conversion failed"})
}
