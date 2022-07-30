package factset

type TimeSeriesFloat64 struct {
	Dates  []*string  `json:"dates"`
	Values []*float64 `json:"values"`
}

type TimeSeriesString struct {
	Dates  []*string `json:"dates"`
	Values []*string `json:"values"`
}

type DataItem struct {
	Error      *int    `json:"error"`
	Formula    *string `json:"formula"`
	RequestId  *string `json:"requestId"`
	DataType   *string `json:"dataType"`
	ObjectType *string `json:"objectType"`
}

type DataItemTimeSeriesFloat64 struct {
	DataItem
	Result *TimeSeriesFloat64 `json:"result"`
}

type DataItemTimeSeriesString struct {
	DataItem
	Result *TimeSeriesString `json:"result"`
}

type DataTimeSeriesFloat64 struct {
	Data []*DataItemTimeSeriesFloat64 `json:"data"`
	Meta *interface{}                 `json:"meta"`
}

type DataTimeSeriesString struct {
	Data []*DataItemTimeSeriesString `json:"data"`
	Meta *interface{}                `json:"meta"`
}

type DataRequest struct {
	Ids      *[]string `json:"ids"`
	Formulas *[]string `json:"formulas"`
}

type Request struct {
	Data *DataRequest `json:"data"`
}
