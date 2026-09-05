// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"sync"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsRegexp() {
	_, err := v2.IsRegexp("^[0-9a-fA-F]+$", "ABcd1234")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsRegexpInvalid(t *testing.T) {
	_, err := v2.IsRegexp("^[0-9a-fA-F]+$", "Onur")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsRegexpValid(t *testing.T) {
	_, err := v2.IsRegexp("^[0-9a-fA-F]+$", "ABcd1234")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckRegexpNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type User struct {
		Username int `checkers:"regexp:^[A-Za-z]$"`
	}

	user := &User{}

	v2.CheckStruct(user)
}

func TestCheckRegexpInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"regexp:^[A-Za-z]+$"`
	}

	user := &User{
		Username: "abcd1234",
	}

	_, ok := v2.CheckStruct(user)
	if ok {
		t.Fatal("expected error")
	}
}

// TestIsRegexpConcurrentSamePattern exercises the compiled-pattern cache
// from many goroutines using the same expression, some new to the process
// and some already cached. Run with `go test -race`.
func TestIsRegexpConcurrentSamePattern(t *testing.T) {
	const expression = "^[a-z]+$"

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if _, err := v2.IsRegexp(expression, "onur"); err != nil {
				t.Error(err)
			}
		}()
	}

	wg.Wait()
}

func BenchmarkIsRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = v2.IsRegexp("^[0-9a-fA-F]+$", "ABcd1234")
	}
}

func TestCheckRegexpValid(t *testing.T) {
	type User struct {
		Username string `checkers:"regexp:^[A-Za-z]+$"`
	}

	user := &User{
		Username: "abcd",
	}

	_, ok := v2.CheckStruct(user)
	if !ok {
		t.Fatal("expected valid")
	}
}
