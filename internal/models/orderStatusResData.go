package models

type OrderStatusResData struct {
	Type                   string    `json:"type"`
	Status                 string    `json:"status"`
	ErrorMessage           string    `json:"errorMessage"`
	EndpointID             string    `json:"endpointID"`
	ProcessorTransactionID string    `json:"processorTransactionID"`
	OrderID                string    `json:"orderID"`
	MerchantOrderID        string    `json:"merchantOrderID"`
	Amount                 string    `json:"amount"`
	Currency               string    `json:"currency"`
	CustomerEmail          string    `json:"customerEmail"`
	CustomParam            string    `json:"customParam"`
	ExtraData              ExtraData `json:"extraData"`
	Request                Request   `json:"request"`
}

// ExtraData represents the extra data structure within the data
type ExtraData struct {
	AmountChanged     bool   `json:"amountChanged"`
	AmountRounded     bool   `json:"amountRounded"`
	AmountManipulated bool   `json:"amountManipulated"`
	DCC               bool   `json:"dcc"`
	OriginalAmount    string `json:"originalAmount"`
	PaymentMethod     string `json:"paymentMethod"`
	SelectedBankCode  string `json:"selectedBankCode"`
	SelectedBankName  string `json:"selectedBankName"`
}

// Request represents the request structure within the data
type Request struct {
	MerchantID      string `json:"merchantID"`
	OrderID         string `json:"orderID"`
	MerchantOrderID string `json:"merchantOrderID"`
	Timestamp       string `json:"timestamp"`
}
