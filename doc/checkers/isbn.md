# ISBN Checker

An [ISBN (International Standard Book Number)](https://en.wikipedia.org/wiki/International_Standard_Book_Number) is a 10 or 13 digit number that is used to identify a book.

The `isbn` checker checks if the given value is a valid ISBN. If the given value is not a valid ISBN, the checker will return the `NOT_ISBN` result. Here is an example:

```golang
type Book struct {
  ISBN string `checkers:"isbn"`
}

book := &Book{
  ISBN: "1430248270",
}

_, valid := checker.Check(book)
if !valid {
  // Send the mistakes back to the user
}
```

The `isbn` checker can also be configured to check for a 10 digit or a 13 digit number. Here is an example:

```golang
type Book struct {
  ISBN string `checkers:"isbn:13"`
}

book := &Book{
  ISBN: "9781430248279",
}

_, valid := checker.Check(book)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `isbn` checker functions below to validate the user input. 

- [`IsISBN`](https://pkg.go.dev/github.com/cinar/checker#IsISBN) checks if the given value is a valid ISBN number.
- [`IsISBN10`](https://pkg.go.dev/github.com/cinar/checker#IsISBN10) checks if the given value is a valid ISBN 10 number.
- [`IsISBN13`](https://pkg.go.dev/github.com/cinar/checker#IsISBN13) checks if the given value is a valid ISBN 13 number.

Here is an example:

```golang
result := checker.IsISBN("1430248270")
if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
