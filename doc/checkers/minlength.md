# Min Length Checker

The ```min-length``` checker checks if the length of the given value is greather than the given minimum length. If the length of the value is below the minimum length, the checker will return the ```NOT_MIN_LENGTH``` result. Here is an example:

```golang
type User struct {
  Password string `checkers:"min-length:4"`
}

user := &User{
  Password: "1234",
}

mistakes, valid := Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

The checker can be applied to all types that are supported by the [reflect.Value.Len()](https://pkg.go.dev/reflect#Value.Len) function.

If you do not want to validate user input stored in a struct, you can individually call the ```min-length``` checker function ```IsMinLength``` to validate the user input. Here is an example:

```golang
s := "1234"

result := IsMinLength(s, 4)

if result != ResultValid {
  // Send the mistakes back to the user
}
```
