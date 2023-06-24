# ASCII Checker

The `ascii` checker checks if the given string consists of only ASCII characters. If the string contains non-ASCII characters, the checker will return the `NOT_ASCII` result. Here is an example:

```golang
type User struct {
  Username string `checkers:"ascii"`
}

user := &User{
  Username: "checker",
}

_, valid := checker.Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `ascii` checker function [`IsASCII`](https://pkg.go.dev/github.com/cinar/checker#IsASCII) to validate the user input. Here is an example:

```golang
result := checker.IsASCII("Checker")

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
