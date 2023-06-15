package checker

import (
	"reflect"
)

// ResultNotSame indicates that the given two values are not equal to each other.
const ResultNotSame = "NOT_SAME"

// makeSame makes a checker function for the same checker.
func makeSame(config string) CheckFunc {
	return func(value, parent reflect.Value) Result {
		return checkSame(value, parent, config)
	}
}

// checkSame checks if the given value is equal to the value of the field with the given name.
func checkSame(value, parent reflect.Value, name string) Result {
	other := parent.FieldByName(name)

	if !other.IsValid() {
		panic("other field not found")
	}

	other = reflect.Indirect(other)

	if !value.Equal(other) {
		return ResultNotSame
	}

	return ResultValid
}
