package kaianodit

import (
	"encoding/json"
	"fmt"
)

// GetBalance returns the balance of the account of given address in wei
func (c *Client) GetBalance(address string, block string) (string, error) {
	if block == "" {
		block = "latest"
	}

	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getBalance",
		Params:  []interface{}{address, block},
		ID:      1,
	}

	var response JSONRPCResponse
	err := c.doRequestAndDecode("POST", c.nodeURL, request, &response)
	if err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", fmt.Errorf("RPC error: %s", response.Error.Message)
	}

	return response.Result, nil
}

// GetGasPrice returns the current gas price in wei
func (c *Client) GetGasPrice() (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_gasPrice",
		Params:  []interface{}{},
		ID:      1,
	}

	var response JSONRPCResponse
	err := c.doRequestAndDecode("POST", c.nodeURL, request, &response)
	if err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", fmt.Errorf("RPC error: %s", response.Error.Message)
	}

	return response.Result, nil
}

// GetBlockNumber returns the current block number
func (c *Client) GetBlockNumber() (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_blockNumber",
		Params:  []interface{}{},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// Call executes a new message call immediately without creating a transaction
func (c *Client) Call(callRequest CallRequest, block string) (string, error) {
	if block == "" {
		block = "latest"
	}

	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_call",
		Params:  []interface{}{callRequest, block},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// GetChainID returns the chain ID of the current network
func (c *Client) GetChainID() (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_chainId",
		Params:  []interface{}{},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// CreateAccessList creates an access list for a transaction
func (c *Client) CreateAccessList(callRequest CallRequest, block string) (*AccessList, error) {
	if block == "" {
		block = "latest"
	}

	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_createAccessList",
		Params:  []interface{}{callRequest, block},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result AccessList
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// EstimateGas estimates the gas needed for a transaction
func (c *Client) EstimateGas(callRequest CallRequest) (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_estimateGas",
		Params:  []interface{}{callRequest},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// GetFeeHistory returns a collection of historical gas information
func (c *Client) GetFeeHistory(blockCount int, newestBlock string, rewardPercentiles []float64) (*FeeHistoryResult, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_feeHistory",
		Params:  []interface{}{blockCount, newestBlock, rewardPercentiles},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result FeeHistoryResult
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetBlockByHash returns information about a block by hash
func (c *Client) GetBlockByHash(hash string, fullTransactions bool) (*Block, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getBlockByHash",
		Params:  []interface{}{hash, fullTransactions},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result Block
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetBlockByNumber returns information about a block by number
func (c *Client) GetBlockByNumber(blockNumber string, fullTransactions bool) (*Block, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getBlockByNumber",
		Params:  []interface{}{blockNumber, fullTransactions},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result Block
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetBlockReceipts returns all transaction receipts for a given block
func (c *Client) GetBlockReceipts(blockHash string) ([]TransactionReceipt, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getBlockReceipts",
		Params:  []interface{}{blockHash},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result []TransactionReceipt
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetBlockTransactionCountByHash returns the number of transactions in a block by block hash
func (c *Client) GetBlockTransactionCountByHash(blockHash string) (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getBlockTransactionCountByHash",
		Params:  []interface{}{blockHash},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// GetBlockTransactionCountByNumber returns the number of transactions in a block by block number
func (c *Client) GetBlockTransactionCountByNumber(blockNumber string) (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getBlockTransactionCountByNumber",
		Params:  []interface{}{blockNumber},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// GetCode returns code at a given address
func (c *Client) GetCode(address string, block string) (string, error) {
	if block == "" {
		block = "latest"
	}

	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getCode",
		Params:  []interface{}{address, block},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// GetFilterChanges returns an array of logs which occurred since last poll
func (c *Client) GetFilterChanges(filterID string) ([]Log, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getFilterChanges",
		Params:  []interface{}{filterID},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result []Log
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetFilterLogs returns an array of all logs matching filter with given id
func (c *Client) GetFilterLogs(filterID string) ([]Log, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getFilterLogs",
		Params:  []interface{}{filterID},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result []Log
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetLogs returns an array of all logs matching a given filter object
func (c *Client) GetLogs(filterOptions FilterOptions) ([]Log, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getLogs",
		Params:  []interface{}{filterOptions},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result []Log
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetProof returns the account and storage values of the specified account including the Merkle-proof
func (c *Client) GetProof(address string, storageKeys []string, block string) (*ProofResponse, error) {
	if block == "" {
		block = "latest"
	}

	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getProof",
		Params:  []interface{}{address, storageKeys, block},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result ProofResponse
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetStorageAt returns the value from a storage position at a given address
func (c *Client) GetStorageAt(address string, position string, block string) (string, error) {
	if block == "" {
		block = "latest"
	}

	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getStorageAt",
		Params:  []interface{}{address, position, block},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// GetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position
func (c *Client) GetTransactionByBlockHashAndIndex(blockHash string, index string) (*Transaction, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getTransactionByBlockHashAndIndex",
		Params:  []interface{}{blockHash, index},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result Transaction
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position
func (c *Client) GetTransactionByBlockNumberAndIndex(blockNumber string, index string) (*Transaction, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getTransactionByBlockNumberAndIndex",
		Params:  []interface{}{blockNumber, index},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result Transaction
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTransactionByHash returns the information about a transaction requested by transaction hash
func (c *Client) GetTransactionByHash(hash string) (*Transaction, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getTransactionByHash",
		Params:  []interface{}{hash},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result Transaction
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTransactionCount returns the number of transactions sent from an address
func (c *Client) GetTransactionCount(address string, block string) (string, error) {
	if block == "" {
		block = "latest"
	}

	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getTransactionCount",
		Params:  []interface{}{address, block},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// GetTransactionReceipt returns the receipt of a transaction by transaction hash
func (c *Client) GetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_getTransactionReceipt",
		Params:  []interface{}{hash},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return nil, err
	}

	var result TransactionReceipt
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetMaxPriorityFeePerGas returns the current maxPriorityFeePerGas
func (c *Client) GetMaxPriorityFeePerGas() (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_maxPriorityFeePerGas",
		Params:  []interface{}{},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// NewBlockFilter creates a filter in the node, to notify when a new block arrives
func (c *Client) NewBlockFilter() (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_newBlockFilter",
		Params:  []interface{}{},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// NewFilter creates a filter object, based on filter options, to notify when the state changes
func (c *Client) NewFilter(filterOptions FilterOptions) (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_newFilter",
		Params:  []interface{}{filterOptions},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// NewPendingTransactionFilter creates a filter in the node, to notify when new pending transactions arrive
func (c *Client) NewPendingTransactionFilter() (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_newPendingTransactionFilter",
		Params:  []interface{}{},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// SendRawTransaction sends a signed transaction
func (c *Client) SendRawTransaction(signedTransactionData string) (string, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_sendRawTransaction",
		Params:  []interface{}{signedTransactionData},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return "", err
	}

	return result, nil
}

// UninstallFilter uninstalls a filter with given ID
func (c *Client) UninstallFilter(filterID string) (bool, error) {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "kaia_uninstallFilter",
		Params:  []interface{}{filterID},
		ID:      1,
	}

	var response JSONRPCResponse
	if err := c.doRequestAndDecode("POST", c.nodeURL, request, &response); err != nil {
		return false, err
	}

	var result bool
	if err := json.Unmarshal(response.Result, &result); err != nil {
		return false, err
	}

	return result, nil
} 
