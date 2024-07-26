package models

import "os"

var (
	EndpointID = os.Getenv("ZOTAPAY_ENDPOINT_ID")
	APIKey     = os.Getenv("ZOTAPAY_API_KEY")
	CURR       = os.Getenv("ZOTAPAY_CURR")
	MerchantID = os.Getenv("ZOTAPAY_MERCHANT_ID")
	BaseURL    = os.Getenv("ZOTAPAY_BASE_URL")
)
