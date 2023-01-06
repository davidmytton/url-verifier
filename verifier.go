// Package urlverifier is a Go library for URL validation and verification: does
// this URL actually work?
// SPDX-License-Identifier: MIT
package urlverifier

import (
	"errors"
	"net/url"

	"github.com/asaskevich/govalidator"
)

// Verifier is a URL Verifier. Create one using NewVerifier()
type Verifier struct {
	httpCheckEnabled bool // Whether to check if the URL is reachable via HTTP (default: false)
}

// Result is the result of a URL verification
type Result struct {
	URL           string   `json:"url"`            // The URL that was checked
	URLComponents *url.URL `json:"url_components"` // The URL components, if the URL is valid
	IsURL         bool     `json:"is_url"`         // Whether the URL is valid
	IsRFC3986URL  bool     `json:"is_rfc3986_url"` // Whether the URL is a valid URL according to RFC 3986. This is the same as IsRFC3986URI but with a check for a scheme.
	IsRFC3986URI  bool     `json:"is_rfc3986_uri"` // Whether the URL is a valid URI according to RFC 3986
	HTTP          *HTTP    `json:"http"`           // The result of a HTTP check, if enabled
}

// NewVerifier creates a new URL Verifier
func NewVerifier() *Verifier {
	return &Verifier{}
}

// Verify verifies a URL. It checks if the URL is valid, parses it if so, and
// checks if it is valid according to RFC 3986 (as a URI without a scheme and a
// URL with a scheme). If the HTTP check is enabled, it also checks if the URL
// is reachable via HTTP.
func (v *Verifier) Verify(rawURL string) (*Result, error) {
	ret := Result{
		URL:          rawURL,
		IsURL:        false,
		IsRFC3986URL: false,
		IsRFC3986URI: false,
	}

	// Check if the URL is valid
	ret.IsURL = govalidator.IsURL(ret.URL)

	// If the URL is valid, parse it
	if ret.IsURL {
		p, err := url.Parse(ret.URL)
		if err != nil {
			return &ret, err
		}
		ret.URLComponents = p
	}

	// Check if the URL is a valid URI according to RFC 3986, plus a check for a
	// scheme.
	ret.IsRFC3986URL = v.IsRequestURL(ret.URL)

	// Check if the URL is a valid URI according to RFC 3986
	ret.IsRFC3986URI = v.IsRequestURI(ret.URL)

	// Check if the URL is reachable via HTTP
	if v.httpCheckEnabled {
		if ret.URLComponents.Scheme == "http" || ret.URLComponents.Scheme == "https" {
			http, err := v.CheckHTTP(ret.URL)
			if err != nil {
				ret.HTTP = http
				return &ret, err
			}
			ret.HTTP = http
		} else {
			return &ret, errors.New("unable to check if the URL is reachable via HTTP: the URL does not have a HTTP or HTTPS scheme")
		}
	}

	return &ret, nil
}

// IsRequestURL checks if the string rawURL, assuming it was received in an HTTP
// request, is a valid URL confirm to RFC 3986. Implemented from govalidator:
// https://github.com/asaskevich/govalidator/blob/f21760c49a8d602d863493de796926d2a5c1138d/validator.go#L130
func (v *Verifier) IsRequestURL(rawURL string) bool {
	url, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false // Couldn't even parse the rawURL
	}
	if len(url.Scheme) == 0 {
		return false // No Scheme found
	}
	return true
}

// IsRequestURI checks if the string rawURL, assuming it was received in an HTTP
// request, is an absolute URI or an absolute path. Implemented from
// govalidator:
// https://github.com/asaskevich/govalidator/blob/f21760c49a8d602d863493de796926d2a5c1138d/validator.go#L144
func (v *Verifier) IsRequestURI(rawURL string) bool {
	_, err := url.ParseRequestURI(rawURL)
	return err == nil
}

// DisableHTTPCheck disables checking if the URL is reachable via HTTP
func (v *Verifier) DisableHTTPCheck() {
	v.httpCheckEnabled = false
}

// EnableHTTPCheck enables checking if the URL is reachable via HTTP
func (v *Verifier) EnableHTTPCheck() {
	v.httpCheckEnabled = true
}
