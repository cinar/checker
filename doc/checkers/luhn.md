# Luhn Checker

The `luhn` checker checks if the given number is valid based on the Luhn Algorithm. If the given number is not valid, it will return the `NOT_LUHN` result. Here is an example:

```golang
type Order struct {
  CreditCard string `checkers:"luhn"`
}

order := &Order{
  CreditCard: "4012888888881881",
}

_, valid := checker.Check(order)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `luhn` checker function [`IsLuhn`](https://pkg.go.dev/github.com/cinar/checker#IsLuhn) to validate the user input. Here is an example:

```golang
result := checker.IsLuhn("4012888888881881")

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
