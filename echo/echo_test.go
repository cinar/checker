// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package echo_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	checker "github.com/cinar/checker/v2"
	checkerecho "github.com/cinar/checker/v2/echo"
)

// Registration is used across the tests as the struct being bound and checked.
type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

func newTestContext(body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()

	return e.NewContext(req, recorder), recorder
}

func TestBindSuccess(t *testing.T) {
	c, recorder := newTestContext(`{"name":"  Onur Cinar  ","email":"onur@example.com"}`)

	var registration Registration

	if !checkerecho.Bind(c, &registration) {
		t.Fatalf("expected bind to succeed, got status %d body %s", recorder.Code, recorder.Body.String())
	}

	if registration.Name != "Onur Cinar" {
		t.Fatalf("expected the trim normalizer to be applied, got %q", registration.Name)
	}
}

func TestBindInvalidJSON(t *testing.T) {
	c, recorder := newTestContext(`{"name":`)

	var registration Registration

	if checkerecho.Bind(c, &registration) {
		t.Fatal("expected bind to fail")
	}

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	var body map[string]string

	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}

	if body["error"] == "" {
		t.Fatalf("expected a bind error message, got %v", body)
	}
}

func TestBindValidationFailure(t *testing.T) {
	c, recorder := newTestContext(`{"name":"","email":""}`)

	var registration Registration

	if checkerecho.Bind(c, &registration) {
		t.Fatal("expected bind to fail")
	}

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	var fields map[string]checker.FieldError

	if err := json.Unmarshal(recorder.Body.Bytes(), &fields); err != nil {
		t.Fatal(err)
	}

	if fields["name"].Code != "REQUIRED" {
		t.Fatalf("expected name required, got %+v", fields["name"])
	}

	if fields["email"].Code != "REQUIRED" {
		t.Fatalf("expected email required, got %+v", fields["email"])
	}
}

func TestCheckSuccess(t *testing.T) {
	c, recorder := newTestContext("")

	registration := &Registration{
		Name:  "Onur Cinar",
		Email: "onur@example.com",
	}

	if !checkerecho.Check(c, registration) {
		t.Fatalf("expected check to succeed, got status %d body %s", recorder.Code, recorder.Body.String())
	}
}

func TestCheckFailure(t *testing.T) {
	c, recorder := newTestContext("")

	registration := &Registration{
		Name:  "Onur Cinar",
		Email: "not-an-email",
	}

	if checkerecho.Check(c, registration) {
		t.Fatal("expected check to fail")
	}

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	var fields map[string]checker.FieldError

	if err := json.Unmarshal(recorder.Body.Bytes(), &fields); err != nil {
		t.Fatal(err)
	}

	if fields["email"].Code != "NOT_EMAIL" {
		t.Fatalf("expected email not valid, got %+v", fields["email"])
	}
}
