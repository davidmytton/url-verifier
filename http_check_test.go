// SPDX-License-Identifier: MIT
package urlverifier

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckHTTP_Status200(t *testing.T) {
	urlToCheck := "http://example.com/"

	verifier := NewVerifier()
	ret, err := verifier.CheckHTTP(urlToCheck)

	expected := &HTTP{
		Reachable:  true,
		StatusCode: 200,
		IsSuccess:  true,
	}

	assert.Equal(t, expected, ret)
	assert.Nil(t, err)
}

func TestCheckHTTP_Status404(t *testing.T) {
	urlToCheck := "http://example.com/notfound"

	verifier := NewVerifier()
	ret, err := verifier.CheckHTTP(urlToCheck)

	expected := &HTTP{
		Reachable:  true,
		StatusCode: 404,
		IsSuccess:  false,
	}

	assert.Equal(t, expected, ret)
	assert.Nil(t, err)
}

func TestCheckHTTP_Unreachable(t *testing.T) {
	urlToCheck := "http://example.unreachable"

	verifier := NewVerifier()
	ret, err := verifier.CheckHTTP(urlToCheck)

	expected := &HTTP{
		Reachable: false,
		IsSuccess: false,
	}

	assert.Equal(t, expected, ret)
	assert.IsType(t, &url.Error{}, err)
	assert.ErrorContains(t, err, "lookup example.unreachable: no such host")
}

func TestCheckHTTP_expiredCert(t *testing.T) {
	urlToCheck := "https://self-signed.badssl.com/"

	verifier := NewVerifier()
	ret, err := verifier.CheckHTTP(urlToCheck)

	expected := &HTTP{
		Reachable: false,
		IsSuccess: false,
	}

	assert.Equal(t, expected, ret)
	assert.IsType(t, &url.Error{}, err)
	assert.ErrorContains(t, err, "x509:")
}

func TestCheckHTTP_expiredCert_allowSkip(t *testing.T) {
	urlToCheck := "https://self-signed.badssl.com/"

	verifier := NewVerifier()
	verifier.AllowSkipCertVerification()
	ret, err := verifier.CheckHTTP(urlToCheck)

	expected := &HTTP{
		Reachable:  true,
		StatusCode: 200,
		IsSuccess:  true,
	}

	assert.Equal(t, expected, ret)
	assert.Nil(t, err)
}
