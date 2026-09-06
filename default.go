// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strconv"
)

const (
	// nameDefault is the name of the default normalizer.
	nameDefault = "default"
)

// Default returns a normalizer that replaces value with fallback when value
// is its zero value, and leaves value untouched otherwise. Unlike a checker,
// it never returns an error.
func Default[T any](fallback T) CheckFunc[T] {
	return func(value T) (T, error) {
		if reflect.ValueOf(value).IsZero() {
			return fallback, nil
		}

		return value, nil
	}
}

// makeDefault creates the default normalizer function from a string
// parameter. Unlike every other maker, the parameter's meaning depends on
// the kind of the field it ends up attached to (a string, a bool, or a
// number), which isn't known until the returned function actually runs
// against a real reflect.Value -- so, unlike a checker with a fixed
// parameter type (e.g. gt's numeric threshold), a malformed default value
// only panics the first time the field is actually zero, not immediately.
func makeDefault(params string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		// A nil pointer field reaches here as-is (Kind Pointer, IsNil true)
		// only via CheckStruct's indirectOrNilPointer, which preserves it
		// specifically so checkers can act on "missing" pointer fields; a
		// non-nil pointer is always already dereferenced to its pointee by
		// the time any checker runs, through every entry point, so this
		// branch is never reached for one.
		if value.Kind() == reflect.Pointer {
			ptr := reflect.New(value.Type().Elem())
			ptr.Elem().Set(parseDefaultValue(value.Type().Elem(), params))

			return ptr, nil
		}

		if !value.IsZero() {
			return value, nil
		}

		return parseDefaultValue(value.Type(), params), nil
	}
}

// parseDefaultValue parses params as a value of t's kind, converted to t
// itself so a defined type (e.g. `type Status string`) round-trips
// correctly. Panics if params can't be parsed as t's kind, or t's kind
// isn't supported (only string, bool, and the int/uint/float kinds are).
func parseDefaultValue(t reflect.Type, params string) reflect.Value {
	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf(params).Convert(t)

	case reflect.Bool:
		b, err := strconv.ParseBool(params)
		if err != nil {
			panic("unable to parse params as bool")
		}

		return reflect.ValueOf(b).Convert(t)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(params, 10, 64)
		if err != nil {
			panic("unable to parse params as int")
		}

		return reflect.ValueOf(n).Convert(t)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(params, 10, 64)
		if err != nil {
			panic("unable to parse params as uint")
		}

		return reflect.ValueOf(n).Convert(t)

	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(params, 64)
		if err != nil {
			panic("unable to parse params as float")
		}

		return reflect.ValueOf(n).Convert(t)

	default:
		panic("default is not supported for kind " + t.Kind().String())
	}
}
