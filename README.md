# Kaia Nodit SDK for Go

The Kaia Nodit SDK for Go provides a convenient way to interact with the Kaia Nodit API from Go applications. This SDK supports all Kaia Nodit API endpoints and provides a simple, intuitive interface for blockchain data retrieval. It supports both Kaia mainnet and testnet networks.

## Features

- Complete coverage of Kaia Nodit API endpoints
- Support for both mainnet and testnet networks
- Easy-to-use interface
- Proper error handling
- Pagination support
- Comprehensive documentation
- Type-safe responses
- Node (JSON-RPC) support for direct blockchain queries
- Webhook integration for real-time notifications

## Installation

To install the SDK, use `go get`:
```
go get github.com/yuzurushX/kaianodit-go
```

## Usage

To use the SDK, import it into your Go code:
```go
import "github.com/yuzurushX/kaianodit-go"

// Create a client for mainnet (default)
client := kaianodit.NewClient("your-api-key")

// Or create a client for testnet
client := kaianodit.NewClient("your-api-key", kaianodit.Network(kaianodit.Testnet))
```

## Documentation

For detailed documentation on each endpoint, refer to the [Nodit API Documentation](https://developer.nodit.io/reference/kaia-getnftcontractmetadatabycontracts)

## Available Methods

### Node (JSON-RPC)
- `GetBalance` - Get account balance
- `GetBlockNumber` - Get current block number
- `Call` - Execute a new message call
- `GetChainID` - Get chain ID
- `CreateAccessList` - Create an access list
- `EstimateGas` - Estimate gas for transaction
- `GetFeeHistory` - Get historical fee information
- `GetBlockByHash` - Get block by hash
- `GetBlockByNumber` - Get block by number
- `GetBlockReceipts` - Get block receipts
- `GetBlockTransactionCountByHash` - Get transaction count by block hash
- `GetBlockTransactionCountByNumber` - Get transaction count by block number
- `GetCode` - Get contract code
- `GetFilterChanges` - Get filter changes
- `GetFilterLogs` - Get filter logs
- `GetLogs` - Get logs matching filter
- `GetProof` - Get account/storage proof
- `GetStorageAt` - Get storage at position
- `GetTransactionByBlockHashAndIndex` - Get transaction by block hash and index
- `GetTransactionByBlockNumberAndIndex` - Get transaction by block number and index
- `GetTransactionByHash` - Get transaction by hash
- `GetTransactionCount` - Get transaction count
- `GetTransactionReceipt` - Get transaction receipt
- `GetMaxPriorityFeePerGas` - Get max priority fee
- `NewBlockFilter` - Create new block filter
- `NewFilter` - Create new log filter
- `NewPendingTransactionFilter` - Create new pending transaction filter
- `SendRawTransaction` - Send raw transaction
- `UninstallFilter` - Uninstall filter

### Blockchain
- `GetBlockByHashOrNumber` - Retrieve block information by hash or number
- `GetBlocksWithinRange` - Get blocks within a specified range
- `GetGasPrice` - Get current gas price
- `GetNextNonceByAccount` - Get next nonce for an account
- `GetTotalTransactionCountByAccount` - Get total transaction count for an account
- `GetTransactionByHash` - Get transaction details by hash
- `IsContract` - Check if an address is a contract
- `SearchEvents` - Search for specific events

### Token
- `GetNativeBalance` - Get native token balance
- `GetTokenAllowance` - Get token allowance
- `GetTokenContractMetadata` - Get token contract metadata
- `GetTokenHolders` - Get token holders
- `GetTokenPrices` - Get token prices
- `GetTokenTransfers` - Get token transfers
- `GetTokenTransfersByContract` - Get transfers for specific contract
- `GetTokenTransfersWithinRange` - Get transfers within time/block range
- `GetTokensOwned` - Get tokens owned by account
- `SearchTokenContractMetadata` - Search token contracts

### NFT
- `GetNFTContractMetadata` - Get NFT contract metadata
- `GetNFTContractsByAccount` - Get NFT contracts by account
- `GetNFTHoldersByContract` - Get NFT holders for contract
- `GetNFTHoldersByTokenId` - Get holders for specific token
- `GetNFTMetadataByContract` - Get metadata for contract
- `GetNFTMetadataByTokenIds` - Get metadata for specific tokens
- `GetNFTTransfersByAccount` - Get NFT transfers by account
- `GetNFTTransfersByContract` - Get transfers for contract
- `GetNFTTransfersByTokenId` - Get transfers for specific token
- `GetNFTTransfersWithinRange` - Get transfers within range
- `