package tuxedo

import (
	"bytes"
	"io"
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

// execute sends an HTTP request with the given method and URL.
//
// It also applies headers, handles request bodies, and ensures proper response handling.
//
// Parameters:
//   - method: HTTP method (GET, POST, etc.).
//   - url: Target URL for the request.
//
// Returns:
//   - *Response: A structured response containing status code, body, and headers.
//   - error: Any error encountered during execution.
func (r *Request) execute(method, url string) (*Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(r.body))
	if err != nil {
		return nil, err
	}
	for key, value := range r.headers {
		req.Header.Set(key, value)
	}
	if r.body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if r.enableTrace {
		// TODO: Add trace logs
	}
	resp, err := r.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer doClose(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       respBody,
		Headers:    resp.Header,
	}, nil
}

// Get sends a GET request to the specified URL.
//
// Parameters:
//   - url: Target URL for the request.
//
// Returns:
//   - *Response: The HTTP response.
//   - error: Any error encountered.
func (r *Request) Get(url string) (*Response, error) {
	return r.execute(http.MethodGet, url)
}

// Post sends a POST request to the specified URL.
//
// Parameters:
//   - url: Target URL for the request.
//
// Returns:
//   - *Response: The HTTP response.
//   - error: Any error encountered.
func (r *Request) Post(url string) (*Response, error) {
	return r.execute(http.MethodPost, url)
}

// Put sends a PUT request to the specified URL.
//
// Parameters:
//   - url: Target URL for the request.
//
// Returns:
//   - *Response: The HTTP response.
//   - error: Any error encountered.
func (r *Request) Put(url string) (*Response, error) {
	return r.execute(http.MethodPut, url)
}

// Delete sends a DELETE request to the specified URL.
//
// Parameters:
//   - url: Target URL for the request.
//
// Returns:
//   - *Response: The HTTP response.
//   - error: Any error encountered.
func (r *Request) Delete(url string) (*Response, error) {
	return r.execute(http.MethodDelete, url)
}
