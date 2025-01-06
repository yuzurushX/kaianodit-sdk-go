package kaianodit

import "fmt"

// GetAccountStats retrieves statistics for a specific account
func (c *Client) GetAccountStats(address string) (*AccountStats, error) {
    if address == "" {
        return nil, fmt.Errorf("address is required")
    }

    req := AccountStatsRequest{
        Address: address,
    }

    var result AccountStats
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/stats/getAccountStats", req, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
} 