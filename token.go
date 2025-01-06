package kaianodit

import (
	"fmt"
)

// GetNativeBalance retrieves the native token balance for a specific account
// accountAddress is required
func (c *Client) GetNativeBalance(accountAddress string) (*NativeBalanceResponse, error) {
	if accountAddress == "" {
		return nil, fmt.Errorf("accountAddress is required")
	}

	req := struct {
		AccountAddress string `json:"accountAddress"`
	}{
		AccountAddress: accountAddress,
	}

	var result NativeBalanceResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/native/getNativeBalanceByAccount", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTokenAllowance(contractAddress, ownerAddress, spenderAddress string) (*TokenAllowanceResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contractAddress is required")
	}
	if ownerAddress == "" {
		return nil, fmt.Errorf("ownerAddress is required")
	}
	if spenderAddress == "" {
		return nil, fmt.Errorf("spenderAddress is required")
	}

	req := TokenAllowanceRequest{
		ContractAddress: contractAddress,
		OwnerAddress:   ownerAddress,
		SpenderAddress: spenderAddress,
	}

	var result TokenAllowanceResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/getTokenAllowance", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTokenContractMetadata retrieves metadata for multiple token contracts
func (c *Client) GetTokenContractMetadata(contractAddresses []string) ([]TokenContractMetadata, error) {
	if len(contractAddresses) == 0 {
		return nil, fmt.Errorf("contractAddresses cannot be empty")
	}

	req := TokenContractMetadataRequest{
		ContractAddresses: contractAddresses,
	}

	var result []TokenContractMetadata
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/getTokenContractMetadataByContracts", req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTokenHolders retrieves token holders for a specific contract
func (c *Client) GetTokenHolders(contractAddress string, opts *TokenHoldersOptions) (*TokenHoldersResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contractAddress is required")
	}

	req := TokenHoldersRequest{
		ContractAddress: contractAddress,
	}

	if opts != nil {
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
	}

	var result TokenHoldersResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/getTokenHoldersByContract", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// TokenHoldersOptions represents optional parameters for GetTokenHolders
type TokenHoldersOptions struct {
	Page      *int
	Rpp       *int
	Cursor    *string
	WithCount *bool
}

// GetTokenPrices retrieves detailed price information for multiple token contracts
func (c *Client) GetTokenPrices(contractAddresses []string) ([]TokenPrice, error) {
	if len(contractAddresses) == 0 {
		return nil, fmt.Errorf("contractAddresses cannot be empty")
	}

	req := TokenPricesRequest{
		ContractAddresses: contractAddresses,
	}

	var result []TokenPrice
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/getTokenPricesByContracts", req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// TokenTransfersOptions represents optional parameters for GetTokenTransfers
type TokenTransfersOptions struct {
	Relation          string
	ContractAddresses []string
	FromBlock         string
	ToBlock           string
	FromDate          string
	ToDate            string
	Page              *int
	Rpp               *int
	Cursor            *string
	WithCount         *bool
	WithMetadata      *bool
	WithZeroValue     *bool
}

// GetTokenTransfers retrieves token transfers for a specific account
func (c *Client) GetTokenTransfers(accountAddress string, opts *TokenTransfersOptions) (*TokenTransfersResponse, error) {
	if accountAddress == "" {
		return nil, fmt.Errorf("accountAddress is required")
	}

	req := TokenTransfersRequest{
		AccountAddress: accountAddress,
	}

	if opts != nil {
		req.Relation = opts.Relation
		req.ContractAddresses = opts.ContractAddresses
		req.FromBlock = opts.FromBlock
		req.ToBlock = opts.ToBlock
		req.FromDate = opts.FromDate
		req.ToDate = opts.ToDate
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
		req.WithMetadata = opts.WithMetadata
		req.WithZeroValue = opts.WithZeroValue
	}

	var result TokenTransfersResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/getTokenTransfersByAccount", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// TokenTransfersByContractOptions represents optional parameters for GetTokenTransfersByContract
type TokenTransfersByContractOptions struct {
	FromBlock     string
	ToBlock       string
	FromDate      string
	ToDate        string
	Page          *int
	Rpp           *int
	Cursor        *string
	WithCount     *bool
	WithZeroValue *bool
}

// GetTokenTransfersByContract retrieves token transfers for a specific contract
func (c *Client) GetTokenTransfersByContract(contractAddress string, opts *TokenTransfersByContractOptions) (*TokenTransfersByContractResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contractAddress is required")
	}

	req := TokenTransfersByContractRequest{
		ContractAddress: contractAddress,
	}

	if opts != nil {
		req.FromBlock = opts.FromBlock
		req.ToBlock = opts.ToBlock
		req.FromDate = opts.FromDate
		req.ToDate = opts.ToDate
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
		req.WithZeroValue = opts.WithZeroValue
	}

	var result TokenTransfersByContractResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/getTokenTransfersByContract", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// TokenTransfersWithinRangeOptions represents optional parameters for GetTokenTransfersWithinRange
type TokenTransfersWithinRangeOptions struct {
	FromBlock     string
	ToBlock       string
	FromDate      string
	ToDate        string
	Page          *int
	Rpp           *int
	Cursor        *string
	WithCount     *bool
	WithMetadata  *bool
	WithZeroValue *bool
}

// GetTokenTransfersWithinRange retrieves token transfers within a specified range
func (c *Client) GetTokenTransfersWithinRange(opts *TokenTransfersWithinRangeOptions) (*TokenTransfersWithinRangeResponse, error) {
	req := TokenTransfersWithinRangeRequest{}

	if opts != nil {
		req.FromBlock = opts.FromBlock
		req.ToBlock = opts.ToBlock
		req.FromDate = opts.FromDate
		req.ToDate = opts.ToDate
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
		req.WithMetadata = opts.WithMetadata
		req.WithZeroValue = opts.WithZeroValue
	}

	var result TokenTransfersWithinRangeResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/getTokenTransfersWithinRange", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// TokensOwnedOptions represents optional parameters for GetTokensOwned
type TokensOwnedOptions struct {
	ContractAddress string
	Page            *int
	Rpp             *int
	Cursor          *string
	WithCount       *bool
}

// GetTokensOwned retrieves tokens owned by an account
func (c *Client) GetTokensOwned(accountAddress string, opts *TokensOwnedOptions) (*TokensOwnedResponse, error) {
	if accountAddress == "" {
		return nil, fmt.Errorf("accountAddress is required")
	}

	req := TokensOwnedRequest{
		AccountAddress: accountAddress,
	}

	if opts != nil {
		req.ContractAddress = opts.ContractAddress
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
	}

	var result TokensOwnedResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/getTokensOwnedByAccount", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// TokenContractMetadataSearchOptions represents optional parameters for SearchTokenContractMetadata
type TokenContractMetadataSearchOptions struct {
	Page      *int
	Rpp       *int
	Cursor    *string
	WithCount *bool
}

// SearchTokenContractMetadata searches for token contract metadata using a keyword
func (c *Client) SearchTokenContractMetadata(keyword string, opts *TokenContractMetadataSearchOptions) (*TokenContractMetadataSearchResponse, error) {
	if keyword == "" {
		return nil, fmt.Errorf("keyword is required")
	}

	req := TokenContractMetadataSearchRequest{
		Keyword: keyword,
	}

	if opts != nil {
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
	}

	var result TokenContractMetadataSearchResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/token/searchTokenContractMetadataByKeyword", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
