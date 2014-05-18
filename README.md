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

