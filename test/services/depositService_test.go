package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/bozhidarv/zota-dev-challenge/internal/models"
	"github.com/bozhidarv/zota-dev-challenge/internal/services"
)

func TestMakeDepostiRequest(t *testing.T) {
	mockResponseData := models.BasicResponse[models.DepositResponseData]{
		Code:    "200",
		Message: "Success",
		Data: models.DepositResponseData{
			DepositURL:      "https://example.com/deposit",
			MerchantOrderID: "QvE8dZshpKhaOmHY",
			OrderID:         "53953",
		},
	}
	mockResponseBody, _ := json.Marshal(mockResponseData)

	// Create a mock HTTP response
	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(mockResponseBody)),
	}

	// Create a mock HTTP client
	mockClient := &models.MockHTTPClient{
		Response: mockResp,
		Err:      nil,
	}

	// Call the function with the mock client
	ip := "103.106.8.104"
	orderId := "QvE8dZshpKhaOmHY"
	checkoutURL := "https://example.com/checkout"

	reqBody := models.CustomerDepositData{
		OrderAmount: "100",
		Email:       "customer@cust.com",
		FirstName:   "John",
		LastName:    "Doe",
		Address:     "123 Main St",
		CountryCode: "US",
		City:        "New York",
		ZipCode:     "10001",
		Phone:       "12819374192",
	}

	response, err := services.MakeDepostiRequest(mockClient, reqBody, ip, orderId, checkoutURL)
	// Verify the results
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != "200" {
		t.Errorf("Expected status code 200, got %v", response.Code)
	}

	if response.Data.MerchantOrderID != mockResponseData.Data.MerchantOrderID {
		t.Errorf(
			"Expected order ID %v, got %v",
			mockResponseData.Data.OrderID,
			response.Data.OrderID,
		)
	}
}
