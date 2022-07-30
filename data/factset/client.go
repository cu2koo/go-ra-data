package factset

import (
	"bytes"
	"errors"
	"github.com/cu2koo/go-ra-data/model"
	"github.com/goccy/go-json"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var source = "factset.com"

type FactSet struct {
	client   *http.Client
	baseURL  string
	username string
	apiKey   string
}

func New() (*FactSet, error) {
	vp := viper.New()
	vp.AddConfigPath("./data/factset")
	vp.SetConfigName("config")
	vp.SetConfigType("json")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &FactSet{
		baseURL:  vp.GetString("baseURL"),
		username: vp.GetString("username"),
		apiKey:   vp.GetString("apiKey"),
		client: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}

func (s *FactSet) sendRequest(ids *[]string, formulas *[]string) (*http.Response, error) {
	data := &Request{
		&DataRequest{
			ids, formulas,
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, s.baseURL+"/time-series", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(s.username, s.apiKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	return s.client.Do(request)
}

func (s *FactSet) GetPrices(symbols *[]string, period *time.Duration) ([]*model.PriceData, error) {
	startDate := time.Now().Add(-*period)
	timeToken := startDate.Format("01/02/2006")

	timesSeriesFormulasString := &[]string{
		"P_SYMBOL",
		"P_CURRENCY(\"ISO\")",
	}
	timeSeriesFormulasFloat64 := &[]string{
		"P_PRICE_OPEN(" + timeToken + ",NOW)",
		"FG_PRICE_HIGH(" + timeToken + ",NOW)",
		"FG_PRICE_LOW(" + timeToken + ",NOW)",
		"FG_PRICE(" + timeToken + ",NOW)",
		"XP_VOLUME(" + timeToken + ",NOW)",
		"FF_PRICE_HIGH_52WK",
		"FF_PRICE_LOW_52WK",
	}

	timesSeriesStringResponse, err := s.sendRequest(symbols, timesSeriesFormulasString)
	if err != nil {
		return nil, err
	}
	timeSeriesFloat64Response, err := s.sendRequest(symbols, timeSeriesFormulasFloat64)
	if err != nil {
		return nil, err
	}

	if timesSeriesStringResponse.StatusCode != http.StatusOK || timeSeriesFloat64Response.StatusCode != http.StatusOK {
		return nil, errors.New("request failed")
	}

	dataTimesSeriesString := DataTimeSeriesString{}
	dataTimeSeriesFloat64 := DataTimeSeriesFloat64{}
	err = json.NewDecoder(timesSeriesStringResponse.Body).Decode(&dataTimesSeriesString)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(timeSeriesFloat64Response.Body).Decode(&dataTimeSeriesFloat64)
	if err != nil {
		return nil, err
	}

	return createPriceData(&dataTimesSeriesString, &dataTimeSeriesFloat64, &timeToken)
}

func createPriceData(dataTimeSeriesString *DataTimeSeriesString, dataTimeSeriesFloat64 *DataTimeSeriesFloat64, timeToken *string) ([]*model.PriceData, error) {
	sortedDataTimeSeriesString := make(map[string][]*DataItemTimeSeriesString)
	for _, value := range dataTimeSeriesString.Data {
		sortedDataTimeSeriesString[*value.RequestId] = append(sortedDataTimeSeriesString[*value.RequestId], value)
	}
	sortedDataTimeSeries64 := make(map[string][]*DataItemTimeSeriesFloat64)
	for _, value := range dataTimeSeriesFloat64.Data {
		sortedDataTimeSeries64[*value.RequestId] = append(sortedDataTimeSeries64[*value.RequestId], value)
	}

	priceDataMap := make(map[string]*model.PriceData)

	for _, ticker := range sortedDataTimeSeriesString {
		var priceData model.PriceData
		var requestId string
		for _, value := range ticker {
			switch *value.Formula {
			case "P_SYMBOL":
				requestId = *value.RequestId
				priceData.Ticker = value.Result.Values[0]
				if priceData.Ticker == nil {
					return nil, errors.New("requested ticker not found")
				}
				continue
			case "P_CURRENCY(\"ISO\")":
				priceData.Currency = value.Result.Values[0]
				continue
			}
		}
		priceData.Source = &source
		priceDataMap[requestId] = &priceData
	}

	for _, ticker := range sortedDataTimeSeries64 {
		var priceData *model.PriceData
		prices := make(map[string]*model.Price)
		var priceOrder []string
		for _, value := range ticker {
			switch *value.Formula {
			case "P_PRICE_OPEN(" + *timeToken + ",NOW)":
				priceData = priceDataMap[*value.RequestId]
				for index, price := range value.Result.Values {
					date := *value.Result.Dates[index]
					if prices[date] == nil {
						prices[date] = &model.Price{
							Date: &date,
							Open: price,
						}
						priceOrder = append(priceOrder, date)
					} else {
						prices[date].Open = price
					}
				}
				continue
			case "FG_PRICE_HIGH(" + *timeToken + ",NOW)":
				for index, price := range value.Result.Values {
					date := *value.Result.Dates[index]
					prices[date].High = price
				}
				continue
			case "FG_PRICE_LOW(" + *timeToken + ",NOW)":
				for index, price := range value.Result.Values {
					date := *value.Result.Dates[index]
					prices[date].Low = price
				}
				continue
			case "FG_PRICE(" + *timeToken + ",NOW)":
				for index, price := range value.Result.Values {
					date := *value.Result.Dates[index]
					prices[date].Close = price
				}
				continue
			case "XP_VOLUME(" + *timeToken + ",NOW)":
				for index, volume := range value.Result.Values {
					date := *value.Result.Dates[index]
					prices[date].Volume = volume
				}
				continue
			case "FF_PRICE_HIGH_52WK":
				if value.Result.Values[0] != nil {
					priceData.PriceHigh52Wk = value.Result.Values[0]
				}
				continue
			case "FF_PRICE_LOW_52WK":
				if value.Result.Values[0] != nil {
					priceData.PriceLow52Wk = value.Result.Values[0]
				}
				continue
			}
		}
		for i := len(priceOrder) - 1; i >= 0; i-- {
			priceData.Prices = append(priceData.Prices, prices[priceOrder[i]])
		}
	}

	var priceData []*model.PriceData
	for _, value := range priceDataMap {
		priceData = append(priceData, value)
	}

	return priceData, nil
}
