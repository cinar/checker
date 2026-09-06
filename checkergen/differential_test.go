// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checkergen_test

import (
	"testing"

	checker "github.com/cinar/checker/v2"
	"github.com/cinar/checker/v2/checkergen/testdata/fixture"
)

// assertMatches is a differential check: for the same starting value, it
// runs both the reflect-based checker.CheckStruct and generated (a
// generated Check<Type> function for the same type) against separate
// copies, and asserts they report errors under exactly the same field keys
// and leave the value normalized identically. This is checked instead of
// hand-verifying generated source, since it directly proves the generated
// code is behaviorally faithful to CheckStruct for everything it covers.
func assertMatches[T comparable](t *testing.T, generated func(*T) (checker.CheckErrors, bool), value T) {
	t.Helper()

	reflectValue := value
	generatedValue := value

	reflectErrs, reflectOK := checker.CheckStruct(&reflectValue)
	generatedErrs, generatedOK := generated(&generatedValue)

	if reflectOK != generatedOK {
		t.Fatalf("ok mismatch: CheckStruct=%v generated=%v", reflectOK, generatedOK)
	}

	if len(reflectErrs) != len(generatedErrs) {
		t.Fatalf("error count mismatch: CheckStruct=%v generated=%v", reflectErrs, generatedErrs)
	}

	for key := range reflectErrs {
		if _, ok := generatedErrs[key]; !ok {
			t.Fatalf("generated code missing error for %q, CheckStruct reported %v", key, reflectErrs[key])
		}
	}

	if reflectValue != generatedValue {
		t.Fatalf("normalized value mismatch:\nCheckStruct: %+v\ngenerated:   %+v", reflectValue, generatedValue)
	}
}

func TestGeneratedMatchesCheckStructSignupRequest(t *testing.T) {
	tests := []struct {
		name string
		req  fixture.SignupRequest
	}{
		{
			name: "all valid",
			req: fixture.SignupRequest{
				Email:           "  ALICE@EXAMPLE.COM  ",
				Password:        "supersecret",
				ConfirmPassword: "supersecret",
				Age:             18,
				Nickname:        "ali",
			},
		},
		{
			name: "all valid, omitempty nickname absent",
			req: fixture.SignupRequest{
				Email:           "alice@example.com",
				Password:        "supersecret",
				ConfirmPassword: "supersecret",
				Age:             30,
			},
		},
		{
			name: "invalid email, short password, mismatched confirm, underage, short nickname",
			req: fixture.SignupRequest{
				Email:           "not-an-email",
				Password:        "short",
				ConfirmPassword: "different",
				Age:             10,
				Nickname:        "a",
			},
		},
		{
			name: "password too short only",
			req: fixture.SignupRequest{
				Email:           "alice@example.com",
				Password:        "short",
				ConfirmPassword: "short",
				Age:             18,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertMatches(t, fixture.CheckSignupRequest, test.req)
		})
	}
}

func TestGeneratedMatchesCheckStructCoverage(t *testing.T) {
	valid := fixture.Coverage{
		Handle:    "@onur",
		Code:      "ABC",
		Sum:       "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		Zip:       "94103",
		Status:    "active",
		Excluded:  "allowed",
		Role:      "admin",
		Weight:    10,
		Ratio:     4,
		Greeting:  "",
		Threshold: 0,
		Country:   "US",
		State:     "CA",
		Type:      "member",
		Email:     "onur@example.com",
		ReturnAt:  "2024-06-01",
		DepartAt:  "2024-01-01",
		BornAt:    "2000-01-01",
		DiesAt:    "2090-01-01",
	}

	tests := []struct {
		name string
		cov  fixture.Coverage
	}{
		{name: "all valid", cov: valid},
		{
			name: "contains/starts-with/ends-with/regexp/hash/postal-code/eq/ne/oneof failures",
			cov: func() fixture.Coverage {
				c := valid
				c.Handle = "nope"
				c.Code = "abcd"
				c.Sum = "not-a-hash"
				c.Zip = "not-a-zip"
				c.Status = "inactive"
				c.Excluded = "banned"
				c.Role = "superuser"
				return c
			}(),
		},
		{
			name: "multiple-of/finite/int failures",
			cov: func() fixture.Coverage {
				c := valid
				c.Weight = 11
				c.Ratio = 3.5
				return c
			}(),
		},
		{
			name: "default fills the zero value",
			cov: func() fixture.Coverage {
				c := valid
				c.Greeting = ""
				c.Threshold = 0
				return c
			}(),
		},
		{
			name: "required-if/required-unless/before-field/after-field failures",
			cov: func() fixture.Coverage {
				c := valid
				c.State = ""
				c.Email = ""
				c.DepartAt = "2024-12-01"
				c.DiesAt = "1999-01-01"
				return c
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertMatches(t, fixture.CheckCoverage, test.cov)
		})
	}
}
