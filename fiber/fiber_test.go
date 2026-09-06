// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package fiber_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"

	checker "github.com/cinar/checker/v2"
	checkerfiber "github.com/cinar/checker/v2/fiber"
)

// Registration is used across the tests as the struct being bound and checked.
type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

func doRequest(t *testing.T, app *fiber.App, method, body string) (*http.Response, []byte) {
	t.Helper()

	req := httptest.NewRequest(method, "/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	return resp, data
}

func newBindApp() *fiber.App {
	app := fiber.New()

	app.Post("/register", func(c fiber.Ctx) error {
		var registration Registration

		if !checkerfiber.Bind(c, &registration) {
			return nil
		}

		return c.JSON(registration)
	})

	return app
}

func TestBindSuccess(t *testing.T) {
	resp, data := doRequest(t, newBindApp(), http.MethodPost, `{"name":"  Onur Cinar  ","email":"onur@example.com"}`)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d body %s", http.StatusOK, resp.StatusCode, data)
	}

	var registration Registration

	if err := json.Unmarshal(data, &registration); err != nil {
		t.Fatal(err)
	}

	if registration.Name != "Onur Cinar" {
		t.Fatalf("expected the trim normalizer to be applied, got %q", registration.Name)
	}
}

func TestBindInvalidJSON(t *testing.T) {
	resp, data := doRequest(t, newBindApp(), http.MethodPost, `{"name":`)

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var body map[string]string

	if err := json.Unmarshal(data, &body); err != nil {
		t.Fatal(err)
	}

	if body["error"] == "" {
		t.Fatalf("expected a bind error message, got %v", body)
	}
}

func TestBindValidationFailure(t *testing.T) {
	resp, data := doRequest(t, newBindApp(), http.MethodPost, `{"name":"","email":""}`)

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var fields map[string]checker.FieldError

	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}

	if fields["name"].Code != "REQUIRED" {
		t.Fatalf("expected name required, got %+v", fields["name"])
	}

	if fields["email"].Code != "REQUIRED" {
		t.Fatalf("expected email required, got %+v", fields["email"])
	}
}

func newCheckApp(registration *Registration) *fiber.App {
	app := fiber.New()

	app.Post("/register", func(c fiber.Ctx) error {
		if !checkerfiber.Check(c, registration) {
			return nil
		}

		return c.JSON(registration)
	})

	return app
}

func TestCheckSuccess(t *testing.T) {
	registration := &Registration{
		Name:  "Onur Cinar",
		Email: "onur@example.com",
	}

	resp, data := doRequest(t, newCheckApp(registration), http.MethodPost, "")

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected check to succeed, got status %d body %s", resp.StatusCode, data)
	}
}

func TestCheckFailure(t *testing.T) {
	registration := &Registration{
		Name:  "Onur Cinar",
		Email: "not-an-email",
	}

	resp, data := doRequest(t, newCheckApp(registration), http.MethodPost, "")

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content type %q", resp.Header.Get("Content-Type"))
	}

	var fields map[string]checker.FieldError

	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}

	if fields["email"].Code != "NOT_EMAIL" {
		t.Fatalf("expected email not valid, got %+v", fields["email"])
	}
}
