// SPDX-License-Identifier: MIT
package urlverifier

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testURLs = []struct {
	rawURL        string
	urlComponents *url.URL
	isURL         bool
	isRFC3986URL  bool
	isRFC3986URI  bool
}{
	{rawURL: "http://example.com",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "https://example.com",
		urlComponents: &url.URL{Scheme: "https", Host: "example.com"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com/",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com/path",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/path"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com/path?query",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/path", RawQuery: "query"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com/path?query#fragment",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/path", RawQuery: "query", Fragment: "fragment"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://user:pass@www.example.com/",
		urlComponents: &url.URL{Scheme: "http", Host: "www.example.com", Path: "/", User: url.UserPassword("user", "pass")},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  false,
		isRFC3986URI:  false},
	{rawURL: "example.com",
		urlComponents: &url.URL{Scheme: "", Host: "", Path: "example.com"},
		isURL:         true,
		isRFC3986URL:  false,
		isRFC3986URI:  false},
	{rawURL: "http://example.dev/",
		urlComponents: &url.URL{Scheme: "http", Host: "example.dev", Path: "/"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.中文网/",
		urlComponents: &url.URL{Scheme: "http", Host: "example.中文网", Path: "/"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com:8080",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com:8080"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "ftp://example.com",
		urlComponents: &url.URL{Scheme: "ftp", Host: "example.com"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "ftp.example.com",
		urlComponents: &url.URL{Scheme: "", Host: "", Path: "ftp.example.com"},
		isURL:         true,
		isRFC3986URL:  false,
		isRFC3986URI:  false},
	{rawURL: "http://127.0.0.1/",
		urlComponents: &url.URL{Scheme: "http", Host: "127.0.0.1", Path: "/"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com/?query=%2F",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/", RawQuery: "query=%2F"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://localhost:3000/",
		urlComponents: &url.URL{Scheme: "http", Host: "localhost:3000", Path: "/"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com/?query",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/", RawQuery: "query"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com?query",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "", RawQuery: "query"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://www.xn--froschgrn-x9a.net/",
		urlComponents: &url.URL{Scheme: "http", Host: "www.xn--froschgrn-x9a.net", Path: "/"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com/a-",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/a-"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.پاکستان/",
		urlComponents: &url.URL{Scheme: "http", Host: "example.پاکستان", Path: "/"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.c_o_m/",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://_example.com/",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example_example.com/",
		urlComponents: &url.URL{Scheme: "http", Host: "example_example.com", Path: "/"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "xyz://example.com",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: ".com",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  false,
		isRFC3986URI:  false},
	{rawURL: "invalid.",
		urlComponents: &url.URL{Scheme: "", Host: "", Path: "invalid."},
		isURL:         true,
		isRFC3986URL:  false,
		isRFC3986URI:  false},
	{rawURL: "http://example.com/~user",
		urlComponents: &url.URL{Scheme: "http", Host: "example.com", Path: "/~user"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "mailto:someone@example.com",
		urlComponents: &url.URL{Scheme: "mailto", Host: "", Opaque: "someone@example.com"},
		isURL:         true,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "/abs/test/dir",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  false,
		isRFC3986URI:  true},
	{rawURL: "./rel/test/dir",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  false,
		isRFC3986URI:  false},
	{rawURL: "http://example-.com/",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://-example.com/",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example_.com/",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://_example.com/",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com:80:80/",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
	{rawURL: "http://example.com://8080",
		urlComponents: nil,
		isURL:         false,
		isRFC3986URL:  true,
		isRFC3986URI:  true},
}

func TestCheckVerify_HTTPCheckDisabledDefault(t *testing.T) {
	for _, test := range testURLs {
		urlToCheck := test.rawURL

		verifier := NewVerifier()
		//verifier.DisableHTTPCheck()
		ret, err := verifier.Verify(urlToCheck)

		expected := Result{
			URL:           urlToCheck,
			URLComponents: test.urlComponents,
			IsURL:         test.isURL,
			IsRFC3986URL:  test.isRFC3986URL,
			IsRFC3986URI:  test.isRFC3986URI,
			HTTP:          nil,
		}

		assert.Equal(t, expected, *ret)
		assert.Nil(t, err)
	}
}

func TestCheckVerify_HTTPCheckDisabledExplicit(t *testing.T) {
	for _, test := range testURLs {
		urlToCheck := test.rawURL

		verifier := NewVerifier()
		verifier.DisableHTTPCheck()
		ret, err := verifier.Verify(urlToCheck)

		expected := Result{
			URL:           urlToCheck,
			URLComponents: test.urlComponents,
			IsURL:         test.isURL,
			IsRFC3986URL:  test.isRFC3986URL,
			IsRFC3986URI:  test.isRFC3986URI,
			HTTP:          nil,
		}

		assert.Equal(t, expected, *ret)
		assert.Nil(t, err)
	}
}

func TestIsRequestURL(t *testing.T) {
	for _, test := range testURLs {
		urlToCheck := test.rawURL

		verifier := NewVerifier()
		ret := verifier.IsRequestURL(urlToCheck)

		assert.Equal(t, test.isRFC3986URL, ret)
	}
}

func TestIsRequestURI(t *testing.T) {
	for _, test := range testURLs {
		urlToCheck := test.rawURL

		verifier := NewVerifier()
		ret := verifier.IsRequestURI(urlToCheck)

		assert.Equal(t, test.isRFC3986URI, ret)
	}
}
