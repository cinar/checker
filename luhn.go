// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"reflect"
)

// CheckerLuhn is the name of the checker.
const CheckerLuhn = "luhn"

// ResultNotLuhn indicates that the given number is not valid based on the Luhn algorithm.
const ResultNotLuhn = "NOT_LUHN"

// doubleTable is the values for the last digits of doubled digits added.
var doubleTable = [10]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// IsLuhn checks if the given number is valid based on the Luhn algorithm.
func IsLuhn(number string) Result {
	digits := number[:len(number)-1]
	check := rune(number[len(number)-1])

	if calculateLuhnCheckDigit(digits) != check {
		return ResultNotLuhn
	}

	return ResultValid
}

// makeLuhn makes a checker function for the Luhn algorithm.
func makeLuhn(_ string) CheckFunc {
	return checkLuhn
}

// checkLuhn checks if the given number is valid based on the Luhn algorithm.
func checkLuhn(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsLuhn(value.String())
}

// Calculates the Luhn algorighm check digit for the given number.
func calculateLuhnCheckDigit(number string) rune {
	digits := []rune(number)
	check := 0

	for i, j := 0, len(digits)-1; i <= j; i++ {
		d := int(digits[j-i] - '0')
		if i%2 == 0 {
			d = doubleTable[d]
		}

		check += d
	}

	check *= 9
	check %= 10

	return rune('0' + check)
}
