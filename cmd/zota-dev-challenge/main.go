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

	"github.com/bozhidarv/zota-dev-challenge/internal/services"
)

func getIPAddress(r *http.Request) string {
	// Try to get the IP from the X-Forwarded-For header
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For can contain multiple IPs, get the first one
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	// Try to get the IP from the X-Real-IP header
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	// Get the IP from RemoteAddr
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

		responseBody, err := services.MakeDepostiRequest(getIPAddress(r), orderId, getFullURL(r))
		if err != nil {
			http.Error(w, "Error making deposit request", http.StatusInternalServerError)
		}

		respBodyStr, _ := json.Marshal(responseBody)

		_, err = w.Write(respBodyStr)
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
