//go:build !test
// +build !test

package checker

import "testing"

// FailIfNoPanic fails if test didn't panic. Use this function with the defer.
func FailIfNoPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Fail()
	}
}
