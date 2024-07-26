package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bozhidarv/zota-dev-challenge/internal/models"
)

func encryptSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func MakeDepostiRequest(
	client models.HTTPClient,
	reqBody models.CustomerDepositData,
	ip string,
	orderId string,
	checkoutURL string,
) (models.BasicResponse[models.DepositResponseData], error) {
	depositReq := models.DepositReq{
		MerchantOrderID:     orderId,
		MerchantOrderDesc:   "Deposit for order " + orderId,
		OrderAmount:         reqBody.OrderAmount,
		OrderCurrency:       "USD",
		CustomerEmail:       reqBody.Email,
		CustomerFirstName:   reqBody.FirstName,
		CustomerLastName:    reqBody.LastName,
		CustomerAddress:     reqBody.Address,
		CustomerCountryCode: reqBody.CountryCode,
		CustomerCity:        reqBody.Email,
		CustomerZipCode:     reqBody.ZipCode,
		CustomerPhone:       reqBody.Phone,
		CustomerIP:          ip,
		RedirectURL:         "https://duckduckgo.com/?q=yes&t=brave&ia=web",
		CheckoutURL:         checkoutURL,
		Signature: encryptSHA256(
			models.EndpointID + orderId + reqBody.OrderAmount + reqBody.Email + models.APIKey,
		),
	}

	orderJSON, err := json.Marshal(depositReq)
	if err != nil {
		fmt.Println("Error marshaling order:", err)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}

	url := fmt.Sprintf("https://%s/api/v1/deposit/request/%s/", models.BaseURL, models.EndpointID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(orderJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}
	defer resp.Body.Close()

	var responseBody models.BasicResponse[models.DepositResponseData]
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		fmt.Println("Error decoding response:", err)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		fmt.Println("message: ", responseBody.Message)
		return models.BasicResponse[models.DepositResponseData]{}, err
	}

	return responseBody, nil
}
