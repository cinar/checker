package checker

import (
	"reflect"
)

// CheckerDigits is the name of the checker.
const CheckerDigits = "digits"

// ResultNotDigits indicates that the given string contains non-digit characters.
const ResultNotDigits = "NOT_DIGITS"

// IsDigits checks if the given string consists of only digit characters.
func IsDigits(value string) Result {
	for i := 0; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return ResultNotDigits
		}
	}

	return ResultValid
}

// makeDigits makes a checker function for the digits checker.
func makeDigits(_ string) CheckFunc {
	return checkDigits
}

// checkDigits checks if the given string consists of only digit characters.
func checkDigits(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsDigits(value.String())
}
