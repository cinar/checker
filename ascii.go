package checker

import (
	"reflect"
	"unicode"
)

// CheckerAscii is the name of the checker.
const CheckerAscii = "ascii"

// ResultNotAscii indicates that the given string contains non-ASCII characters.
const ResultNotAscii = "NOT_ASCII"

// IsAscii checks if the given string consists of only ASCII characters.
func IsAscii(value string) Result {
	for i := 0; i < len(value); i++ {
		if value[i] > unicode.MaxASCII {
			return ResultNotAscii
		}
	}

	return ResultValid
}

// makeAscii makes a checker function for the ascii checker.
func makeAscii(_ string) CheckFunc {
	return checkAscii
}

// checkAscii checks if the given string consists of only ASCII characters.
func checkAscii(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsAscii(value.String())
}
