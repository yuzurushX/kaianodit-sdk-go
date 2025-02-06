package kaianodit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) doRequestAndDecode(method, path string, body interface{}, result interface{}) error {
	var url string
	
	// Check if this is a node request (full URL) or regular API request (path only)
	if strings.HasPrefix(path, "http") {
		url = path
	} else {
		// Ensure path starts with /v1/kaia/ and includes the network for regular API requests
		if !strings.HasPrefix(path, "/v1/kaia/") {
			path = fmt.Sprintf("/v1/kaia/%s/%s", c.network, strings.TrimPrefix(path, "/"))
		}
		url = fmt.Sprintf("%s%s", c.baseURL, path)
	}

	var bodyReader io.Reader
	switch v := body.(type) {
	case []byte:
		bodyReader = bytes.NewReader(v)
	case nil:
		bodyReader = nil
	default:
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var apiError Error
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		return &apiError
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
} 