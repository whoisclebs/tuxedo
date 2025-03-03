package tuxedo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
	}
}

type Request struct {
	client      *Client
	headers     map[string]string
	body        []byte
	enableTrace bool
}

func (c *Client) R() *Request {
	return &Request{
		client:  c,
		headers: make(map[string]string),
	}
}

func (r *Request) SetHeader(key, value string) *Request {
	r.headers[key] = value
	return r
}

func (r *Request) SetBody(body []byte) *Request {
	r.body = body
	return r
}

func (r *Request) EnableTrace() *Request {
	r.enableTrace = true
	return r
}

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

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

func (r *Request) Get(url string) (*Response, error) {
	return r.execute(http.MethodGet, url)
}

func (r *Request) Post(url string) (*Response, error) {
	return r.execute(http.MethodPost, url)
}

func (r *Request) Put(url string) (*Response, error) {
	return r.execute(http.MethodPut, url)
}

func (r *Request) Delete(url string) (*Response, error) {
	return r.execute(http.MethodDelete, url)
}

func (res *Response) Unmarshal(target interface{}) error {
	return json.Unmarshal(res.Body, target)
}
