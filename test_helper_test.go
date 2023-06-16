package checker

import "testing"

func TestFailIfNoPanicValid(t *testing.T) {
	defer FailIfNoPanic(t)
	panic("")
}

func TestFailIfNoPanicInvalid(t *testing.T) {
	defer FailIfNoPanic(t)
	defer FailIfNoPanic(nil)
}
