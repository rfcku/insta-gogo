package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ApiRequestHandler represents a client for making API requests.
type ApiRequestHandler struct {
	Client *http.Client
}

// NewApiRequestHandler creates and returns a new ApiRequestHandler.
func NewApiRequestHandler() *ApiRequestHandler {
	return &ApiRequestHandler{
		Client: &http.Client{},
	}
}

// handleError checks for HTTP errors and logs them. It returns an error if the status code is not 200.
func (h *ApiRequestHandler) handleError(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		err := fmt.Errorf("error: %d - %s", resp.StatusCode, string(body))
		log.Println(err)
		return err
	}
	return nil
}

// Get sends a GET request to the specified URL with parameters and returns the response body as JSON.
func (h *ApiRequestHandler) Get(url string, params map[string]string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := h.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if err := h.handleError(resp); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("error decoding JSON:", err)
		return nil, err
	}

	log.Printf("Response: %d\n", resp.StatusCode)
	return result, nil
}

// Post sends a POST request to the specified URL with parameters and returns the response body as JSON.
func (h *ApiRequestHandler) Post(url string, params map[string]string) (map[string]interface{}, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := h.Client.Post(url, "application/json", bytes.NewBuffer(jsonParams))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if err := h.handleError(resp); err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("error decoding JSON:", err)
		return nil, err
	}

	log.Printf("Response: %d\n", resp.StatusCode)
	return result, nil
}
