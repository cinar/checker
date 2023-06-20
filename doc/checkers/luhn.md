# Luhn Checker

The ```luhn``` checker checks if the given number is valid based on the Luhn Algorithm. If the given number is not valid, it will return the ```NOT_LUHN``` result. Here is an example:

```golang
type Order struct {
  CreditCard string `checkers:"luhn"`
}

order := &Order{
  CreditCard: "4012888888881881",
}

_, valid := Check(order)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the ```luhn``` checker function ```IsLuhn``` to validate the user input. Here is an example:

```golang
result := IsLuhn("4012888888881881")

if result != ResultValid {
  // Send the mistakes back to the user
}
```
