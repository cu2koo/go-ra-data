package service

import (
	"github.com/cu2koo/go-ra-data/model"
	"github.com/cu2koo/go-ra-data/model/data"
	"github.com/cu2koo/go-ra-data/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const year = time.Hour * 24 * 365
const month = time.Hour * 24 * 30
const day = time.Hour * 24

type RequestData struct {
	Symbol *string `form:"symbol"`
	Period *string `form:"period"`
}

type MarketService struct {
	data *[]data.Market
}

func NewMarketService(data *[]data.Market) *MarketService {
	ms := &MarketService{
		data: data,
	}
	return ms
}

func (ms *MarketService) GetPrices(c *gin.Context) {
	if ms.data == nil {
		response.NoDataProviderDefined(c)
		return
	}

	var requestData RequestData
	err := c.ShouldBind(&requestData)
	if err != nil {
		response.QueryBindingFailed(c)
		return
	}

	if requestData.Symbol == nil {
		response.NoSymbolDefined(c)
		return
	}
	symbols := strings.Split(*requestData.Symbol, ",")

	period := year
	if requestData.Period != nil {
		periodString := strings.ToLower(*requestData.Period)
		if strings.HasSuffix(periodString, "y") {
			periodStringNumber := strings.TrimSuffix(periodString, "y")
			periodNumber, err := strconv.Atoi(periodStringNumber)
			if err != nil {
				response.PeriodNumberConversionFailed(c)
				return
			}
			period = year * time.Duration(periodNumber)
		}
		if strings.HasSuffix(periodString, "m") {
			periodStringNumber := strings.TrimSuffix(periodString, "m")
			periodNumber, err := strconv.Atoi(periodStringNumber)
			if err != nil {
				response.PeriodNumberConversionFailed(c)
				return
			}
			period = month * time.Duration(periodNumber)
		}
		if strings.HasSuffix(periodString, "d") {
			periodStringNumber := strings.TrimSuffix(periodString, "d")
			periodNumber, err := strconv.Atoi(periodStringNumber)
			if err != nil {
				response.PeriodNumberConversionFailed(c)
				return
			}
			period = day * time.Duration(periodNumber)
		}

	}

	var result []*model.PriceData
	for _, dataProvider := range *ms.data {
		res, err := dataProvider.GetPrices(&symbols, &period)
		if err != nil {
			continue
		}
		result = append(result, res...)
	}

	if len(result) <= 0 {
		response.NoDataFound(c)
		return
	}

	c.JSON(http.StatusOK, result)
}
