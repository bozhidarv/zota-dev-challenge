package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/bozhidarv/zota-dev-challenge/internal/models"
	"github.com/bozhidarv/zota-dev-challenge/internal/services"
)

func getIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func getFullURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	return fmt.Sprintf("%s://%s%s", scheme, r.Host, r.URL.RequestURI())
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		orderId := uuid.New().String()

		defer r.Body.Close()

		var reqBody models.CustomerDepositData
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			fmt.Println("Error decoding req body:", err)
			http.Error(w, "Error making deposit request", http.StatusInternalServerError)
		}

		client := &http.Client{}
		resBody, err := services.MakeDepostiRequest(
			client,
			reqBody,
			getIPAddress(r),
			orderId,
			getFullURL(r),
		)
		if err != nil {
			http.Error(w, "Error making deposit request", http.StatusInternalServerError)
		}

		go services.MakeOrderStatusRequest(client, orderId, resBody.Data.OrderID)

		respBodyStr, _ := json.Marshal(resBody)

		_, err = w.Write(respBodyStr)
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
