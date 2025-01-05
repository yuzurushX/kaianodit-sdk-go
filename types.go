package nodit

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