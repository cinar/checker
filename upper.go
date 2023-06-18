package checker

import (
	"reflect"
	"strings"
)

// NormalizerUpper is the name of the normalizer.
const NormalizerUpper = "upper"

// makeUpper makes a normalizer function for the upper normalizer.
func makeUpper(_ string) CheckFunc {
	return normalizeUpper
}

// normalizeUpper maps all Unicode letters in the given value to their upper case.
func normalizeUpper(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(strings.ToUpper(value.String()))

	return ResultValid
}
