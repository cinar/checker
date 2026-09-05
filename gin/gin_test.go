// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package gin_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	checker "github.com/cinar/checker/v2"
	checkergin "github.com/cinar/checker/v2/gin"
)

// Registration is used across the tests as the struct being bound and checked.
type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

func newTestContext(method, body string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	req := httptest.NewRequest(method, "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	return c, recorder
}

func TestBindSuccess(t *testing.T) {
	c, recorder := newTestContext(http.MethodPost, `{"name":"  Onur Cinar  ","email":"onur@example.com"}`)

	var registration Registration

	if !checkergin.Bind(c, &registration) {
		t.Fatalf("expected bind to succeed, got status %d body %s", recorder.Code, recorder.Body.String())
	}

	if registration.Name != "Onur Cinar" {
		t.Fatalf("expected the trim normalizer to be applied, got %q", registration.Name)
	}
}

func TestBindInvalidJSON(t *testing.T) {
	c, recorder := newTestContext(http.MethodPost, `{"name":`)

	var registration Registration

	if checkergin.Bind(c, &registration) {
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
	c, recorder := newTestContext(http.MethodPost, `{"name":"","email":""}`)

	var registration Registration

	if checkergin.Bind(c, &registration) {
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
	c, recorder := newTestContext(http.MethodPost, "")

	registration := &Registration{
		Name:  "Onur Cinar",
		Email: "onur@example.com",
	}

	if !checkergin.Check(c, registration) {
		t.Fatalf("expected check to succeed, got status %d body %s", recorder.Code, recorder.Body.String())
	}
}

func TestCheckFailure(t *testing.T) {
	c, recorder := newTestContext(http.MethodPost, "")

	registration := &Registration{
		Name:  "Onur Cinar",
		Email: "not-an-email",
	}

	if checkergin.Check(c, registration) {
		t.Fatal("expected check to fail")
	}

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	if recorder.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content type %q", recorder.Header().Get("Content-Type"))
	}

	var fields map[string]checker.FieldError

	if err := json.Unmarshal(recorder.Body.Bytes(), &fields); err != nil {
		t.Fatal(err)
	}

	if fields["email"].Code != "NOT_EMAIL" {
		t.Fatalf("expected email not valid, got %+v", fields["email"])
	}
}
