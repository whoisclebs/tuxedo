package tuxedo

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// Response represents an HTTP response.
//
// It contains the status code, headers, and body of the response.
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// Json unmarshals the response body into the provided target struct.
//
// Parameters:
//   - target: A pointer to the struct where the JSON data should be unmarshaled.
//
// Returns:
//   - error: Any error encountered during unmarshaling.
func (res *Response) Json(target interface{}) error {
	return json.Unmarshal(res.Body, target)
}

// Xml unmarshals the response body into the provided target struct.
//
// Parameters:
//   - target: A pointer to the struct where the XML data should be unmarshaled.
//
// Returns:
//   - error: Any error encountered during unmarshaling.
func (res *Response) Xml(target interface{}) error {
	return xml.Unmarshal(res.Body, target)
}
