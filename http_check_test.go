// SPDX-License-Identifier: MIT
package urlverifier

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckVerify_HTTPCheckEnabledStatus200(t *testing.T) {
	urlToCheck := "http://example.com/"

	verifier := NewVerifier()
	verifier.EnableHTTPCheck()
	ret, err := verifier.Verify(urlToCheck)

	expected := Result{
		URL:           urlToCheck,
		URLComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/"},
		IsURL:         true,
		IsRFC3986URL:  true,
		IsRFC3986URI:  true,
		HTTP: &HTTP{
			Reachable:  true,
			StatusCode: 200,
			IsSuccess:  true,
		},
	}

	assert.Equal(t, expected, *ret)
	assert.Nil(t, err)
}

func TestCheckVerify_HTTPCheckEnabledStatus404(t *testing.T) {
	urlToCheck := "http://example.com/notfound"

	verifier := NewVerifier()
	verifier.EnableHTTPCheck()
	ret, err := verifier.Verify(urlToCheck)

	expected := Result{
		URL:           urlToCheck,
		URLComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/notfound"},
		IsURL:         true,
		IsRFC3986URL:  true,
		IsRFC3986URI:  true,
		HTTP: &HTTP{
			Reachable:  true,
			StatusCode: 404,
			IsSuccess:  false,
		},
	}

	assert.Equal(t, expected, *ret)
	assert.Nil(t, err)
}

func TestCheckVerify_HTTPCheckEnabledUnreachable(t *testing.T) {
	urlToCheck := "http://example.unreachable/"

	verifier := NewVerifier()
	verifier.EnableHTTPCheck()
	ret, err := verifier.Verify(urlToCheck)

	expected := Result{
		URL:           urlToCheck,
		URLComponents: &url.URL{Scheme: "http", Host: "example.unreachable", Path: "/"},
		IsURL:         true,
		IsRFC3986URL:  true,
		IsRFC3986URI:  true,
		HTTP: &HTTP{
			Reachable: false,
			IsSuccess: false,
		},
	}

	assert.Equal(t, expected, *ret)
	assert.IsType(t, &url.Error{}, err)
	assert.ErrorContains(t, err, "dial tcp: lookup example.unreachable: no such host")
}
