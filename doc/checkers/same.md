# Same Checker

The `same` checker checks if the given value is equal to the value of the other field specified by its name. If they are not equal, the checker will return the `NOT_SAME` result. In the example below, the `same` checker ensures that the value in the `Confirm` field matches the value in the `Password` field.

```golang
type User struct {
    Password string
    Confirm  string `checkers:"same:Password"`
}

user := &User{
    Password: "1234",
    Confirm:  "1234",
}

mistakes, valid := checker.Check(user)
if !valid {
    // Send the mistakes back to the user
}
```