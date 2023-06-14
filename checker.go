// Package checker is a Go library for validating user input through struct tags.
package checker

import (
	"fmt"
	"reflect"
	"strings"
)

// Result is a unique textual identifier for the mistake.
type Result string

// CheckFunc defines the checker function.
type CheckFunc func(value, parent reflect.Value) Result

// MakeFunc defines the maker function.
type MakeFunc func(params string) CheckFunc

// Mistake provides the field where the mistake was made and a result for the mistake.
type Mistake struct {
	Field  string
	Result Result
}

// ResultValid result indicates that the user input is valid.
const ResultValid Result = "VALID"

// makers provdes mapping to maker function for the checkers.
var makers = map[string]MakeFunc{
	"required": makeRequired,
}

// Register registers the given checker name and the maker function.
func Register(name string, maker MakeFunc) {
	makers[name] = maker
}

// initCheckers initializes the checkers provided in the config.
func initCheckers(config string) []CheckFunc {
	fields := strings.Fields(config)
	checkers := make([]CheckFunc, len(fields))

	for i, field := range fields {
		name, params, _ := strings.Cut(field, ":")

		maker, ok := makers[name]
		if !ok {
			panic(fmt.Sprintf("checker %s is unkown", name))
		}

		checkers[i] = maker(params)
	}

	return checkers
}
