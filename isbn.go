package checker

import (
	"reflect"
	"strings"
)

// Program to check for ISBN
// https://www.geeksforgeeks.org/program-check-isbn/

// How to Verify an ISBN
// https://www.instructables.com/How-to-verify-a-ISBN/

// CheckerISBN is the name of the checker.
const CheckerISBN = "isbn"

// ResultNotISBN indicates that the given value is not a valid ISBN.
const ResultNotISBN = "NOT_ISBN"

// IsISBN10 checks if the given value is a valid ISBN-10 number.
func IsISBN10(value string) Result {
	value = strings.ReplaceAll(value, "-", "")

	if len(value) != 10 {
		return ResultNotISBN
	}

	digits := []rune(value)
	sum := 0

	for i, e := 0, len(digits); i < e; i++ {
		n := isbnDigitToInt(digits[i])
		sum += n * (e - i)
	}

	if sum%11 != 0 {
		return ResultNotISBN
	}

	return ResultValid
}

// IsISBN13 checks if the given value is a valid ISBN-13 number.
func IsISBN13(value string) Result {
	value = strings.ReplaceAll(value, "-", "")

	if len(value) != 13 {
		return ResultNotISBN
	}

	digits := []rune(value)
	sum := 0

	for i, d := range digits {
		n := isbnDigitToInt(d)
		if i%2 != 0 {
			n *= 3
		}

		sum += n
	}

	if sum%10 != 0 {
		return ResultNotISBN
	}

	return ResultValid
}

// IsISBN checks if the given value is a valid ISBN number.
func IsISBN(value string) Result {
	value = strings.ReplaceAll(value, "-", "")

	if len(value) == 10 {
		return IsISBN10(value)
	} else if len(value) == 13 {
		return IsISBN13(value)
	}

	return ResultNotISBN
}

// isbnDigitToInt returns the integer value of given ISBN digit.
func isbnDigitToInt(r rune) int {
	if r == 'X' {
		return 10
	}

	return int(r - '0')
}

// makeISBN makes a checker function for the URL checker.
func makeISBN(config string) CheckFunc {
	if config != "" && config != "10" && config != "13" {
		panic("invalid format")
	}

	return func(value, parent reflect.Value) Result {
		return checkISBN(value, parent, config)
	}
}

// checkISBN checks if the given value is a valid ISBN number.
func checkISBN(value, _ reflect.Value, mode string) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	number := value.String()

	switch mode {
	case "10":
		return IsISBN10(number)

	case "13":
		return IsISBN13(number)

	default:
		return IsISBN(number)
	}
}
