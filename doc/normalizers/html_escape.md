# HTML Escape Normalizer

The `html-escape` normalizer uses [html.EscapeString](https://pkg.go.dev/html#EscapeString) to escape special characters like "<" to become "&lt;". It escapes only five such characters: <, >, &, ' and ".

```golang
type Comment struct {
  Body string `checkers:"html-escape"`
}

comment := &Comment{
  Body: "<tag> \"Checker\" & 'Library' </tag>",
}

checker.Check(comment)

// Outputs: 
// &lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;
fmt.Println(comment.Body)
```
