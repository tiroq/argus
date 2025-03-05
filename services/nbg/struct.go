package nbg

type CurrencyInfo struct {
	Code          string  `json:"code"`
	Quantity      int     `json:"quantity"`
	RateFormatted string  `json:"rateFormatted"`
	DiffFormatted string  `json:"diffFormatted"`
	Rate          float64 `json:"rate"`
	Name          string  `json:"name"`
	Diff          float64 `json:"diff"`
	Date          string  `json:"date"`
	ValidFromDate string  `json:"validFromDate"`
}

type HistEntry struct {
	Date       string         `json:"date"`
	Currencies []CurrencyInfo `json:"currencies"`
}
