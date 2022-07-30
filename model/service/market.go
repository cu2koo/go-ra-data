package service

import "github.com/gin-gonic/gin"

type Market interface {
	GetPrices(c *gin.Context)
}
