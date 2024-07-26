package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bozhidarv/zota-dev-challenge/internal/models"
)

func encryptSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

var (
	EndpointID = os.Getenv("ZOTAPAY_ENDPOINT_ID")
	APIKey     = os.Getenv("ZOTAPAY_API_KEY")
	CURR       = os.Getenv("ZOTAPAY_CURR")
	MerchantID = os.Getenv("ZOTAPAY_MERCHANT_ID")
	BaseURL    = os.Getenv("ZOTAPAY_BASE_URL")
)

func MakeDepostiRequest(
	ip string,
	orderId string,
	checkoutURL string,
) (models.BasicResponse[models.DepositResponseData], error) {
	depositReq := models.DepositReq{
		MerchantOrderID:     orderId,
		MerchantOrderDesc:   "Test order",
		OrderAmount:         "500.00",
		OrderCurrency:       "USD",
		CustomerEmail:       "customer@email-address.com",
		CustomerFirstName:   "John",
		CustomerLastName:    "Doe",
		CustomerAddress:     "5/5 Moo 5 Thong Nai Pan Noi Beach, Baan Tai, Koh Phangan",
		CustomerCountryCode: "TH",
		CustomerCity:        "Surat Thani",
		CustomerZipCode:     "84280",
		CustomerPhone:       "+66-77999110",
		CustomerIP:          ip,
		RedirectURL:         "https://duckduckgo.com/?q=yes&t=brave&ia=web",
		CheckoutURL:         checkoutURL,
		Signature: encryptSHA256(
			EndpointID + orderId + "500.0customer@email-address.com" + APIKey,
		),
	}

	orderJSON, err := json.Marshal(depositReq)
	if err != nil {
		fmt.Println("Error marshaling order:", err)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}

	// Create a new HTTP POST request
	url := fmt.Sprintf("https://%s/api/v1/deposit/request/%s/", BaseURL, EndpointID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(orderJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the http.DefaultClient
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}
	defer resp.Body.Close()

	// Print the response body (if needed)
	var responseBody models.BasicResponse[models.DepositResponseData]
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		fmt.Println("Error decoding response:", err)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}

	return responseBody, nil
}
