# IPv6 Checker

The `ipv6` checker checks if the value is an IPv6 address. If the value is not an IPv6 address, the checker will return the `NOT_IP_V6` result. Here is an example:

```golang
type Request struct {
  RemoteIP string `checkers:"ipv6"`
}

request := &Request{
  RemoteIP: "2001:db8::68",
}

_, valid := checker.Check(request)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `ipv6` checker function [`IsIPV6`](https://pkg.go.dev/github.com/cinar/checker#IsIPV6) to validate the user input. Here is an example:

```golang
result := checker.IsIPV6("2001:db8::68")

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
