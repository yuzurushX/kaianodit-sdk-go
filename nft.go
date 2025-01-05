package kaianodit

import (
	"encoding/json"
	"fmt"
)

// GetNFTContractMetadata retrieves metadata for multiple NFT contracts
// contractAddresses is required and must not be empty
func (c *Client) GetNFTContractMetadata(contractAddresses []string) ([]NFTContractMetadata, error) {
	// Validate required parameter
	if contractAddresses == nil {
		return nil, fmt.Errorf("contractAddresses is required")
	}
	if len(contractAddresses) == 0 {
		return nil, fmt.Errorf("contractAddresses cannot be empty")
	}

	// Validate individual addresses are not empty
	for i, addr := range contractAddresses {
		if addr == "" {
			return nil, fmt.Errorf("contract address at index %d cannot be empty", i)
		}
	}

	req := NFTContractMetadataRequest{
		ContractAddresses: contractAddresses,
	}

	resp, err := c.doRequest("POST", "/v1/kaia/mainnet/nft/getNftContractMetadataByContracts", req)
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

	var metadata []NFTContractMetadata
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return metadata, nil
}

// NFTContractsByAccountOptions represents optional parameters for GetNFTContractsByAccount
type NFTContractsByAccountOptions struct {
	ContractAddresses []string
	Page             *int
	Rpp              *int
	Cursor           *string
	WithCount        *bool
}

// GetNFTContractsByAccount retrieves NFT contracts for a specific account
// accountAddress is required
func (c *Client) GetNFTContractsByAccount(accountAddress string, opts *NFTContractsByAccountOptions) (*NFTContractsByAccountResponse, error) {
	// Validate required parameter
	if accountAddress == "" {
		return nil, fmt.Errorf("accountAddress is required")
	}

	// Prepare request
	req := NFTContractsByAccountRequest{
		AccountAddress: accountAddress,
	}

	// Apply optional parameters if provided
	if opts != nil {
		req.ContractAddresses = opts.ContractAddresses
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
	}

	resp, err := c.doRequest("POST", "/v1/kaia/mainnet/nft/getNftContractsByAccount", req)
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

	var result NFTContractsByAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
} 