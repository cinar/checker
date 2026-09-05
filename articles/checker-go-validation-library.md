---
title: Stop Hand-Rolling Validation in Go — Meet Checker
published: true
description: A zero-dependency Go library that validates and normalizes input from struct tags, ships 23 translated locales, generates JSON Schema from your rules, and plugs straight into Gin and Echo.
tags: go, golang, opensource, validation
canonical_url: https://dev.to/onurcinar/stop-hand-rolling-validation-in-go-meet-checker-nk1
cover_image: https://dev-to-uploads.s3.us-east-2.amazonaws.com/uploads/articles/tvu8xdvg2eumm3aeq5l0.png
---

If you've written more than one HTTP handler in Go, you've written this code:

```go
if strings.TrimSpace(req.Name) == "" {
	return errors.New("name is required")
}

if !strings.Contains(req.Email, "@") {
	return errors.New("invalid email")
}

if req.ConfirmPassword != req.Password {
	return errors.New("passwords do not match")
}
```

It starts small. Then a new field shows up, then a cross-field rule, then someone asks for a nicer error message, then someone else asks for that error message in Spanish, and suddenly your handler is 200 lines of `if` statements that nobody wants to touch.

[**Checker**](https://github.com/cinar/checker) is a Go library built to make that pile of `if` statements disappear — declaratively, with zero external dependencies. Here's why it's worth a look.

## The one-line pitch

```go
type Registration struct {
	Name            string `checkers:"trim required"`
	Email           string `checkers:"required email"`
	Password        string `checkers:"required min-len:8"`
	ConfirmPassword string `checkers:"eq-field:Password"`
}

errors, valid := checker.CheckStruct(&registration)
if !valid {
	// errors is a map[string]error, keyed by field name
}
```

That's it. Trim the name, require it, validate the email format, enforce a minimum password length, and confirm the two password fields match — all declared next to the field they apply to, all in one call.

## Why it's worth your attention

### Zero dependencies, really zero

The core `checker` module imports nothing beyond the Go standard library. Not `reflect`-based validator forks with a dozen transitive deps, not a YAML parser you didn't ask for. `go get github.com/cinar/checker/v2` pulls in exactly one thing: Checker itself. That matters for supply-chain surface area, build times, and not having to explain to a security review why your validation library needs 40 packages.

### Checkers *and* normalizers, in one pipeline

Most validation libraries only check — they tell you your input is wrong. Checker also fixes it. `trim`, `lower`, `upper`, `title`, HTML/URL escaping — these are normalizers, and they share the exact same pipeline as checkers like `required` or `email`. Mix them freely:

```go
type Person struct {
	Name string `checkers:"trim title required"`
}
```

Trim the whitespace, title-case it, then make sure something is actually left. One declaration, no separate "sanitize" pass before your "validate" pass.

### Cross-field and conditional rules, without a callback

A password-confirmation field. A "State" field that's only required if "Country" is "US". A return date that has to be after the departure date. These usually mean dropping out of struct tags entirely and writing custom validation functions. Checker handles them as tags:

```go
type Trip struct {
	Country  string `checkers:"required"`
	State    string `checkers:"required-if:Country:US"`
	DepartAt string `checkers:"required"`
	ReturnAt string `checkers:"required after-field:DateOnly:DepartAt"`
}
```

`eq-field`, `required-if`, `required-unless`, `before-field`, `after-field` — all comparing sibling struct fields, all declared inline.

### Slices and maps, checked at both levels

Need to make sure a slice has at most 2 emails, *and* that each email is at most 64 characters? The `@` prefix separates container-level rules from item-level ones:

```go
type Person struct {
	Emails map[string]string `checkers:"@max-len:2 trim max-len:64"`
}
```

`@max-len:2` caps the map at two entries. `trim max-len:64` runs on every value in it. Nested structs and pointers inside a slice or map get walked and checked too — this isn't a shallow, top-level-fields-only validator.

### 23 languages, opt-in

Ship a SaaS product to a global audience and "Not a valid email address." in every locale stops being acceptable pretty fast. Checker ships **23 translated locales** out of the box — the same set `go-playground/validator` supports — and none of them cost you anything unless you ask for them:

```go
checker.RegisterLocale(locales.DeDE, locales.DeDEMessages)

_, err := checker.IsEmail("abcd")
fmt.Println(err.ErrorWithLocale(locales.DeDE))
// Keine gültige E-Mail-Adresse.
```

Only `en-US` is registered by default, so importing `checker` never silently pulls translation data into your binary that you don't use.

### Structured errors that are API-ready

`CheckStruct` doesn't just hand you a generic `error`. It returns `CheckErrors`, a `map[string]error` keyed by field name that *also* implements the `error` interface, so you can return it directly. When you're building an HTTP API, call `.JSON()` and you're done:

```go
errs, valid := checker.CheckStruct(&registration)
if !valid {
	data, _ := errs.JSON()
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)
	// {"Name":{"code":"REQUIRED","message":"Required value is missing."}}
	return
}
```

Machine-readable `code`, human-readable `message`, per field, ready to serialize. `JSONWithLocale()` does the same thing localized.

### Your validation rules, turned into a JSON Schema — for free

This is the feature that stops people mid-scroll: Checker can generate a **JSON Schema** document directly from your struct tags.

```go
type Person struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

schema := checker.JSONSchema(&Person{})
```

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "Person",
  "type": "object",
  "properties": {
    "email": { "type": "string", "format": "email" },
    "name": { "type": "string" }
  },
  "required": ["email", "name"]
}
```

`required` becomes `required`, `min-len`/`max-len` become `minLength`/`maxLength` (or `minItems`/`maxItems`, `minProperties`/`maxProperties` for slices and maps), `gte`/`lte` become `minimum`/`maximum`, and `email`/`url`/`ipv4`/`fqdn` become a `format`. A checker with no schema equivalent isn't silently dropped — it lands in an `x-checker` vendor extension so nothing goes missing.

Translation: your Go validation tags *are* your API documentation and your frontend validation spec. Stop maintaining the same rules three times.

### Drop-in adapters for Gin and Echo

If you're on Gin or Echo, binding and validating a request body is one call:

```go
import checkergin "github.com/cinar/checker/v2/gin"

router.POST("/register", func(c *gin.Context) {
	var registration Registration

	if !checkergin.Bind(c, &registration) {
		return // 400 already written
	}

	c.JSON(http.StatusOK, registration)
})
```

Both adapters are separately-versioned Go modules, so Gin or Echo only enters your dependency tree if you actually `go get` the adapter — the core library stays dependency-free regardless of which framework you use.

### Extensible when the built-ins aren't enough

Checker ships 30+ checkers (email, URL, IP/IPv4/IPv6, CIDR, MAC, credit card, ISBN, hashes, country and language codes, an Ethereum address checker, and more), but you're not boxed in. Register your own:

```go
checker.RegisterMaker("is-fruit", func(params string) checker.CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		s := value.Interface().(string)
		if s == "apple" || s == "banana" {
			return value, nil
		}
		return value, checker.NewCheckError("NOT_FRUIT")
	}
})
```

Once registered, `is-fruit` works in struct tags exactly like a built-in checker — and you can teach `JSONSchema` how to represent it too, via `RegisterSchemaMaker`.

## Built like a library that means it

A detail that's easy to gloss over: this project enforces **100% test coverage**. Every checker, every normalizer, every branch has a matching test. There's also a `locales_test.go` that fails the build if any locale is missing a message for a given error code, or if a placeholder doesn't match `en-US`. That's the kind of quiet discipline that keeps a validation library from lying to you in production.

## Try it

```bash
go get github.com/cinar/checker/v2
```

```go
import checker "github.com/cinar/checker/v2"
```

Check out the project on GitHub: **[github.com/cinar/checker](https://github.com/cinar/checker)**

If you're validating structs by hand in Go right now, give it five minutes. If you find a checker missing, or a locale that's rough around the edges — PRs are welcome.
