package tuxedo

// Request represents an HTTP request with custom headers, body, and tracing options.
//
// It is created using the R() method of the Client.
type Request struct {
	client      *Client
	headers     map[string]string
	body        []byte
	enableTrace bool
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
