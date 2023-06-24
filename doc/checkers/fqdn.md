# FQDN Checker

The Full Qualified Domain Name (FQDN) is the complete domain name for a computer or host on the internet. The `fqdn` checker checks if the given string consists of a FQDN. If the string is not a valid FQDN, the checker will return the `NOT_FQDN` result. Here is an example:

```golang
type Request struct {
  Domain string `checkers:"fqdn"`
}

request := &Request{
  Domain: "zdo.com",
}

_, valid := checker.Check(request)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `fqdn` checker function [`IsFqdn`](https://pkg.go.dev/github.com/cinar/checker#IsFqdn) to validate the user input. Here is an example:

```golang
result := checker.IsFqdn("zdo.com")

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
