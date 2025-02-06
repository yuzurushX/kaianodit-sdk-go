package kaianodit

// NFTContractMetadataRequest represents the request body for getNftContractMetadataByContracts
type NFTContractMetadataRequest struct {
	ContractAddresses []string `json:"contractAddresses"`
}

// NFTContractMetadata represents a single NFT contract metadata
type NFTContractMetadata struct {
	Address                string `json:"address"`
	DeployedTransactionHash string `json:"deployedTransactionHash"`
	DeployedAt            string `json:"deployedAt"`
	DeployerAddress       string `json:"deployerAddress"`
	LogoURL              string `json:"logoUrl"`
	Type                 string `json:"type"`
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	TotalSupply          string `json:"totalSupply"`
	Decimals             int    `json:"decimals"`
}

// NFTContractsByAccountRequest represents the request body for getNftContractsByAccount
type NFTContractsByAccountRequest struct {
	AccountAddress    string   `json:"accountAddress"`              // Required
	ContractAddresses []string `json:"contractAddresses,omitempty"` // Optional
	Page             *int     `json:"page,omitempty"`              // Optional
	Rpp              *int     `json:"rpp,omitempty"`              // Optional
	Cursor           *string  `json:"cursor,omitempty"`            // Optional
	WithCount        *bool    `json:"withCount,omitempty"`         // Optional
}

// NFTContractBalance represents the balance information for a contract
type NFTContractBalance struct {
	TotalBalance  string `json:"totalBalance"`
	UniqueBalance string `json:"uniqueBalance"`
}

// NFTContractsByAccountResponse represents the response from getNftContractsByAccount
type NFTContractsByAccountResponse struct {
	Page    int                 `json:"page"`
	Rpp     int                `json:"rpp"`
	Cursor  string             `json:"cursor"`
	Count   *int               `json:"count,omitempty"` // Only present when withCount is true
	Items   []NFTContractBalance `json:"items"`
}

// NativeBalanceResponse represents the response from getNativeBalanceByAccount
type NativeBalanceResponse struct {
	OwnerAddress string `json:"ownerAddress"`
	Balance      string `json:"balance"`
}

// TokenAllowanceRequest represents the request for getTokenAllowance
type TokenAllowanceRequest struct {
	ContractAddress string `json:"contractAddress"`
	OwnerAddress   string `json:"ownerAddress"`
	SpenderAddress string `json:"spenderAddress"`
}

// TokenAllowanceResponse represents the response from getTokenAllowance
type TokenAllowanceResponse struct {
	Allowance string `json:"allowance"`
}

// TokenContractMetadataRequest represents the request for getTokenContractMetadataByContracts
type TokenContractMetadataRequest struct {
	ContractAddresses []string `json:"contractAddresses"`
}

// TokenContractMetadata represents a single token contract metadata
type TokenContractMetadata struct {
	Address                string `json:"address"`
	DeployedTransactionHash string `json:"deployedTransactionHash"`
	DeployedAt            string `json:"deployedAt"`
	DeployerAddress       string `json:"deployerAddress"`
	LogoURL              string `json:"logoUrl"`
	Type                 string `json:"type"`
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	TotalSupply          string `json:"totalSupply"`
	Decimals             int    `json:"decimals"`
}

// TokenHolder represents a single token holder's information
type TokenHolder struct {
	OwnerAddress string `json:"ownerAddress"`
	Balance      string `json:"balance"`
}

// TokenHoldersRequest represents the request for getTokenHoldersByContract
type TokenHoldersRequest struct {
	ContractAddress string `json:"contractAddress"`
	Page           *int   `json:"page,omitempty"`
	Rpp            *int   `json:"rpp,omitempty"`
	Cursor         *string `json:"cursor,omitempty"`
	WithCount      *bool   `json:"withCount,omitempty"`
}

// TokenHoldersResponse represents the response from getTokenHoldersByContract
type TokenHoldersResponse struct {
	Page    int           `json:"page,omitempty"`
	Rpp     int           `json:"rpp"`
	Cursor  string        `json:"cursor"`
	Count   *int          `json:"count,omitempty"` // Only present when withCount is true
	Items   []TokenHolder `json:"items"`
}

// TokenPriceContract represents the contract information in token price response
type TokenPriceContract struct {
	Address                string `json:"address"`
	DeployedTransactionHash string `json:"deployedTransactionHash"`
	DeployedAt            string `json:"deployedAt"`
	DeployerAddress       string `json:"deployerAddress"`
	LogoURL              string `json:"logoUrl"`
	Type                 string `json:"type"`
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	Decimals             int    `json:"decimals"`
	TotalSupply          string `json:"totalSupply"`
}

// TokenPrice represents the detailed price information for a token
type TokenPrice struct {
	Currency              string           `json:"currency"`
	Price                 string           `json:"price"`
	VolumeFor24h         string           `json:"volumeFor24h"`
	VolumeChangeFor24h   string           `json:"volumeChangeFor24h"`
	PercentChangeFor1h   string           `json:"percentChangeFor1h"`
	PercentChangeFor24h  string           `json:"percentChangeFor24h"`
	PercentChangeFor7d   string           `json:"percentChangeFor7d"`
	MarketCap            string           `json:"marketCap"`
	UpdatedAt            string           `json:"updatedAt"`
	Listings             []string         `json:"listings"`
	Contract             TokenPriceContract `json:"contract"`
}

// TokenPricesRequest represents the request for getTokenPricesByContracts
type TokenPricesRequest struct {
	ContractAddresses []string `json:"contractAddresses"`
}

type TokenContract struct{
	Name            string           `json:"name"`
	Symbol          string           `json:"symbol"`
	TotalSupply     string           `json:"totalSupply"`
	Decimals       	int64            `json:"decimals"`

}
// TokenTransfer represents a single token transfer record
type TokenTransfer struct {
	From            string           `json:"from"`
	To              string           `json:"to"`
	Value           string           `json:"value"`
	Timestamp       int64            `json:"timestamp"`
	BlockNumber     int64            `json:"blockNumber"`
	TransactionHash string           `json:"transactionHash"`
	LogIndex        int              `json:"logIndex"`
	Contract        TokenContract    `json:"contract"`
}

// TokenTransfersRequest represents the request for getTokenTransfersByAccount
type TokenTransfersRequest struct {
	AccountAddress    string   `json:"accountAddress"`              // Required
	Relation         string   `json:"relation,omitempty"`
	ContractAddresses []string `json:"contractAddresses,omitempty"`
	FromBlock        string   `json:"fromBlock,omitempty"`
	ToBlock          string   `json:"toBlock,omitempty"`
	FromDate         string   `json:"fromDate,omitempty"`
	ToDate           string   `json:"toDate,omitempty"`
	Page             *int     `json:"page,omitempty"`
	Rpp              *int     `json:"rpp,omitempty"`
	Cursor           *string  `json:"cursor,omitempty"`
	WithCount        *bool    `json:"withCount,omitempty"`
	WithMetadata     *bool    `json:"withMetadata,omitempty"`
	WithZeroValue    *bool    `json:"withZeroValue,omitempty"`
}

// TokenTransfersResponse represents the response from getTokenTransfersByAccount
type TokenTransfersResponse struct {
	Rpp    int            `json:"rpp"`
	Count  *int           `json:"count,omitempty"`
	Cursor string         `json:"cursor"`
	Items  []TokenTransfer `json:"items"`
}

// TokenTransfersByContractRequest represents the request for getTokenTransfersByContract
type TokenTransfersByContractRequest struct {
	ContractAddress string  `json:"contractAddress"`          // Required
	FromBlock      string  `json:"fromBlock,omitempty"`
	ToBlock        string  `json:"toBlock,omitempty"`
	FromDate       string  `json:"fromDate,omitempty"`
	ToDate         string  `json:"toDate,omitempty"`
	Page           *int    `json:"page,omitempty"`
	Rpp            *int    `json:"rpp,omitempty"`
	Cursor         *string `json:"cursor,omitempty"`
	WithCount      *bool   `json:"withCount,omitempty"`
	WithZeroValue  *bool   `json:"withZeroValue,omitempty"`
}

// TokenTransfersByContractResponse represents the response from getTokenTransfersByContract
type TokenTransfersByContractResponse struct {
	Rpp    int            `json:"rpp"`
	Count  *int           `json:"count,omitempty"`
	Cursor string         `json:"cursor"`
	Items  []TokenTransfer `json:"items"`
}

// TokenTransfersWithinRangeRequest represents the request for getTokenTransfersWithinRange
type TokenTransfersWithinRangeRequest struct {
	FromBlock     string  `json:"fromBlock,omitempty"`
	ToBlock       string  `json:"toBlock,omitempty"`
	FromDate      string  `json:"fromDate,omitempty"`
	ToDate        string  `json:"toDate,omitempty"`
	Page          *int    `json:"page,omitempty"`
	Rpp           *int    `json:"rpp,omitempty"`
	Cursor        *string `json:"cursor,omitempty"`
	WithCount     *bool   `json:"withCount,omitempty"`
	WithMetadata  *bool   `json:"withMetadata,omitempty"`
	WithZeroValue *bool   `json:"withZeroValue,omitempty"`
}

// TokenTransfersWithinRangeResponse represents the response from getTokenTransfersWithinRange
type TokenTransfersWithinRangeResponse struct {
	Rpp    int            `json:"rpp"`
	Count  *int           `json:"count,omitempty"`
	Cursor string         `json:"cursor"`
	Items  []TokenTransfer `json:"items"`
}

// TokensOwnedRequest represents the request for getTokensOwnedByAccount
type TokensOwnedRequest struct {
	AccountAddress   string  `json:"accountAddress"`
	ContractAddress string  `json:"contractAddress,omitempty"`
	Page            *int    `json:"page,omitempty"`
	Rpp             *int    `json:"rpp,omitempty"`
	Cursor          *string `json:"cursor,omitempty"`
	WithCount       *bool   `json:"withCount,omitempty"`
}

// TokenOwnership represents a single token ownership record
type TokenOwnership struct {
	OwnerAddress string        `json:"ownerAddress"`
	Balance      string        `json:"balance"`
	Contract     TokenContract `json:"contract"`
}

// TokensOwnedResponse represents the response from getTokensOwnedByAccount
type TokensOwnedResponse struct {
	Rpp    int             `json:"rpp"`
	Count  *int            `json:"count,omitempty"`
	Cursor string          `json:"cursor"`
	Items  []TokenOwnership `json:"items"`
}

// TokenContractMetadataSearchRequest represents the request for searchTokenContractMetadataByKeyword
type TokenContractMetadataSearchRequest struct {
	Keyword    string  `json:"keyword"`              // Required
	Page       *int    `json:"page,omitempty"`
	Rpp        *int    `json:"rpp,omitempty"`
	Cursor     *string `json:"cursor,omitempty"`
	WithCount  *bool   `json:"withCount,omitempty"`
}

// TokenContractMetadataSearchResponse represents the response from searchTokenContractMetadataByKeyword
type TokenContractMetadataSearchResponse struct {
	Rpp    int                   `json:"rpp"`
	Count  *int                  `json:"count,omitempty"`
	Cursor string                `json:"cursor"`
	Items  []TokenContractMetadata `json:"items"`
}

// NFTHolder represents a single NFT holder's information
type NFTHolder struct {
	OwnerAddress  string `json:"ownerAddress"`
	TotalBalance  string `json:"totalBalance"`
	UniqueBalance string `json:"uniqueBalance"`
}

// NFTHoldersByContractRequest represents the request for getNftHoldersByContract
type NFTHoldersByContractRequest struct {
	ContractAddress string  `json:"contractAddress"`
	Page           *int    `json:"page,omitempty"`
	Rpp            *int    `json:"rpp,omitempty"`
	Cursor         *string `json:"cursor,omitempty"`
	WithCount      *bool   `json:"withCount,omitempty"`
}

// NFTHoldersByTokenIdRequest represents the request for getNftHoldersByTokenId
type NFTHoldersByTokenIdRequest struct {
	ContractAddress string  `json:"contractAddress"`
	TokenId        string  `json:"tokenId"`
	Page           *int    `json:"page,omitempty"`
	Rpp            *int    `json:"rpp,omitempty"`
	Cursor         *string `json:"cursor,omitempty"`
	WithCount      *bool   `json:"withCount,omitempty"`
}

// NFTMetadata represents a single NFT metadata
type NFTMetadata struct {
	Contract         NFTContract `json:"contract"`
	TokenId         string      `json:"tokenId"`
	TokenUri        string      `json:"tokenUri"`
	TokenUriSyncedAt string      `json:"tokenUriSyncedAt"`
	RawMetadata     string      `json:"rawMetadata"`
	MetadataSyncedAt string      `json:"metadataSyncedAt"`
}

// NFTContract represents NFT contract information
type NFTContract struct {
	Address                string `json:"address"`
	DeployedTransactionHash string `json:"deployedTransactionHash"`
	DeployedAt            string `json:"deployedAt"`
	DeployerAddress       string `json:"deployerAddress"`
	LogoURL              *string `json:"logoUrl"`
	Type                 string `json:"type"`
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
}

// NFTMetadataByContractRequest represents the request for getNftMetadataByContract
type NFTMetadataByContractRequest struct {
	ContractAddress string  `json:"contractAddress"`
	Page           *int    `json:"page,omitempty"`
	Rpp            *int    `json:"rpp,omitempty"`
	Cursor         *string `json:"cursor,omitempty"`
	WithCount      *bool   `json:"withCount,omitempty"`
}

// NFTHoldersResponse represents the response for both holders endpoints
type NFTHoldersResponse struct {
	Rpp    int         `json:"rpp"`
	Count  *int        `json:"count,omitempty"`
	Cursor string      `json:"cursor"`
	Items  []NFTHolder `json:"items"`
}

// NFTMetadataResponse represents the response for getNftMetadataByContract
type NFTMetadataResponse struct {
	Rpp    int          `json:"rpp"`
	Count  *int         `json:"count,omitempty"`
	Cursor string       `json:"cursor"`
	Items  []NFTMetadata `json:"items"`
}

// TokenWithContract represents a token identifier with its contract address
type TokenWithContract struct {
	ContractAddress string `json:"contractAddress"`
	TokenId         string `json:"tokenId"`
}

// NFTMetadataByTokenIdsRequest represents the request for getNftMetadataByTokenIds
type NFTMetadataByTokenIdsRequest struct {
	Tokens []TokenWithContract `json:"tokens"`
}

// NFTTransfersByAccountRequest represents the request for getNftTransfersByAccount
type NFTTransfersByAccountRequest struct {
	AccountAddress     string   `json:"accountAddress"`
	Relation          string   `json:"relation,omitempty"`
	ContractAddresses []string `json:"contractAddresses,omitempty"`
	FromBlock         string   `json:"fromBlock,omitempty"`
	ToBlock           string   `json:"toBlock,omitempty"`
	FromDate          string   `json:"fromDate,omitempty"`
	ToDate            string   `json:"toDate,omitempty"`
	Page              *int     `json:"page,omitempty"`
	Rpp               *int     `json:"rpp,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"`
	WithCount         *bool    `json:"withCount,omitempty"`
	WithMetadata      *bool    `json:"withMetadata,omitempty"`
	WithZeroValue     *bool    `json:"withZeroValue,omitempty"`
}


// NFTTransfersByContractRequest represents the request for getNftTransfersByContract
type NFTTransfersByContractRequest struct {
	ContractAddress string  `json:"contractAddress"`
	FromBlock      string  `json:"fromBlock,omitempty"`
	ToBlock        string  `json:"toBlock,omitempty"`
	FromDate       string  `json:"fromDate,omitempty"`
	ToDate         string  `json:"toDate,omitempty"`
	Page           *int    `json:"page,omitempty"`
	Rpp            *int    `json:"rpp,omitempty"`
	Cursor         *string `json:"cursor,omitempty"`
	WithCount      *bool   `json:"withCount,omitempty"`
	WithMetadata   *bool   `json:"withMetadata,omitempty"`
	WithZeroValue  *bool   `json:"withZeroValue,omitempty"`
}

type NFTTransfersByTokenIdRequest struct {
	ContractAddress string  `json:"contractAddress"`
	TokenId         string  `json:"tokenId"`
	FromBlock      string  `json:"fromBlock,omitempty"`
	ToBlock        string  `json:"toBlock,omitempty"`
	FromDate       string  `json:"fromDate,omitempty"`
	ToDate         string  `json:"toDate,omitempty"`
	Page           *int    `json:"page,omitempty"`
	Rpp            *int    `json:"rpp,omitempty"`
	Cursor         *string `json:"cursor,omitempty"`
	WithCount      *bool   `json:"withCount,omitempty"`
	WithMetadata   *bool   `json:"withMetadata,omitempty"`
	WithZeroValue  *bool   `json:"withZeroValue,omitempty"`
}

type NFTTransfersWithinRangeRequest struct {
	FromBlock     string  `json:"fromBlock,omitempty"`
	ToBlock       string  `json:"toBlock,omitempty"`
	FromDate      string  `json:"fromDate,omitempty"`
	ToDate        string  `json:"toDate,omitempty"`
	Page          *int    `json:"page,omitempty"`
	Rpp           *int    `json:"rpp,omitempty"`
	Cursor        *string `json:"cursor,omitempty"`
	WithCount     *bool   `json:"withCount,omitempty"`
	WithMetadata  *bool   `json:"withMetadata,omitempty"`
	WithZeroValue *bool   `json:"withZeroValue,omitempty"`
}

// NFTTransferResponse represents the response for getNftTransfersByAccount
type NFTTransferResponse struct {
	Rpp    int           `json:"rpp"`
	Count  int           `json:"count,omitempty"`
	Cursor string        `json:"cursor,omitempty"`
	Items  []NFTTransfer `json:"items"`
}

type NFTTransfer struct {
	From            string      `json:"from"`
	To              string      `json:"to"`
	Value           string      `json:"value"`
	Timestamp       int64       `json:"timestamp"`
	BlockNumber     int64       `json:"blockNumber"`
	TransactionHash string      `json:"transactionHash"`
	LogIndex        int         `json:"logIndex"`
	BatchIndex      int         `json:"batchIndex,omitempty"`
	Contract        NFTContract `json:"contract"`
	NFT            *NFTMetadata `json:"nft,omitempty"`
}


type NFTItem struct {
	TokenId          string  `json:"tokenId"`
	TokenUri         *string `json:"tokenUri"`
	TokenUriSyncedAt *string `json:"tokenUriSyncedAt"`
	RawMetadata      *string `json:"rawMetadata,omitempty"`
	MetadataSyncedAt *string `json:"metadataSyncedAt,omitempty"`
}

// NFT Ownership Types
type NFTsOwnedRequest struct {
	AccountAddress    string   `json:"accountAddress"`
	ContractAddresses []string `json:"contractAddresses,omitempty"`
	Page             *int     `json:"page,omitempty"`
	Rpp              *int     `json:"rpp,omitempty"`
	Cursor           *string  `json:"cursor,omitempty"`
	WithCount        *bool    `json:"withCount,omitempty"`
	WithMetadata     *bool    `json:"withMetadata,omitempty"`
}

type NFTsOwnedResponse struct {
	Rpp    int           `json:"rpp"`
	Count  int           `json:"count,omitempty"`
	Cursor string        `json:"cursor,omitempty"`
	Items  []NFTOwned    `json:"items"`
}

type NFTOwned struct {
	Contract         NFTContract `json:"contract"`
	OwnerAddress     string      `json:"ownerAddress"`
	Balance          string      `json:"balance"`
	TokenId          string      `json:"tokenId"`
	TokenUri         string      `json:"tokenUri"`
	TokenUriSyncedAt string      `json:"tokenUriSyncedAt"`
	RawMetadata      *string     `json:"rawMetadata"`
	MetadataSyncedAt *string     `json:"metadataSyncedAt"`
}

// NFT Contract Metadata Search Types
type NFTContractMetadataSearchRequest struct {
	Keyword    string  `json:"keyword"`
	Page       *int    `json:"page,omitempty"`
	Rpp        *int    `json:"rpp,omitempty"`
	Cursor     *string `json:"cursor,omitempty"`
	WithCount  *bool   `json:"withCount,omitempty"`
}

type NFTContractMetadataSearchResponse struct {
	Rpp    int           `json:"rpp"`
	Count  int           `json:"count,omitempty"`
	Cursor string        `json:"cursor,omitempty"`
	Items  []NFTContract `json:"items"`
}

// NFT Metadata Sync Types
type SyncNFTMetadataRequest struct {
	Tokens []TokenWithContract `json:"tokens"`
}

type SyncNFTMetadataResponse struct {
	Message string `json:"message"`
}

// Blockchain Types
type BlockRequest struct {
	Block string `json:"block"`
}

type Block struct {
	Hash              string       `json:"hash"`
	Number            int64        `json:"number"`
	Timestamp         int64        `json:"timestamp"`
	ParentHash        string       `json:"parentHash"`
	Nonce            string       `json:"nonce"`
	StateRoot        string       `json:"stateRoot"`
	ReceiptsRoot     string       `json:"receiptsRoot"`
	TransactionsRoot string       `json:"transactionsRoot"`
	Miner            string       `json:"miner"`
	Difficulty       string       `json:"difficulty"`
	TotalDifficulty  string       `json:"totalDifficulty"`
	MixHash          string       `json:"mixHash"`
	GasLimit         string       `json:"gasLimit"`
	GasUsed          string       `json:"gasUsed"`
	BaseFeePerGas    string       `json:"baseFeePerGas"`
	BlobGasUsed      string       `json:"blobGasUsed,omitempty"`
	Size             string       `json:"size"`
	LogsBloom        string       `json:"logsBloom"`
	ExtraData        string       `json:"extraData"`
	Sha3Uncles       string       `json:"sha3Uncles"`
	TransactionsCount int         `json:"transactionsCount"`
	Transactions     []string     `json:"transactions"`
	WithdrawalsCount int         `json:"withdrawalsCount,omitempty"`
	WithdrawalsRoot  string       `json:"withdrawalsRoot,omitempty"`
	Withdrawals      []Withdrawal `json:"withdrawals,omitempty"`
}

type Withdrawal struct {
	Index          string `json:"index"`
	ValidatorIndex string `json:"validatorIndex"`
	Address        string `json:"address"`
	Amount         string `json:"amount"`
}

type BlocksWithinRangeRequest struct {
	FromBlock string  `json:"fromBlock,omitempty"`
	ToBlock   string  `json:"toBlock,omitempty"`
	FromDate  string  `json:"fromDate,omitempty"`
	ToDate    string  `json:"toDate,omitempty"`
	Page      *int    `json:"page,omitempty"`
	Rpp       *int    `json:"rpp,omitempty"`
	Cursor    *string `json:"cursor,omitempty"`
	WithCount *bool   `json:"withCount,omitempty"`
}

type BlocksResponse struct {
	Page    int     `json:"page,omitempty"`
	Rpp     int     `json:"rpp"`
	Cursor  string  `json:"cursor,omitempty"`
	Count   int     `json:"count,omitempty"`
	Items   []Block `json:"items"`
}

type GasPrice struct {
	High        string `json:"high"`
	Average     string `json:"average"`
	Low         string `json:"low"`
	LatestBlock string `json:"latestBlock"`
	BaseFee     string `json:"baseFee"`
}

// Account Request Types
type AccountRequest struct {
	AccountAddress string `json:"accountAddress"`
}

type NonceResponse struct {
	Nonce string `json:"nonce"`
}

type TransactionCountResponse struct {
	TransactionCount int `json:"transactionCount"`
}

// Transaction Types
type TransactionRequest struct {
	TransactionHash string `json:"transactionHash"`
	WithLogs        *bool  `json:"withLogs,omitempty"`
	WithDecode      *bool  `json:"withDecode,omitempty"`
}

type Transaction struct {
	TransactionHash       string           `json:"transactionHash"`
	TransactionIndex     string           `json:"transactionIndex"`
	BlockHash           string           `json:"blockHash"`
	BlockNumber         int64            `json:"blockNumber"`
	From                string           `json:"from"`
	To                  string           `json:"to"`
	Value               string           `json:"value"`
	Input               string           `json:"input"`
	DecodedInput        *DecodedInput     `json:"decodedInput,omitempty"`
	FunctionSelector    string           `json:"functionSelector"`
	Nonce              string           `json:"nonce"`
	Gas                string           `json:"gas"`
	GasPrice           string           `json:"gasPrice"`
	MaxFeePerGas       string           `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string         `json:"maxPriorityFeePerGas"`
	GasUsed            string           `json:"gasUsed"`
	CumulativeGasUsed  string           `json:"cumulativeGasUsed"`
	EffectiveGasPrice  string           `json:"effectiveGasPrice"`
	ContractAddress    string           `json:"contractAddress"`
	Type               string           `json:"type"`
	Status             string           `json:"status"`
	LogsBloom          string           `json:"logsBloom"`
	AccessList         []AccessListItem  `json:"accessList"`
	Timestamp         int64            `json:"timestamp"`
	Logs              []TransactionLog  `json:"logs,omitempty"`
}

type DecodedInput struct {
	Type      string     `json:"type"`
	Name      string     `json:"name"`
	Signature string     `json:"signature"`
	Args      []Argument `json:"args"`
}

type Argument struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AccessListItem struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

type TransactionLog struct {
	Name          string     `json:"name"`
	EventFragment string     `json:"eventFragment"`
	Signature     string     `json:"signature"`
	EventHash     string     `json:"eventHash"`
	Args          []Argument `json:"args"`
}

// IsContract Types
type IsContractRequest struct {
	Address string `json:"address"`
}

type IsContractResponse struct {
	Result bool `json:"result"`
}

// SearchEvents Types
type SearchEventsRequest struct {
	ContractAddress string      `json:"contractAddress"`
	EventNames      []string    `json:"eventNames"`
	ABI             interface{} `json:"abi"`
	FromBlock       string      `json:"fromBlock,omitempty"`
	ToBlock         string      `json:"toBlock,omitempty"`
	FromDate        string      `json:"fromDate,omitempty"`
	ToDate          string      `json:"toDate,omitempty"`
	Page            *int        `json:"page,omitempty"`
	Rpp             *int        `json:"rpp,omitempty"`
	Cursor          *string     `json:"cursor,omitempty"`
}

type EventsResponse struct {
	Page    int     `json:"page,omitempty"`
	Rpp     int     `json:"rpp"`
	Cursor  string  `json:"cursor,omitempty"`
	Count   int     `json:"count,omitempty"`
	Items   []Event `json:"items"`
}

type Event struct {
	Logs EventLog `json:"logs"`
}

type EventLog struct {
	ContractAddress   string   `json:"contractAddress"`
	TransactionHash   string   `json:"transactionHash"`
	TransactionIndex  string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      int64    `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
}

// Account Statistics Types
type AccountStatsRequest struct {
	Address string `json:"address"`
}

type AccountStats struct {
	TransactionCounts TransactionCounts `json:"transactionCounts"`
	TransferCounts    TransferCounts    `json:"transferCounts"`
	Assets            Assets            `json:"assets"`
}

type TransactionCounts struct {
	External int `json:"external"`
	Internal int `json:"internal"`
}

type TransferCounts struct {
	Tokens int `json:"tokens"`
	NFTs   int `json:"nfts"`
}

type Assets struct {
	Tokens        int `json:"tokens"`
	NFTs          int `json:"nfts"`
	NFTContracts int `json:"nftContracts"`
}

// JSONRPCRequest represents a JSON-RPC 2.0 request
type JSONRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int          `json:"id"`
}

// JSONRPCResponse represents a JSON-RPC 2.0 response
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
	ID      int             `json:"id"`
}

// RPCError represents a JSON-RPC 2.0 error
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Ethereum JSON-RPC specific types
type CallRequest struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value,omitempty"`
	Data     string `json:"data,omitempty"`
}

type AccessListEntry struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

type AccessList struct {
	AccessList []AccessListEntry `json:"accessList"`
	GasUsed    string           `json:"gasUsed"`
}

type FeeHistoryResult struct {
	OldestBlock  string     `json:"oldestBlock"`
	BaseFeePerGas []string   `json:"baseFeePerGas"`
	GasUsedRatio []float64  `json:"gasUsedRatio"`
	Reward      [][]string  `json:"reward,omitempty"`
}

type Log struct {
	Removed          bool     `json:"removed"`
	LogIndex         string   `json:"logIndex"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}

type FilterOptions struct {
	FromBlock string     `json:"fromBlock,omitempty"`
	ToBlock   string     `json:"toBlock,omitempty"`
	Address   string     `json:"address,omitempty"`
	Topics    []string   `json:"topics,omitempty"`
}

type ProofResponse struct {
	Address      string          `json:"address"`
	Balance      string          `json:"balance"`
	CodeHash     string          `json:"codeHash"`
	Nonce        string          `json:"nonce"`
	StorageHash  string          `json:"storageHash"`
	AccountProof []string        `json:"accountProof"`
	StorageProof []StorageProof  `json:"storageProof"`
}

type StorageProof struct {
	Key   string   `json:"key"`
	Value string   `json:"value"`
	Proof []string `json:"proof"`
}

type TransactionReceipt struct {
	TransactionHash   string   `json:"transactionHash"`
	TransactionIndex  string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	CumulativeGasUsed string   `json:"cumulativeGasUsed"`
	GasUsed          string   `json:"gasUsed"`
	ContractAddress  string   `json:"contractAddress"`
	Logs             []Log    `json:"logs"`
	LogsBloom        string   `json:"logsBloom"`
	Status           string   `json:"status"`
	EffectiveGasPrice string  `json:"effectiveGasPrice"`
	Type             string   `json:"type"`
}

// Update JSONRPCResponse to handle different result types
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
	ID      int             `json:"id"`
}

// Webhook Types
type WebhookEventType string

const (
	EventTypeAddressActivity      WebhookEventType = "ADDRESS_ACTIVITY"
	EventTypeMinedTransaction     WebhookEventType = "MINED_TRANSACTION"
	EventTypeSuccessfulTransaction WebhookEventType = "SUCCESSFUL_TRANSACTION"
	EventTypeFailedTransaction    WebhookEventType = "FAILED_TRANSACTION"
	EventTypeTokenTransfer        WebhookEventType = "TOKEN_TRANSFER"
	EventTypeBelowThresholdBalance WebhookEventType = "BELOW_THRESHOLD_BALANCE"
	EventTypeBlockPeriod          WebhookEventType = "BLOCK_PERIOD"
	EventTypeBlockListCaller      WebhookEventType = "BLOCK_LIST_CALLER"
	EventTypeAllowListCaller      WebhookEventType = "ALLOW_LIST_CALLER"
	EventTypeLog                  WebhookEventType = "LOG"
)

type WebhookNotification struct {
	WebhookURL string `json:"webhookUrl"`
}

// TokenCondition represents a token contract condition
type TokenCondition struct {
	ContractAddress string `json:"contractAddress"`
}

// WebhookCondition represents different types of conditions based on event type
type WebhookCondition struct {
	// For ADDRESS_ACTIVITY, MINED_TRANSACTION, SUCCESSFUL_TRANSACTION, FAILED_TRANSACTION
	Addresses []string `json:"addresses,omitempty"`

	// For TOKEN_TRANSFER
	Tokens []TokenCondition `json:"tokens,omitempty"`

	// For BELOW_THRESHOLD_BALANCE
	Address                string `json:"address,omitempty"`
	BelowThresholdBalance string `json:"belowThresholdBalance,omitempty"`

	// For BLOCK_PERIOD
	Period int `json:"period,omitempty"`

	// For BLOCK_LIST_CALLER
	BlockListCallers []string `json:"blockListCallers,omitempty"`

	// For ALLOW_LIST_CALLER
	AllowListCallers []string `json:"allowListCallers,omitempty"`

	// For LOG
	Address string   `json:"address,omitempty"` // Contract address for LOG event type
	Topics  []string `json:"topics,omitempty"`
}

type CreateWebhookRequest struct {
	EventType    WebhookEventType    `json:"eventType"`
	Description  string              `json:"description"`
	Notification WebhookNotification `json:"notification"`
	Condition    WebhookCondition    `json:"condition"`
}

type Webhook struct {
	SubscriptionID string              `json:"subscriptionId"`
	EventType      WebhookEventType    `json:"eventType"`
	Description    string              `json:"description"`
	Notification   WebhookNotification `json:"notification"`
	Condition      WebhookCondition    `json:"condition"`
	CreatedAt      string              `json:"createdAt,omitempty"`
	UpdatedAt      string              `json:"updatedAt,omitempty"`
}

type ListWebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type WebhookResponse struct {
	Message string  `json:"message,omitempty"`
	Data    Webhook `json:"data,omitempty"`
}

type UpdateWebhookRequest struct {
	Description  string              `json:"description,omitempty"`
	Notification WebhookNotification `json:"notification,omitempty"`
	Condition    WebhookCondition    `json:"condition,omitempty"`
} 