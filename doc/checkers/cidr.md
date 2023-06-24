# CIDR Checker

The `cidr` checker checks if the value is a valid CIDR notation IP address and prefix length, like `192.0.2.0/24` or `2001:db8::/32`, as defined in [RFC 4632](https://rfc-editor.org/rfc/rfc4632.html) and [RFC 4291](https://rfc-editor.org/rfc/rfc4291.html). If the value is not a valid CIDR notation, the checker will return the `NOT_CIDR` result. Here is an example:

```golang
type Network struct {
  Subnet string `checkers:"cidr"`
}

network := &Network{
  Subnet: "192.0.2.0/24",
}

_, valid := checker.Check(network)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `cidr` checker function [`IsCidr`](https://pkg.go.dev/github.com/cinar/checker#IsASCII) to validate the user input. Here is an example:

```golang
result := checker.IsCidr("2001:db8::/32")

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
