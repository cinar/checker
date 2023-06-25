# Required Checker

The `required` checker checks for the presence of required input. If the input is not present, the checker will return the `REQUIRED` result. Here is an example:

```golang
type Person struct {
    Name string `checkers:"required"`
}

person := &Person{}

mistakes, valid := checker.Check(person)
if !valid {
    // Send the mistakes back to the user
}
```

If you do not want to validate user input stored in a struct, you can individually call the `required` checker function [`IsRequired`](https://pkg.go.dev/github.com/cinar/checker#IsRequired) to validate the user input. Here is an example:

```golang
var name string

result := checker.IsRequired(name)
if result != checker.ResultValid {
    // Send the result back to the user
}
```
