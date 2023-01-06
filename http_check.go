// SPDX-License-Identifier: MIT
package urlverifier

import "net/http"

// HTTP is the result of a HTTP check
type HTTP struct {
	Reachable  bool `json:"reachable"`   // Whether the URL is reachable via HTTP. This may be true even if the response is an HTTP error e.g. a 500 error.
	StatusCode int  `json:"status_code"` // The HTTP status code
	IsSuccess  bool `josn:"is_success"`  // Whether the HTTP response is a success (2xx) or success-like code (3xx)
}

// CheckHTTP checks if the URL is reachable via HTTP
func (v *Verifier) CheckHTTP(urlToCheck string) (*HTTP, error) {
	if !v.httpCheckEnabled {
		return nil, nil
	}

	ret := HTTP{
		Reachable: false,
		IsSuccess: false,
	}

	// Check if the URL is reachable via HTTP
	resp, err := http.Get(urlToCheck)
	if err != nil {
		return &ret, err
	}
	defer resp.Body.Close()

	ret.Reachable = true
	ret.StatusCode = resp.StatusCode

	// Check if the HTTP response is a success (2xx) or success-like code (3xx)
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		ret.IsSuccess = true
	}

	return &ret, nil
}
