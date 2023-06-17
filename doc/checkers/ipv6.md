# IPv6 Checker

The ```ipv6``` checker checks if the value is an IPv6 address. If the value is not an IPv6 address, the checker will return the ```NOT_IP_V6``` result. Here is an example:

```golang
type Request struct {
  RemoteIp string `checkers:"ipv6"`
}

request := &Request{
  RemoteIp: "2001:db8::68",
}

_, valid := Check(request)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the ```ipv6``` checker function ```IsIpV6``` to validate the user input. Here is an example:

```golang
result := IsIpV6("2001:db8::68")

if result != ResultValid {
  // Send the mistakes back to the user
}
```
