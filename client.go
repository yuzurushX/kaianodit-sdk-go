package nodit

import (
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://web3.nodit.io"
	defaultTimeout = 30 * time.Second
)

// Client represents the Nodit API client
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// ClientOption is a function that modifies the client
type ClientOption func(*Client)

// NewClient creates a new Nodit API client
func NewClient(apiKey string, opts ...ClientOption) *Client {
	client := &Client{
		baseURL: defaultBaseURL,
		apiKey:  apiKey,
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