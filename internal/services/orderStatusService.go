package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bozhidarv/zota-dev-challenge/internal/models"
)

func MakeOrderStatusRequest(merchantOrderID string, orderId string) {
	time.Sleep(10 * time.Second)

	currTimestamp := strconv.FormatInt(time.Now().Unix(), 10)

	signature := encryptSHA256(
		models.MerchantID + merchantOrderID + orderId + currTimestamp + models.APIKey,
	)

	url := fmt.Sprintf(
		"https://%s/api/v1/order/status/%s/?merchantID=%s&merchantOrderID=%s&orderID=%s&timestamp=%s&signature=%s",
		models.BaseURL,
		models.EndpointID,
		models.MerchantID,
		merchantOrderID,
		orderId,
		currTimestamp,
		signature,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	var responseBody models.BasicResponse[models.OrderStatusResData]
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		fmt.Println("message: ", responseBody.Message)
		return
	}

	if responseBody.Data.Status == "PROCESSING" {
		MakeOrderStatusRequest(merchantOrderID, orderId)
	}
}
