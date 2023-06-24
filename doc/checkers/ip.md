# IP Checker

The `ip` checker checks if the value is an IP address. If the value is not an IP address, the checker will return the `NOT_IP` result. Here is an example:

```golang
type Request struct {
  RemoteIP string `checkers:"ip"`
}

request := &Request{
  RemoteIP: "192.168.1.1",
}

_, valid := checker.Check(request)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `ip` checker function [`IsIP`](https://pkg.go.dev/github.com/cinar/checker#IsIP) to validate the user input. Here is an example:

```golang
result := checker.IsIP("2001:db8::68")

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
