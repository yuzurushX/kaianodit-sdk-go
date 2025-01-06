# Kaia Nodit SDK for Go

The Kaia Nodit SDK for Go provides a convenient way to interact with the Kaia Nodit API from Go applications. This SDK supports all Kaia Nodit API endpoints and provides a simple, intuitive interface for blockchain data retrieval.

## Features

- Complete coverage of Kaia Nodit API endpoints
- Easy-to-use interface
- Proper error handling
- Pagination support
- Comprehensive documentation
- Type-safe responses

## Installation

To install the SDK, use `go get`:
```
go get github.com/yuzurushX/kaianodit-go
```

## Usage

To use the SDK, import it into your Go code:
```
import "github.com/yuzurushX/kaianodit-go"
```

## Documentation

For detailed documentation on each endpoint, refer to the [Kaian ODit API documentation](https://apidoc.kaian.io/docs/kaian-odit-api).

## Available Methods

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
- `GetNFTsOwnedByAccount` - Get NFTs owned by account
- `SearchNFTContractMetadata` - Search NFT contracts
- `SyncNFTMetadata` - Sync NFT metadata

### Statistics
- `GetAccountStats` - Get account statistics

## Rate Limiting

The SDK respects the API's rate limits. When limits are exceeded, the API will return appropriate error responses.
