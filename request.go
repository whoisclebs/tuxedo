package tuxedo

import (
	"bytes"
	"io"
	"net/http"
)

// Request represents an HTTP request with custom headers, body, and tracing options.
//
// It is created using the R() method of the Client.
type Request struct {
	client      *Client
	headers     map[string]string
	body        []byte
	enableTrace bool
}

// AddHeader sets a custom header for the request.
//
// Parameters:
//   - key: The header name.
//   - value: The header value.
//
// Returns:
//   - *Request: The request instance to allow method chaining.
func (r *Request) AddHeader(key, value string) *Request {
	r.headers[key] = value
	return r
}

// SetBody sets the request body.
//
// Parameters:
//   - body: The request body as a byte slice.
//
// Returns:
//   - *Request: The request instance to allow method chaining.
func (r *Request) SetBody(body []byte) *Request {
	r.body = body
	return r
}

// EnableTrace enables tracing for the request.
//
// When enabled, it allows collecting debug information (to be implemented).
//
// Returns:
//   - *Request: The request instance to allow method chaining.
func (r *Request) EnableTrace() *Request {
	r.enableTrace = true
	return r
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
