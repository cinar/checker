# URL Checker

The ```url``` checker checks if the given value is a valid URL. If the given value is not a valid URL, the checker will return the ```NOT_URL``` result. The checker uses [ParseRequestURI](https://pkg.go.dev/net/url#ParseRequestURI) function to parse the URL, and then checks if the schema or the host are both set.

Here is an example:

```golang
type Bookmark struct {
  URL string `checkers:"url"`
}

bookmark := &Bookmark{
  URL: "https://zdo.com",
}

_, valid := checker.Check(bookmark)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the ```url``` checker function ```IsURL``` to validate the user input. Here is an example:

```golang
result := checker.IsURL("https://zdo.com")
if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
