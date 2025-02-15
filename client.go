package kaianodit

import (
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://web3.nodit.io"
	defaultTimeout = 30 * time.Second
)

// Network types
const (
	Mainnet = "mainnet"
	Testnet = "kairos"
)

// Node URLs for different networks
const (
	mainnetNodeURL = "https://kaia-mainnet.nodit.io"
	testnetNodeURL = "https://kaia-kairos.nodit.io"
)

// Client represents the Nodit API client
type Client struct {
	baseURL    string
	nodeURL    string
	apiKey     string
	network    string
	httpClient *http.Client
}

// ClientOption is a function that modifies the client
type ClientOption func(*Client)

// NewClient creates a new Nodit API client
func NewClient(apiKey string, opts ...ClientOption) *Client {
	client := &Client{
		baseURL: defaultBaseURL,
		nodeURL: mainnetNodeURL, // Default to mainnet node URL
		apiKey:  apiKey,
		network: Mainnet, // Default to mainnet
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	// Apply any custom options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Network sets the network (Mainnet or Testnet)
func Network(network string) ClientOption {
	return func(c *Client) {
		c.network = network
		// Update node URL based on network
		if network == Testnet {
			c.nodeURL = testnetNodeURL
		} else {
			c.nodeURL = mainnetNodeURL
		}
	}
}

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
} 