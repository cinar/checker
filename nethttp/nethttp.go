// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package nethttp integrates Checker with the standard library's net/http,
// decoding and validating a JSON request body in a single call, and writing
// a JSON 400 response automatically when decoding or validation fails.
package nethttp

import (
	"encoding/json"
	"net/http"

	checker "github.com/cinar/checker/v2"
)

// jsonContentType is the Content-Type header value written on every JSON
// response this package produces.
const jsonContentType = "application/json; charset=utf-8"

// Bind decodes r's JSON body into target, then runs Check on it. If
// decoding fails, it writes a 400 JSON response describing the decode error
// to w and returns false. Otherwise, it returns the result of Check.
func Bind(w http.ResponseWriter, r *http.Request, target any) bool {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		writeJSONError(w, err)
		return false
	}

	return Check(w, target)
}

// Check runs checker.CheckStruct on target. If validation fails, it writes
// a 400 JSON response to w, whose body is target's checker.CheckErrors
// marshaled with CheckErrors.JSON, and returns false. Otherwise, it returns
// true, leaving target normalized by any checker normalizers for the
// handler to use.
func Check(w http.ResponseWriter, target any) bool {
	errs, ok := checker.CheckStruct(target)
	if !ok {
		// json.Marshal cannot fail for FieldError's string fields.
		data, _ := errs.JSON()

		w.Header().Set("Content-Type", jsonContentType)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(data)

		return false
	}

	return true
}

// writeJSONError writes a 400 JSON response with a single "error" field
// describing err, matching the shape of the gin and echo adapters' bind
// error responses.
func writeJSONError(w http.ResponseWriter, err error) {
	// json.Marshal cannot fail for a single string field.
	data, _ := json.Marshal(map[string]string{"error": err.Error()})

	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(data)
}
