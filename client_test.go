package tuxedo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type testResponse struct {
	Method string      `json:"method"`
	URL    string      `json:"url"`
	Body   interface{} `json:"body,omitempty"`
}

var ts *httptest.Server

func newTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test-Header", "test-value")
		bodyBytes, _ := io.ReadAll(r.Body)
		var reqBody interface{}
		_ = json.Unmarshal(bodyBytes, &reqBody)
		resp := testResponse{
			Method: r.Method,
			URL:    r.URL.String(),
			Body:   reqBody,
		}
		json.NewEncoder(w).Encode(resp)
	}))
}

func TestGet(t *testing.T) {
	ts = newTestServer()
	client := NewClient(5 * time.Second)
	res, err := client.R().
		EnableTrace().
		SetHeader("Authorization", "Bearer token").
		Get(ts.URL)
	assertNoError(t, err, "Error during GET request")
	assetEqual(t, http.StatusOK, res.StatusCode, "Unexpected status code for POST")
	assetEqual(t, res.Headers.Get("X-Test-Header"), "test-value", "Unexpected value for X-Test-Header")
}
