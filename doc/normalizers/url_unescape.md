# HTML Unescape Normalizer

The `url-unescape` normalizer uses [net.url.QueryUnescape](https://pkg.go.dev/net/url#QueryUnescape) to converte each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB.

```golang
type Comment struct {
  Body string `checkers:"url-unescape"`
}

comment := &Comment{
  Body: "&lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;",
}

_, valid := checker.Check(comment)
if !valid {
  t.Fail()
}

// Outputs:
// <tag> \"Checker\" & 'Library' </tag>
fmt.Println(comment.Body)
```
