package kaianodit

import (
	"fmt"
	"net/url"
)

// CreateWebhook creates a new webhook subscription
func (c *Client) CreateWebhook(request CreateWebhookRequest) (*WebhookResponse, error) {
	var response WebhookResponse
	path := fmt.Sprintf("/v1/kaia/%s/webhooks", c.network)
	err := c.doRequestAndDecode("POST", path, request, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook: %w", err)
	}
	return &response, nil
}

// GetWebhook retrieves a webhook by subscription ID
func (c *Client) GetWebhook(subscriptionID string) (*WebhookResponse, error) {
	var response WebhookResponse
	
	// Add query parameter
	params := url.Values{}
	params.Add("subscriptionId", subscriptionID)
	
	path := fmt.Sprintf("/v1/kaia/%s/webhooks?%s", c.network, params.Encode())
	err := c.doRequestAndDecode("GET", path, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get webhook: %w", err)
	}
	return &response, nil
}

// UpdateWebhook updates an existing webhook configuration
func (c *Client) UpdateWebhook(subscriptionID string, request UpdateWebhookRequest) (*WebhookResponse, error) {
	var response WebhookResponse
	path := fmt.Sprintf("/v1/kaia/%s/webhooks/subscriptionId=%s", c.network, subscriptionID)
	err := c.doRequestAndDecode("PATCH", path, request, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update webhook: %w", err)
	}
	return &response, nil
}

// DeleteWebhook deletes a webhook configuration
func (c *Client) DeleteWebhook(subscriptionID string) (*WebhookResponse, error) {
	var response WebhookResponse
	path := fmt.Sprintf("/v1/kaia/%s/webhooks/subscriptionId=%s", c.network, subscriptionID)
	err := c.doRequestAndDecode("DELETE", path, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to delete webhook: %w", err)
	}
	return &response, nil
} 