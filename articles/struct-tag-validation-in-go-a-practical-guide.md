---
title: Struct Tag Validation in Go: A Practical Guide
published: true
description: A hands-on walkthrough of validating and normalizing a real signup form in Go — trimming input, cross-field rules, slices, custom messages, and JSON error responses — using struct tags.
tags: go, golang, tutorial, webdev
canonical_url: https://dev.to/onurcinar/struct-tag-validation-in-go-a-practical-guide-j4d
---

Every Go API ends up needing the same thing: take a JSON body, clean it up, make sure it's valid, and tell the client exactly what's wrong if it isn't. This guide walks through building that for a real signup form, using struct tags instead of hand-written `if` chains, with [Checker](https://github.com/cinar/checker) as the tool doing the work.

By the end you'll have a `SignupRequest` that trims and lowercases its own input, validates it, enforces a cross-field password match, checks a slice of roles, returns friendly per-field error messages, and serializes cleanly to JSON — in about 15 lines of tags.

## Setup

```bash
go get github.com/cinar/checker/v2
```

```go
import checker "github.com/cinar/checker/v2"
```

## Step 1: A struct with checkers, not `if` statements

Start with a plain struct and describe each field's rules as a `checkers` tag:

```go
type SignupRequest struct {
	Email           string `json:"email" checkers:"trim lower required email"`
	Password        string `json:"password" checkers:"required min-len:8"`
	ConfirmPassword string `json:"confirm_password" checkers:"required eq-field:Password"`
}
```

Read a tag left to right — it's a pipeline, executed in order:

- `trim lower required email` on `Email`: trim whitespace, lowercase it, make sure something is left, then check it looks like an email. Order matters here — `required` runs *after* `trim`, so a field of just `"   "` correctly fails.
- `required min-len:8` on `Password`: must be present, must be at least 8 characters.
- `required eq-field:Password` on `ConfirmPassword`: must be present, and must equal the sibling `Password` field.

Note that `eq-field` isn't comparing against a hardcoded value — it's reading another field off the same struct at validation time. That's a "field-relative" checker; Checker also ships `after-field`, `before-field`, `required-if`, and `required-unless` for the same pattern (a state field required only when country is `"US"`, a return date that has to be after a departure date, and so on).

## Step 2: Validate it

```go
func main() {
	req := &SignupRequest{
		Email:           "  ALICE@EXAMPLE.COM  ",
		Password:        "supersecret123",
		ConfirmPassword: "supersecret123",
	}

	errs, valid := checker.CheckStruct(req)
	if !valid {
		data, _ := errs.JSON()
		fmt.Println(string(data))
		return
	}

	fmt.Println(req.Email) // "alice@example.com"
}
```

`CheckStruct` does two things in one pass: it normalizes the struct **in-place** and validates it. Look at `req.Email` after the call — it's already trimmed and lowercased, ready to hash the password against or write to a database, with no separate "sanitize" step beforehand.

If validation fails, `errs` is a `CheckErrors` — a `map[string]error` keyed by field name, which also implements `error` itself, so you can `return errs` directly from a function that expects one. Calling `.JSON()` on it gives you an HTTP-API-ready body:

```json
{
  "Password": {
    "code": "NOT_MIN_LEN",
    "message": "Value cannot be less than 8."
  }
}
```

Machine-readable `code` for client-side logic, human-readable `message` for display — per field, no manual formatting.

## Step 3: Add a slice field

Signup forms often collect more than scalars. Say `SignupRequest` also takes a list of roles, capped at 3, each one trimmed and alphanumeric-only:

```go
type SignupRequest struct {
	Email           string   `json:"email" checkers:"trim lower required email"`
	Password        string   `json:"password" checkers:"required min-len:8"`
	ConfirmPassword string   `json:"confirm_password" checkers:"required eq-field:Password"`
	Roles           []string `json:"roles" checkers:"@max-len:3 trim alphanumeric"`
}
```

The `@` prefix is what separates *container*-level rules from *item*-level ones in the same tag. `@max-len:3` caps the slice at three entries; `trim alphanumeric` (no `@`) runs on every individual string inside it. Given `Roles: []string{"  admin  ", "editor!"}`, after `CheckStruct` runs, `Roles[0]` becomes `"admin"` and `Roles[1]` fails `alphanumeric` because of the `!`. The same `@` split works for maps, and nested structs inside slices/maps are walked recursively — this isn't a top-level-fields-only validator.

## Step 4: Make a field optional

Not every field should be required. Add an optional website field:

```go
Website string `json:"website" checkers:"omitempty url"`
```

`omitempty` skips every other checker in the tag when the field is its zero value, but still runs them normally once a value is present. An empty `Website` is fine; `"not-a-url"` is not. It looks at the field's *original* value, so `trim omitempty required` on an all-whitespace string still fails `required` after trimming — whitespace isn't the zero value to begin with. (Don't pair `omitempty` with `required` on the same field — that's a contradiction, and `omitempty` wins, so `required` never runs.)

## Step 5: Give one field a friendlier error message

The default messages are fine for logs, but a signup form probably wants nicer copy for `Email` specifically. Add a `checkersMsg` tag alongside `checkers` — a semicolon-separated list of `name=message` pairs, keyed by the bare checker name:

```go
Email string `json:"email" checkers:"trim lower required email" checkersMsg:"required=Email is required;email=Enter a valid email address"`
```

This overrides the message only for this field, only for these checkers — every other field using `required` or `email` elsewhere in your codebase keeps the default (or localized) wording.

## Step 6: Wire it into an HTTP handler

Put it all together with the standard library — no framework required:

```go
func signupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	errs, valid := checker.CheckStruct(&req)
	if !valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		data, _ := errs.JSON()
		w.Write(data)
		return
	}

	// req is normalized and valid — safe to persist.
	w.WriteHeader(http.StatusCreated)
}
```

If you're on Gin or Echo instead, the separately-versioned adapter modules collapse decode-and-validate into one call:

```go
import checkergin "github.com/cinar/checker/v2/gin"

router.POST("/signup", func(c *gin.Context) {
	var req SignupRequest
	if !checkergin.Bind(c, &req) {
		return // 400 already written
	}
	c.JSON(http.StatusOK, req)
})
```

## What you get for free once the tags exist

Because the rules live as tags on the type rather than scattered across handler code, they're reusable outside validation itself. `checker.JSONSchema(&SignupRequest{})` walks the same tags and produces a Draft 2020-12 JSON Schema document — `required` becomes `required`, `min-len` becomes `minLength`, `email` becomes `format: "email"` — so your Go validation rules can double as API documentation or a frontend contract, instead of being hand-copied into a second place that quietly drifts out of sync.

## Try it yourself

```bash
go get github.com/cinar/checker/v2
```

There's a runnable version of the core trim/validate/cross-field pattern on the [Go Playground](https://go.dev/play/p/FfkXm5oC9ii), and the full checker list — 30+ built-ins covering emails, URLs, IPs, credit cards, hashes, country codes, and more — is in the [README](https://github.com/cinar/checker#checkers-provided). If a rule you need isn't built in, `RegisterMaker` lets you add your own and it behaves exactly like a first-party checker in struct tags.
