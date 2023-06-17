# Digits Checker

The ```digits``` checker checks if the given string consists of only digit characters. If the string contains non-digit characters, the checker will return the ```NOT_DIGITS``` result. Here is an example:

```golang
type User struct {
  Id string `checkers:"digits"`
}

user := &User{
  Id: "1234",
}

_, valid := Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the ```digits``` checker function ```IsDigits``` to validate the user input. Here is an example:

```golang
result := IsDigits("1234")

if result != ResultValid {
  // Send the mistakes back to the user
}
```
