// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

// nameOmitEmpty is the checkers tag token that skips the rest of a field's
// (or, with an "@" prefix, a slice/map container's) checks when the value is
// its zero value. Unlike a checker or normalizer, it has no maker of its
// own: reflectCheckFieldWithConfig strips it out of the config string and
// acts on it directly, since it needs the value being checked to decide
// anything, not just the checker's own parameters.
const nameOmitEmpty = "omitempty"
