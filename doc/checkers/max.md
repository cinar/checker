# Max Checker

The `max` checker checks if the given ```int``` or ```float``` value is less than the given maximum. If the value is above the maximum, the checker will return the `NOT_MAX` result. Here is an example:

```golang
type Order struct {
  Quantity int `checkers:"max:10"`
}

order := &Order{
  Quantity: 5,
}

mistakes, valid := checker.Check(order)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `max` checker function [`IsMax`](https://pkg.go.dev/github.com/cinar/checker#IsMax) to validate the user input. Here is an example:

```golang
quantity := 5

result := checker.IsMax(quantity, 10)

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
