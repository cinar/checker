# URL Unescape Normalizer

The `url-unescape` normalizer uses [net.url.QueryUnescape](https://pkg.go.dev/net/url#QueryUnescape) to converte each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB.

```golang
type Request struct {
  Query string `checkers:"url-unescape"`
}

request := &Request{
  Query: "param1%2Fparam2+%3D+1+%2B+2+%26+3+%2B+4",
}

_, valid := checker.Check(request)
if !valid {
  t.Fail()
}

if request.Query != "param1/param2 = 1 + 2 & 3 + 4" {
  t.Fail()
}

// Outputs:
// param1/param2 = 1 + 2 & 3 + 4
fmt.Println(comment.Body)
```
