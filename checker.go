// Package checker is a Go library for validating user input through struct tags.
package checker

import "reflect"

// Result is a unique textual identifier for the mistake.
type Result string

// CheckerFunc defines the checker function.
type CheckerFunc func(value, parent reflect.Value) Result

// MakerFunc defines the maker function.
type MakerFunc func(config string) CheckerFunc

// Mistake provides the field where the mistake was made and a result for the mistake.
type Mistake struct {
	Field  string
	Result Result
}

// ResultValid result indicates that the user input is valid.
const ResultValid Result = "VALID"

var makers = map[string]MakerFunc{
	"required": makeRequired,
}

// Register registers the given checker name and the maker function.
func Register(name string, maker MakerFunc) {

}
