// Copyright (c) 2024-2026 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
//
// Try this on Go Playground: https://go.dev/play/p/5X8ukfSOnZ1

package main

import (
	"fmt"

	checker "github.com/cinar/checker/v2"
)

type UserRegistration struct {
	// Normalization happens first (trim whitespace, lowercase), followed by validation
	Email           string   `json:"email" checkers:"trim lower required email"`
	Password        string   `json:"password" checkers:"required min-len:8"`
	ConfirmPassword string   `json:"confirm_password" checkers:"required eq-field:Password"`
	Age             int      `json:"age" checkers:"gte:18"`
	Roles           []string `json:"roles" checkers:"@max-len:3 trim alphanumeric"`
}

func main() {
	fmt.Println("=== 1. Valid Input with Normalization ===")
	validUser := &UserRegistration{
		Email:           "  GOPHER@EXAMPLE.COM  ",
		Password:        "secret123",
		ConfirmPassword: "secret123",
		Age:             25,
		Roles:           []string{"  admin ", "editor "},
	}

	errs, ok := checker.CheckStruct(validUser)
	if !ok {
		panic("expected valid user")
	}

	fmt.Printf("Normalized Email: %q\n", validUser.Email)
	fmt.Printf("Normalized Roles: %v\n", validUser.Roles)

	fmt.Println("\n=== 2. Invalid Input with Structured JSON Errors ===")
	invalidUser := &UserRegistration{
		Email:           "invalid-email",
		Password:        "secret123",
		ConfirmPassword: "mismatched-password",
		Age:             16,
		Roles:           []string{"admin", "guest", "viewer", "extra"},
	}

	errs, ok = checker.CheckStruct(invalidUser)
	if !ok {
		jsonErr, _ := errs.JSON()
		fmt.Println("Validation Errors (JSON):")
		fmt.Println(string(jsonErr))
	}
}
