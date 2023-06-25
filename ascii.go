package checker

import (
	"reflect"
	"unicode"
)

// CheckerASCII is the name of the checker.
const CheckerASCII = "ascii"

// ResultNotASCII indicates that the given string contains non-ASCII characters.
const ResultNotASCII = "NOT_ASCII"

// IsASCII checks if the given string consists of only ASCII characters.
func IsASCII(value string) Result {
	for _, c := range value {
		if c > unicode.MaxASCII {
			return ResultNotASCII
		}
	}

	return ResultValid
}

// makeASCII makes a checker function for the ASCII checker.
func makeASCII(_ string) CheckFunc {
	return checkASCII
}

// checkASCII checks if the given string consists of only ASCII characters.
func checkASCII(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsASCII(value.String())
}
