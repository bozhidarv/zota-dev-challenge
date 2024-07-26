package models

type CustomerDepositData struct {
	OrderAmount string `json:"orderAmount"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	Phone       string `json:"phone"`
}
