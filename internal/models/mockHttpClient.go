package models

import "net/http"

type MockResponseWriter struct {
	StatusCode int
	Body       []byte
	HeaderMap  map[string][]string
}

func (w *MockResponseWriter) Header() http.Header {
	if w.HeaderMap == nil {
		w.HeaderMap = make(map[string][]string)
	}
	return w.HeaderMap
}

func (w *MockResponseWriter) Write(body []byte) (int, error) {
	w.Body = body
	return len(body), nil
}

func (w *MockResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
}

type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (c *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.Response, c.Err
}
