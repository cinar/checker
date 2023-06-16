# ASCII Checker

The ```ascii``` checkr checks if the given string consists of only ASCII characters. If the string contains non-ASCII characters, the checker will return the ```NOT_ASCII``` result. Here is an example:

```golang
type User struct {
  Username string `checkers:"ascii"`
}

user := &User{
  Username: "checker",
}

_, valid := Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the ```ascii``` checker function ```IsAscii``` to validate the user input. Here is an example:

```golang
result := IsAscii("Checker")

if result != ResultValid {
  // Send the mistakes back to the user
}
```
