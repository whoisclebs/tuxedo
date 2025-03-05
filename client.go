package tuxedo

import (
	"net/http"
	"time"
)

// Client represents an HTTP client with custom configurations.
//
// It provides methods to send requests, apply middlewares, and handle responses.
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new HTTP client with a specified timeout.
//
// Parameters:
//   - timeout: Duration for request timeouts.
//
// Returns:
//   - A new instance of Client.
func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
	}
}

// R initializes a new Request instance.
//
// Returns:
//   - *Request: A new request object associated with the Client.
func (c *Client) R() *Request {
	return &Request{
		client:  c,
		headers: make(map[string]string),
	}
}
