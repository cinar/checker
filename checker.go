// Package checker is a Go library for validating user input through struct tags.
package checker

// Result is a unique textual identifier for the mistake.
type Result string

// ResultValid result indicates that the user input is valid.
const ResultValid Result = "VALID"

// Mistake provides the field where the mistake was made and a result for the mistake.
type Mistake struct {
	Field  string
	Result Result
}
