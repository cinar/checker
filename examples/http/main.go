// Copyright (c) 2024-2026 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
//
// Try this on Go Playground: https://go.dev/play/p/M_tKEKwL38G

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	checker "github.com/cinar/checker/v2"
)

type RegistrationRequest struct {
	Email           string `json:"email" checkers:"trim lower required email"`
	Password        string `json:"password" checkers:"required min-len:8"`
	ConfirmPassword string `json:"confirm_password" checkers:"required eq-field:Password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	errs, valid := checker.CheckStruct(&req)
	if !valid {
		data, _ := errs.JSON()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(data)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
		"email":   req.Email,
	})
}

func main() {
	// 1. Simulate POST /register with unnormalized, valid input:
	validPayload := `{"email":"  ALICE@EXAMPLE.COM  ","password":"secretpassword","confirm_password":"secretpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(validPayload))
	rec := httptest.NewRecorder()
	RegisterHandler(rec, req)
	fmt.Printf("=== Valid Request ===\nStatus: %d\nResponse: %s\n", rec.Code, rec.Body.String())

	// 2. Simulate POST /register with invalid input:
	invalidPayload := `{"email":"bad-email","password":"short","confirm_password":"mismatch"}`
	req = httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(invalidPayload))
	rec = httptest.NewRecorder()
	RegisterHandler(rec, req)
	fmt.Printf("=== Invalid Request ===\nStatus: %d\nResponse: %s\n", rec.Code, rec.Body.String())
}
