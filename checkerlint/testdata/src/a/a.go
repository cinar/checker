// Package a is a checkerlint analysistest fixture.
package a

import (
	"strings"
	"time"

	checker "github.com/cinar/checker/v2"
)

func init() {
	checker.RegisterMaker("my-custom-checker", nil)

	// Each of these is a call that recordCustomRegistration must recognize
	// as *not* a checker name registration, and skip without panicking:
	println("not a selector call")      // call.Fun is an *ast.Ident, not a selector
	time.Now()                          // a selector call with zero args
	strings.TrimSpace("not registered") // right shape, wrong package
	checker.OtherFunc("not registered") // right package, wrong function name

	notALiteral := "still not registered"
	checker.RegisterMaker(notALiteral, nil) // first arg isn't a string literal
}

// Address is embedded (by value and by pointer, in separate structs below)
// to exercise checkStruct's promoted-field-name approximation.
type Address struct {
	Street string
}

// Valid has no issues: every checker name is known, eq-field/required-if
// targets are real siblings, numeric/string checkers are on compatible
// fields, and the slice/array's container/item split is used correctly.
type Valid struct {
	Password        string            `checkers:"trim required"`
	ConfirmPassword string            `checkers:"required eq-field:Password"`
	Age             int               `checkers:"gte:18"`
	Tags            []string          `checkers:"@max-len:5 trim alphanumeric"`
	Fixed           [3]string         `checkers:"@max-len:1 trim"`
	Notes           map[string]string `checkers:"@max-len:5 trim"`
	Website         *string           `checkers:"email"`
	Country         string            `checkers:"required"`
	State           string            `checkers:"required-if:Country:US"`
	Custom          string            `checkers:"my-custom-checker"`
	Legacy          string            `validate:"trim required"`
	Optional        string            `checkers:"omitempty email"`
	Ignored         string            `json:"ignored"`
	Untouched       bool
	Embedded        struct{ X string } `checkers:"required"`
}

// WithEmbedded exercises the anonymous-field name approximation used by
// checkStruct when resolving eq-field/required-if/required-unless targets:
// a plain embedded identifier, an embedded pointer, and an embedded,
// package-qualified type.
type WithEmbedded struct {
	Address
	*checker.Named
	CopyOfAddress string `checkers:"eq-field:Address"`
	CopyOfNamed   string `checkers:"eq-field:Named"`
}

// Invalid exercises every diagnostic checkerlint reports.
type Invalid struct {
	Name             string  `checkers:"totallyBogus"` // want `checkerlint: unknown checker "totallyBogus"`
	Age              int     `checkers:"email"`        // want `checkerlint: email requires a string, but the field's type is int`
	Password         string  `checkers:"required"`
	ConfirmPassword  string  `checkers:"eq-field:Passwrd"`        // want `checkerlint: eq-field references field "Passwrd", which doesn't exist on this struct`
	Score            string  `checkers:"gte:1"`                   // want `checkerlint: gte requires a numeric type, but the field's type is string`
	Country          string  `checkers:"required-if:Countryy:US"` // want `checkerlint: required-if references field "Countryy", which doesn't exist on this struct`
	NotABasicChecker Address `checkers:"email"`
}
