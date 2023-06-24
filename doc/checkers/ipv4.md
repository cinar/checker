# IPv4 Checker

The `ipv4` checker checks if the value is an IPv4 address. If the value is not an IPv4 address, the checker will return the `NOT_IP_V4` result. Here is an example:

```golang
type Request struct {
  RemoteIP string `checkers:"ipv4"`
}

request := &Request{
  RemoteIP: "192.168.1.1",
}

_, valid := checker.Check(request)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `ipv4` checker function [`IsIPV4`](https://pkg.go.dev/github.com/cinar/checker#IsIPV4) to validate the user input. Here is an example:

```golang
result := checker.IsIPV4("192.168.1.1")

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
