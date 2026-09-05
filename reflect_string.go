// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import "reflect"

// reflectString returns the string content of value. It supports defined
// types whose underlying kind is string (e.g. `type Email string`), not
// just the built-in string type, so checkers and normalizers don't panic
// on interface conversion when applied to a custom string type. It still
// panics if value's kind isn't string at all, since that means a
// string-only checker was applied to a field of the wrong type.
func reflectString(value reflect.Value) string {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return value.String()
}
