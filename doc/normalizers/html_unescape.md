# HTML Unescape Normalizer

The `html-unescape` normalizer uses [html.UnescapeString](https://pkg.go.dev/html#UnescapeString) to unescape entities like "&lt;" to become "<". It unescapes a larger range of entities than EscapeString escapes. For example, "&aacute;" unescapes to "á", as does "&#225;" and "&#xE1;".

```golang
type Comment struct {
  Body string `checkers:"html-unescape"`
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
