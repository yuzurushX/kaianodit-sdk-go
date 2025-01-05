package kaianodit

import (
	"encoding/json"
	"fmt"
)

// GetNativeBalance retrieves the native token balance for a specific account
// accountAddress is required
func (c *Client) GetNativeBalance(accountAddress string) (*NativeBalanceResponse, error) {
	// Validate required parameter
	if accountAddress == "" {
		return nil, fmt.Errorf("accountAddress is required")
	}

	// Prepare request
	req := struct {
		AccountAddress string `json:"accountAddress"`
	}{
		AccountAddress: accountAddress,
	}

	resp, err := c.doRequest("POST", "/v1/kaia/mainnet/native/getNativeBalanceByAccount", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var apiError Error
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		return nil, &apiError
	}

	var result NativeBalanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
} 