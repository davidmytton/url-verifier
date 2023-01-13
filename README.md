# url-verifier

🔗 A Go library for URL validation and verification: does this URL actually
work?

[![Build Status](https://github.com/davidmytton/url-verifier/actions/workflows/go.yml/badge.svg)](https://github.com/davidmytton/url-verifier/actions)
[![codecov](https://codecov.io/gh/davidmytton/url-verifier/branch/main/graph/badge.svg?token=HXSXEHU79J)](https://codecov.io/gh/davidmytton/url-verifier)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidmytton/url-verifier)](https://goreportcard.com/report/github.com/davidmytton/url-verifier)

## Features

- **URL Validation:** validates whether a string is a valid URL.
- **Different Validation Types:** validates whether the URL is valid according
  to a "human" definition of a correct URL, strict compliance with
  [RFC3986](https://www.rfc-editor.org/rfc/rfc3986) (Uniform Resource Identifier
  (URI): Generic Syntax), and/or compliance with RFC3986 with the addition of a
  schema e.g. HTTPS.
- **Reachability:** verifies whether the URL is actually reachable via an HTTP
  GET request and provides the status code returned.

## Rationale

There are several methods of validating URLs in Go depending on what you're
trying to achieve. Strict, technical validation can be done through a simple
call to [`url.Parse`](https://pkg.go.dev/net/url#Parse) in Go's Standard library
or a more "human" definition of a valid URL using
[govalidator](https://github.com/asaskevich/govalidator) (which is what this
library uses internally for syntax verification).

However, this will successfully validate all types of URLs, from relative paths
through to hostnames without a scheme. Often, when building user-facing
applications, what we actually want is a way to check whether the URL input
provided will actually work i.e. it's valid, it resolves, and it can be loaded
in a web browser.

## Install

Use `go get` to install this package.

```shell
go get -u github.com/davidmytton/url-verifier
```

## Usage

### Basic usage

Use `Verify` to check whether a URL is correct:

```go
package main

import (
 "fmt"

 urlverifier "github.com/davidmytton/url-verifier"
)

func main() {
 url := "https://example.com/"

 verifier := urlverifier.NewVerifier()
 ret, err := verifier.Verify(url)

 if err != nil {
  fmt.Errorf("Error: %s", err)
 }

 fmt.Printf("Result: %+v\n", ret)
 /*
   Result: &{
    URL:https://example.com/
    URLComponents:https://example.com/
    IsURL:true
    IsRFC3986URL:true
    IsRFC3986URI:true
    HTTP:<nil>
   }
 */
}

```

### URL reachability check

Call `EnableHTTPCheck()` to issue a `GET` request to the HTTP or HTTPS URL and
check whether it is reachable and successfully returns a response (a success
(2xx) or success-like code (3xx)). Non-HTTP(S) URLs will return an error.

```go
package main

import (
 "fmt"

 urlverifier "github.com/davidmytton/url-verifier"
)

func main() {
 url := "https://example.com/"

 verifier := urlverifier.NewVerifier()
 verifier.EnableHTTPCheck()
 ret, err := verifier.Verify(url)

 if err != nil {
  fmt.Errorf("Error: %s", err)
 }

 fmt.Printf("Result: %+v\n", ret)
 fmt.Printf("HTTP: %+v\n", ret.HTTP)

 if ret.HTTP.IsSuccess {
  fmt.Println("The URL is reachable with status code", ret.HTTP.StatusCode)
 }
 /*
   Result: &{
    URL:https://example.com/
    URLComponents:https://example.com/
    IsURL:true
    IsRFC3986URL:true
    IsRFC3986URI:true
    HTTP:0x140000b6a50
   }
   HTTP: &{
    Reachable:true
    StatusCode:200
    IsSuccess:true
   }
   The URL is reachable with status code 200
 */
}
```

## HTTP checks against internal URLs

By default, the reachability checks are only executed if the host resolves to a
non-internal IP address. An internal IP address is defined as any of:
[private](https://pkg.go.dev/net#IP.IsPrivate),
[loopback](https://pkg.go.dev/net#IP.IsLoopback), [link-local
unicast](https://pkg.go.dev/net#IP.IsLinkLocalUnicast), [link-local
multicast](https://pkg.go.dev/net#IP.IsLinkLocalMulticast), [interface-local
multicast](https://pkg.go.dev/net#IP.IsInterfaceLocalMulticast), or
[unspecified](https://pkg.go.dev/net#IP.IsUnspecified).

This is one layer of protection against [Server Side Request
Forgery](https://cheatsheetseries.owasp.org/cheatsheets/Server_Side_Request_Forgery_Prevention_Cheat_Sheet.html#application-layer_1)
(SSRF) requests.

To allow internal HTTP checks, call `verifier.AllowHTTPCheckInternal()`:

```go
urlToCheck := "http://localhost:3000"

verifier := NewVerifier()
verifier.EnableHTTPCheck()
// Danger: Makes SSRF easier!
verifier.AllowHTTPCheckInternal()
ret, err := verifier.Verify(urlToCheck)
...
```

## Credits

This library is heavily inspired by
[`email-verifier`](https://github.com/AfterShip/email-verifier).

## License

This package is licensed under the MIT License.
