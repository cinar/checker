# HTML Escape Normalizer

The `url-escape` normalizer uses [net.url.QueryEscape](https://pkg.go.dev/net/url#QueryEscape) to escape the string so it can be safely placed inside a URL query.

```golang
type Comment struct {
  Body string `checkers:"url-escape"`
}

comment := &Comment{
  Body: "<tag> \"Checker\" & 'Library' </tag>",
}

checker.Check(comment)

// Outputs: 
// &lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;
fmt.Println(comment.Body)
```
