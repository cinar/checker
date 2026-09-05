// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

// TestRegistryConcurrentAccess registers custom makers, field makers, schema
// makers, and locales concurrently with the checks and schema generation
// that read those registries. Run with `go test -race`: before the
// package-level registries were guarded by a mutex, this reliably tripped
// the race detector (and could crash outright with "fatal error:
// concurrent map read and map write" outside of -race).
func TestRegistryConcurrentAccess(t *testing.T) {
	type Plain struct {
		Email string `checkers:"email"`
	}

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		i := i

		wg.Add(6)

		go func() {
			defer wg.Done()

			v2.RegisterMaker(fmt.Sprintf("race-maker-%d", i), func(_ string) v2.CheckFunc[reflect.Value] {
				return func(v reflect.Value) (reflect.Value, error) { return v, nil }
			})
		}()

		go func() {
			defer wg.Done()

			v2.RegisterFieldMaker(fmt.Sprintf("race-field-maker-%d", i), func(_ string) v2.CheckFieldFunc {
				return func(_, v reflect.Value) (reflect.Value, error) { return v, nil }
			})
		}()

		go func() {
			defer wg.Done()

			v2.RegisterSchemaMaker(fmt.Sprintf("race-schema-maker-%d", i), func(_ *v2.Schema, _ string) {})
		}()

		go func() {
			defer wg.Done()

			v2.RegisterLocale(fmt.Sprintf("race-locale-%d", i), map[string]string{
				"NOT_EMAIL": "race",
			})
		}()

		go func() {
			defer wg.Done()

			_, _ = v2.CheckStruct(&Plain{Email: "test@example.com"})
		}()

		go func() {
			defer wg.Done()

			_ = v2.JSONSchema(&Plain{})
		}()
	}

	wg.Wait()
}
