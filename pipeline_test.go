// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

type pipelineUser struct {
	Email         string
	Role          string
	HasMFAEnabled bool
}

type ctxKey string

const usernameTakenKey ctxKey = "username-taken"

func ExampleField() {
	type User struct {
		Email string
	}

	user := &User{Email: "  ALICE@EXAMPLE.COM  "}

	pipeline := v2.NewPipeline[User]().Step(
		v2.Field("Email", func(u *User) *string { return &u.Email }, v2.TrimSpace, v2.Lower, v2.Required, v2.IsEmail),
	)

	errs, ok := pipeline.Validate(context.Background(), user)
	if !ok {
		fmt.Println(errs)
		return
	}

	fmt.Println(user.Email)
	// Output: alice@example.com
}

func ExampleRule() {
	user := &pipelineUser{Role: "admin", HasMFAEnabled: false}

	pipeline := v2.NewPipeline[pipelineUser]().Step(
		v2.Rule("MFA", func(_ context.Context, u *pipelineUser) error {
			if u.Role == "admin" && !u.HasMFAEnabled {
				return v2.NewCheckError("MFA_REQUIRED_FOR_ADMIN")
			}

			return nil
		}),
	)

	errs, ok := pipeline.Validate(context.Background(), user)
	if !ok {
		fmt.Println(errs)
	}
	// Output: MFA: MFA_REQUIRED_FOR_ADMIN
}

func TestPipelineFieldNormalizesAndValidates(t *testing.T) {
	user := &pipelineUser{Email: "  Onur@Example.com  "}

	pipeline := v2.NewPipeline[pipelineUser]().Step(
		v2.Field("Email", func(u *pipelineUser) *string { return &u.Email }, v2.TrimSpace, v2.Lower, v2.Required, v2.IsEmail),
	)

	errs, ok := pipeline.Validate(context.Background(), user)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if user.Email != "onur@example.com" {
		t.Fatalf("expected normalized email, got %q", user.Email)
	}
}

func TestPipelineFieldStopsAtFirstFailingCheck(t *testing.T) {
	user := &pipelineUser{Email: ""}

	pipeline := v2.NewPipeline[pipelineUser]().Step(
		v2.Field("Email", func(u *pipelineUser) *string { return &u.Email }, v2.Required, v2.IsEmail),
	)

	errs, ok := pipeline.Validate(context.Background(), user)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Email"], v2.ErrRequired) {
		t.Fatalf("expected Email required, got %v", errs)
	}
}

func TestPipelineRuleReceivesContext(t *testing.T) {
	user := &pipelineUser{Email: "taken@example.com"}

	pipeline := v2.NewPipeline[pipelineUser]().Step(
		v2.Rule("Email", func(ctx context.Context, u *pipelineUser) error {
			if taken, _ := ctx.Value(usernameTakenKey).(string); taken == u.Email {
				return v2.NewCheckError("EMAIL_TAKEN")
			}

			return nil
		}),
	)

	ctx := context.WithValue(context.Background(), usernameTakenKey, "taken@example.com")

	errs, ok := pipeline.Validate(ctx, user)
	if ok {
		t.Fatal("expected errors")
	}

	var checkErr *v2.CheckError
	if !errors.As(errs["Email"], &checkErr) || checkErr.Code != "EMAIL_TAKEN" {
		t.Fatalf("expected EMAIL_TAKEN, got %v", errs)
	}
}

func TestPipelineRunsEveryStepRegardlessOfEarlierFailures(t *testing.T) {
	user := &pipelineUser{Email: "", Role: "admin", HasMFAEnabled: false}

	pipeline := v2.NewPipeline[pipelineUser]().
		Step(v2.Field("Email", func(u *pipelineUser) *string { return &u.Email }, v2.Required)).
		Step(v2.Rule("MFA", func(_ context.Context, u *pipelineUser) error {
			if u.Role == "admin" && !u.HasMFAEnabled {
				return v2.NewCheckError("MFA_REQUIRED_FOR_ADMIN")
			}

			return nil
		}))

	errs, ok := pipeline.Validate(context.Background(), user)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Email"], v2.ErrRequired) {
		t.Fatalf("expected Email required, got %v", errs)
	}

	if errs["MFA"] == nil {
		t.Fatalf("expected MFA error alongside Email error, got %v", errs)
	}
}

func TestPipelineStepAcceptsMultipleAtOnce(t *testing.T) {
	user := &pipelineUser{Email: "onur@example.com", Role: "user", HasMFAEnabled: true}

	pipeline := v2.NewPipeline[pipelineUser]().Step(
		v2.Field("Email", func(u *pipelineUser) *string { return &u.Email }, v2.Required, v2.IsEmail),
		v2.Rule("MFA", func(_ context.Context, u *pipelineUser) error {
			if u.Role == "admin" && !u.HasMFAEnabled {
				return v2.NewCheckError("MFA_REQUIRED_FOR_ADMIN")
			}

			return nil
		}),
	)

	if errs, ok := pipeline.Validate(context.Background(), user); !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}
