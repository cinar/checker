# URL Escape Normalizer

The `url-escape` normalizer uses [net.url.QueryEscape](https://pkg.go.dev/net/url#QueryEscape) to escape the string so it can be safely placed inside a URL query.

```golang
type Request struct {
  Query string `checkers:"url-escape"`
}

request := &Request{
  Query: "param1/param2 = 1 + 2 & 3 + 4",
}

_, valid := checker.Check(request)
if !valid {
  t.Fail()
}

// Outputs: 
// param1%2Fparam2+%3D+1+%2B+2+%26+3+%2B+4
fmt.Println(request.Query)
```
