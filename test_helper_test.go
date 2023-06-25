//
package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestFailIfNoPanicValid(t *testing.T) {
	defer checker.FailIfNoPanic(t)
	panic("")
}

func TestFailIfNoPanicInvalid(t *testing.T) {
	defer checker.FailIfNoPanic(t)
	defer checker.FailIfNoPanic(nil)
}
