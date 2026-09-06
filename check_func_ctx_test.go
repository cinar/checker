// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

type tenantKey struct{}

func tenantFromContext(ctx context.Context) string {
	tenant, _ := ctx.Value(tenantKey{}).(string)
	return tenant
}

func TestCheckWithContextSuccess(t *testing.T) {
	ctx := context.WithValue(context.Background(), tenantKey{}, "acme")

	upperWithTenant := func(ctx context.Context, value string) (string, error) {
		return tenantFromContext(ctx) + ":" + value, nil
	}

	actual, err := v2.CheckWithContext(ctx, "widget", upperWithTenant)
	if err != nil {
		t.Fatal(err)
	}

	if actual != "acme:widget" {
		t.Fatalf("actual %s expected acme:widget", actual)
	}
}

func TestCheckWithContextStopsAtFirstError(t *testing.T) {
	sentinel := errors.New("boom")

	calls := 0

	failing := func(_ context.Context, value string) (string, error) {
		calls++
		return value, sentinel
	}

	neverCalled := func(_ context.Context, value string) (string, error) {
		calls++
		return value, nil
	}

	_, err := v2.CheckWithContext(context.Background(), "value", failing, neverCalled)
	if !errors.Is(err, sentinel) {
		t.Fatalf("got unexpected error %v", err)
	}

	if calls != 1 {
		t.Fatalf("expected only the first check to run, got %d calls", calls)
	}
}

func TestRegisterCtxMakerAndCheckStructWithContext(t *testing.T) {
	v2.RegisterCtxMaker("unique-in-tenant", func(_ string) v2.CheckFuncCtx[reflect.Value] {
		return func(ctx context.Context, value reflect.Value) (reflect.Value, error) {
			if tenantFromContext(ctx) == "acme" && value.String() == "taken" {
				return value, v2.NewCheckError("NOT_UNIQUE")
			}

			return value, nil
		}
	})

	type User struct {
		Email string `checkers:"required" checkersCtx:"unique-in-tenant"`
	}

	ctx := context.WithValue(context.Background(), tenantKey{}, "acme")

	user := &User{Email: "taken"}

	errs, ok := v2.CheckStructWithContext(ctx, user)
	if ok {
		t.Fatal("expected errors")
	}

	if errs["Email"] == nil {
		t.Fatalf("expected an Email error, got %v", errs)
	}

	user.Email = "free"

	errs, ok = v2.CheckStructWithContext(ctx, user)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}

func TestCheckStructWithContextSkipsCtxCheckAfterFieldError(t *testing.T) {
	called := false

	v2.RegisterCtxMaker("never-called", func(_ string) v2.CheckFuncCtx[reflect.Value] {
		return func(_ context.Context, value reflect.Value) (reflect.Value, error) {
			called = true
			return value, nil
		}
	})

	type User struct {
		Email string `checkers:"required" checkersCtx:"never-called"`
	}

	user := &User{}

	errs, ok := v2.CheckStructWithContext(context.Background(), user)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Email"], v2.ErrRequired) {
		t.Fatalf("expected required error, got %v", errs)
	}

	if called {
		t.Fatal("expected the ctx checker to be skipped after the field-level check failed")
	}
}

func TestCheckStructIgnoresCheckersCtxTag(t *testing.T) {
	called := false

	v2.RegisterCtxMaker("should-not-run", func(_ string) v2.CheckFuncCtx[reflect.Value] {
		return func(_ context.Context, value reflect.Value) (reflect.Value, error) {
			called = true
			return value, nil
		}
	})

	type User struct {
		Email string `checkers:"required" checkersCtx:"should-not-run"`
	}

	user := &User{Email: "present"}

	_, ok := v2.CheckStruct(user)
	if !ok {
		t.Fatal("expected no errors")
	}

	if called {
		t.Fatal("expected CheckStruct to ignore the checkersCtx tag entirely")
	}
}

func TestCheckStructWithContextUnknownCtxChecker(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type User struct {
		Email string `checkersCtx:"unknown-ctx-checker"`
	}

	v2.CheckStructWithContext(context.Background(), &User{Email: "present"})
}

func TestRegisteredCtxMakerNamesIncludesRegistered(t *testing.T) {
	v2.RegisterCtxMaker("registered-for-names-test", func(_ string) v2.CheckFuncCtx[reflect.Value] {
		return func(_ context.Context, value reflect.Value) (reflect.Value, error) {
			return value, nil
		}
	})

	names := v2.RegisteredCtxMakerNames()

	found := false
	for _, name := range names {
		if name == "registered-for-names-test" {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected \"registered-for-names-test\" among registered ctx maker names, got %v", names)
	}
}
