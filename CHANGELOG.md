# Change log

## 0.2.0 (2023-01-06)

- Limit HTTP reachability checks to only execute against hosts with HTTP or
  HTTPS schemas.
- Check the IPs the host resolves to and prevent executing reachability checks
  againsts internal IPs. This provides a layer of protection agains SSRF
  attacks, but can be disabled with `verifier.AllowHTTPCheckInternal()`.

## v0.1.1 (2023-01-06)

- Fixed module path declaration.

## v0.1.0 (2023-01-06)

- Initial version.
