# Min Checker

The `min` checker checks if the given ```int``` or ```float``` value is greather than the given minimum. If the value is below the minimum, the checker will return the `NOT_MIN` result. Here is an example:

```golang
type User struct {
  Age int `checkers:"min:21"`
}

user := &User{
  Age: 45,
}

mistakes, valid := checker.Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `min` checker function [`IsMin`](https://pkg.go.dev/github.com/cinar/checker#IsMin) to validate the user input. Here is an example:

```golang
age := 45

result := checker.IsMin(age, 21)

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
