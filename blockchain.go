package kaianodit

import (
    "encoding/json"
    "fmt"
)

// GetBlockByHashOrNumber retrieves block information by hash or number
func (c *Client) GetBlockByHashOrNumber(block string) (*Block, error) {
    if block == "" {
        return nil, fmt.Errorf("block is required")
    }

    req := BlockRequest{
        Block: block,
    }

    var result Block
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/blockchain/getBlockByHashOrNumber", req, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

// BlocksWithinRangeOptions represents optional parameters for GetBlocksWithinRange
type BlocksWithinRangeOptions struct {
    FromBlock string
    ToBlock   string
    FromDate  string
    ToDate    string
    Page      *int
    Rpp       *int
    Cursor    *string
    WithCount *bool
}

// GetBlocksWithinRange retrieves blocks within a specified range
func (c *Client) GetBlocksWithinRange(opts *BlocksWithinRangeOptions) (*BlocksResponse, error) {
    req := BlocksWithinRangeRequest{}

    if opts != nil {
        req.FromBlock = opts.FromBlock
        req.ToBlock = opts.ToBlock
        req.FromDate = opts.FromDate
        req.ToDate = opts.ToDate
        req.Page = opts.Page
        req.Rpp = opts.Rpp
        req.Cursor = opts.Cursor
        req.WithCount = opts.WithCount
    }

    var result BlocksResponse
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/blockchain/getBlocksWithinRange", req, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

// GetGasPrice retrieves current gas price information
func (c *Client) GetGasPrice() (*GasPrice, error) {
    var result GasPrice
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/blockchain/getGasPrice", nil, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

// GetNextNonceByAccount retrieves the next nonce for a specific account
func (c *Client) GetNextNonceByAccount(accountAddress string) (*NonceResponse, error) {
    if accountAddress == "" {
        return nil, fmt.Errorf("accountAddress is required")
    }

    req := AccountRequest{
        AccountAddress: accountAddress,
    }

    var result NonceResponse
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/blockchain/getNextNonceByAccount", req, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

// GetTotalTransactionCountByAccount retrieves the total transaction count for a specific account
func (c *Client) GetTotalTransactionCountByAccount(accountAddress string) (*TransactionCountResponse, error) {
    if accountAddress == "" {
        return nil, fmt.Errorf("accountAddress is required")
    }

    req := AccountRequest{
        AccountAddress: accountAddress,
    }

    var result TransactionCountResponse
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/blockchain/getTotalTransactionCountByAccount", req, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

// TransactionOptions represents optional parameters for GetTransactionByHash
type TransactionOptions struct {
    WithLogs   *bool
    WithDecode *bool
}

// GetTransactionByHash retrieves transaction details by hash
func (c *Client) GetTransactionByHash(transactionHash string, opts *TransactionOptions) (*Transaction, error) {
    if transactionHash == "" {
        return nil, fmt.Errorf("transactionHash is required")
    }

    req := TransactionRequest{
        TransactionHash: transactionHash,
    }

    if opts != nil {
        req.WithLogs = opts.WithLogs
        req.WithDecode = opts.WithDecode
    }

    var result Transaction
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/blockchain/getTransactionByHash", req, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

// IsContract checks if an address is a contract
func (c *Client) IsContract(address string) (*IsContractResponse, error) {
    if address == "" {
        return nil, fmt.Errorf("address is required")
    }

    req := IsContractRequest{
        Address: address,
    }

    var result IsContractResponse
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/blockchain/isContract", req, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

// SearchEventsOptions represents optional parameters for SearchEvents
type SearchEventsOptions struct {
    FromBlock string
    ToBlock   string
    FromDate  string
    ToDate    string
    Page      *int
    Rpp       *int
    Cursor    *string
}

// SearchEvents searches for contract events
func (c *Client) SearchEvents(contractAddress string, eventNames []string, abi interface{}, opts *SearchEventsOptions) (*EventsResponse, error) {
    if contractAddress == "" {
        return nil, fmt.Errorf("contractAddress is required")
    }
    if len(eventNames) == 0 {
        return nil, fmt.Errorf("eventNames is required")
    }
    if abi == nil {
        return nil, fmt.Errorf("abi is required")
    }

    req := SearchEventsRequest{
        ContractAddress: contractAddress,
        EventNames:     eventNames,
        ABI:            abi,
    }

    if opts != nil {
        req.FromBlock = opts.FromBlock
        req.ToBlock = opts.ToBlock
        req.FromDate = opts.FromDate
        req.ToDate = opts.ToDate
        req.Page = opts.Page
        req.Rpp = opts.Rpp
        req.Cursor = opts.Cursor
    }

    var result EventsResponse
    err := c.doRequestAndDecode("POST", "/v1/kaia/mainnet/blockchain/searchEvents", req, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
} 