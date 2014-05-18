# simplerequest

A simple wrapper for Go's http library.

Usage
-----

```go
header := map[string]string
header["User-Agent"] = "My User-Agent String Goes Here"

resp, err := simplerequest.Get("http://jeffkehler.com", header)

println(resp.Body)
```

Response
--------

Url: Original requested Url
FinalUrl: Final url after all redirects
Body: Text body of text or json responses
ContentType: Content-Type header
ContentLength: Content-Length header

