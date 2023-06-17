# Alphanumeric Checker

The ```alphanumeric``` checker checks if the given string consists of only alphanumeric characters. If the string contains non-alphanumeric characters, the checker will return the ```NOT_ALPHANUMERIC``` result. Here is an example:

```golang
type User struct {
  Username string `checkers:"alphanumeric"`
}

user := &User{
  Username: "ABcd1234",
}

_, valid := Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the ```alphanumeric``` checker function ```IsAlphanumeric``` to validate the user input. Here is an example:

```golang
result := IsAlphanumeric("ABcd1234")

if result != ResultValid {
  // Send the mistakes back to the user
}
```
