package kaianodit

import (
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

	var result []NFTContractMetadata
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftContractMetadataByContracts", req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
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

	var result NFTContractsByAccountResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftContractsByAccount", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
} 

// NFTHoldersOptions represents optional parameters for GetNFTHolders methods
type NFTHoldersOptions struct {
	Page      *int
	Rpp       *int
	Cursor    *string
	WithCount *bool
}

// GetNFTHoldersByContract retrieves NFT holders for a specific contract
func (c *Client) GetNFTHoldersByContract(contractAddress string, opts *NFTHoldersOptions) (*NFTHoldersResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contractAddress is required")
	}

	req := NFTHoldersByContractRequest{
		ContractAddress: contractAddress,
	}

	if opts != nil {
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
	}

	var result NFTHoldersResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftHoldersByContract", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetNFTHoldersByTokenId retrieves NFT holders for a specific token ID
func (c *Client) GetNFTHoldersByTokenId(contractAddress, tokenId string, opts *NFTHoldersOptions) (*NFTHoldersResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contractAddress is required")
	}
	if tokenId == "" {
		return nil, fmt.Errorf("tokenId is required")
	}

	req := NFTHoldersByTokenIdRequest{
		ContractAddress: contractAddress,
		TokenId:        tokenId,
	}

	if opts != nil {
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
	}

	var result NFTHoldersResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftHoldersByTokenId", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetNFTMetadataByContract retrieves NFT metadata for a specific contract
func (c *Client) GetNFTMetadataByContract(contractAddress string, opts *NFTHoldersOptions) (*NFTMetadataResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contractAddress is required")
	}

	req := NFTMetadataByContractRequest{
		ContractAddress: contractAddress,
	}

	if opts != nil {
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
	}

	var result NFTMetadataResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftMetadataByContract", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetNFTMetadataByTokenIds retrieves NFT metadata for specific token IDs
func (c *Client) GetNFTMetadataByTokenIds(tokens []TokenWithContract) ([]NFTMetadata, error) {
	if len(tokens) == 0 {
		return nil, fmt.Errorf("at least one token must be specified")
	}

	req := NFTMetadataByTokenIdsRequest{
		Tokens: tokens,
	}

	var result []NFTMetadata
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftMetadataByTokenIds", req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NFTTransfersByAccountOptions represents optional parameters for GetNFTTransfersByAccount
type NFTTransfersByAccountOptions struct {
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

// GetNFTTransfersByAccount retrieves NFT transfers for a specific account
func (c *Client) GetNFTTransfersByAccount(accountAddress string, opts *NFTTransfersByAccountOptions) (*NFTTransferResponse, error) {
	if accountAddress == "" {
		return nil, fmt.Errorf("accountAddress is required")
	}

	req := NFTTransfersByAccountRequest{
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

	var result NFTTransferResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftTransfersByAccount", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// NFTTransfersByContractOptions represents optional parameters for GetNFTTransfersByContract
type NFTTransfersByContractOptions struct {
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

// GetNFTTransfersByContract retrieves NFT transfers for a specific contract
func (c *Client) GetNFTTransfersByContract(contractAddress string, opts *NFTTransfersByContractOptions) (*NFTTransferResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contractAddress is required")
	}

	req := NFTTransfersByContractRequest{
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
		req.WithMetadata = opts.WithMetadata
		req.WithZeroValue = opts.WithZeroValue
	}

	var result NFTTransferResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftTransfersByContract", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetNFTTransfersByTokenId retrieves NFT transfers for a specific token ID
func (c *Client) GetNFTTransfersByTokenId(contractAddress, tokenId string, opts *NFTTransfersByContractOptions) (*NFTTransferResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contractAddress is required")
	}
	if tokenId == "" {
		return nil, fmt.Errorf("tokenId is required")
	}

	req := NFTTransfersByTokenIdRequest{
		ContractAddress: contractAddress,
		TokenId:        tokenId,
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
		req.WithMetadata = opts.WithMetadata
		req.WithZeroValue = opts.WithZeroValue
	}

	var result NFTTransferResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftTransfersByTokenId", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetNFTTransfersWithinRange retrieves NFT transfers within a specified range
func (c *Client) GetNFTTransfersWithinRange(opts *NFTTransfersByContractOptions) (*NFTTransferResponse, error) {
	req := NFTTransfersWithinRangeRequest{}

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

	var result NFTTransferResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftTransfersWithinRange", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// NFTsOwnedOptions represents optional parameters for GetNFTsOwnedByAccount
type NFTsOwnedOptions struct {
	ContractAddresses []string
	Page             *int
	Rpp              *int
	Cursor           *string
	WithCount        *bool
	WithMetadata     *bool
}

// GetNFTsOwnedByAccount retrieves NFTs owned by a specific account
func (c *Client) GetNFTsOwnedByAccount(accountAddress string, opts *NFTsOwnedOptions) (*NFTsOwnedResponse, error) {
	if accountAddress == "" {
		return nil, fmt.Errorf("accountAddress is required")
	}

	req := NFTsOwnedRequest{
		AccountAddress: accountAddress,
	}

	if opts != nil {
		req.ContractAddresses = opts.ContractAddresses
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
		req.WithMetadata = opts.WithMetadata
	}

	var result NFTsOwnedResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/getNftsOwnedByAccount", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// NFTContractMetadataSearchOptions represents optional parameters for SearchNFTContractMetadata
type NFTContractMetadataSearchOptions struct {
	Page      *int
	Rpp       *int
	Cursor    *string
	WithCount *bool
}

// SearchNFTContractMetadata searches for NFT contract metadata using a keyword
func (c *Client) SearchNFTContractMetadata(keyword string, opts *NFTContractMetadataSearchOptions) (*NFTContractMetadataSearchResponse, error) {
	if keyword == "" {
		return nil, fmt.Errorf("keyword is required")
	}

	req := NFTContractMetadataSearchRequest{
		Keyword: keyword,
	}

	if opts != nil {
		req.Page = opts.Page
		req.Rpp = opts.Rpp
		req.Cursor = opts.Cursor
		req.WithCount = opts.WithCount
	}

	var result NFTContractMetadataSearchResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/searchNftContractMetadataByKeyword", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SyncNFTMetadata submits NFT metadata for synchronization
func (c *Client) SyncNFTMetadata(tokens []TokenWithContract) (*SyncNFTMetadataResponse, error) {
	if len(tokens) == 0 {
		return nil, fmt.Errorf("tokens cannot be empty")
	}

	req := SyncNFTMetadataRequest{
		Tokens: tokens,
	}

	var result SyncNFTMetadataResponse
	err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/nft/syncNftMetadata", req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
