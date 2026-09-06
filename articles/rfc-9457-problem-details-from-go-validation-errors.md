---
title: Returning RFC 9457 Problem Details from Go Validation Errors
published: true
description: How to turn struct-tag validation failures into a standard "application/problem+json" response — the RFC 9457 format — with one method call, no hand-rolled error envelope.
tags: go, golang, api, webdev
canonical_url: https://dev.to/onurcinar/returning-rfc-9457-problem-details-from-go-validation-errors-4nb7
---

Most Go APIs invent their own validation error shape. One team returns `{"errors": [...]}`, another `{"field_errors": {...}}`, a third just a flat `{"error": "message"}` and hopes the client parses it. Every one of those is a private contract the client has to learn from your docs, because there's no shared shape for "here's what's wrong with your request."

[RFC 9457](https://www.rfc-editor.org/rfc/rfc9457) — *Problem Details for HTTP APIs* — is the IETF standard that fixes this: a `application/problem+json` body with `type`, `title`, `status`, and room for problem-specific extensions. [Checker](https://github.com/cinar/checker) now builds one of these directly from a failed struct validation, via `CheckErrors.ProblemDetails()`.

## The shape

RFC 9457 defines four base members — `type`, `title`, `status`, `detail`, `instance` — and lets a specific problem type add its own. For validation errors, [RFC 9457 §3.1](https://www.rfc-editor.org/rfc/rfc9457#section-3.1) sketches exactly this extension: an `invalid-params` array listing which fields failed and why. That's what Checker produces.

## From a failed struct to a problem+json body

Take a struct with a missing required field:

```go
type Person struct {
	Name string `checkers:"required"`
}

person := &Person{}

errs, ok := checker.CheckStruct(person)
if !ok {
	data, _ := json.Marshal(errs.ProblemDetails())
	fmt.Println(string(data))
}
```

```json
{
  "type": "about:blank",
  "title": "Your request parameters failed validation.",
  "status": 400,
  "invalid-params": [
    { "name": "Name", "reason": "Required value is missing.", "code": "REQUIRED" }
  ]
}
```

One method call — `errs.ProblemDetails()` — turns the same `CheckErrors` you'd otherwise call `.JSON()` on into a `*ProblemDetails` value, ready to marshal. `type` defaults to `"about:blank"` (RFC 9457's own default for "no more specific problem type registered"), `status` defaults to `400`, and each `invalid-params` entry carries the field `name`, a localized human-readable `reason`, and the machine-readable `code` — the same code you'd branch on client-side with the plain `.JSON()` output.

## Wiring it into a handler

The one thing RFC 9457 requires that a plain JSON body doesn't is the content type:

```go
func signupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	errs, valid := checker.CheckStruct(&req)
	if !valid {
		pd := errs.ProblemDetails()

		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(pd.Status)
		json.NewEncoder(w).Encode(pd)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
```

`pd.Status` drives both the header and the body, so they can't drift out of sync.

## Overriding the defaults

`Type`, `Title`, and `Status` are plain exported fields on the returned `*ProblemDetails` — RFC 9457 deliberately leaves them to the API producer, so Checker fills in reasonable defaults and gets out of your way. Set them before marshaling if you want a registered problem type or a different status code:

```go
pd := errs.ProblemDetails()
pd.Type = "https://api.example.com/problems/validation-error"
pd.Title = "Signup request failed validation"
pd.Status = http.StatusUnprocessableEntity // 422 instead of 400
```

Nothing else about the call changes — `invalid-params` is still built from the same `CheckErrors` map.

## Localized problem details

Since `ProblemDetails()` is just `ProblemDetailsWithLocale(DefaultLocale)` under the hood, a localized variant is one argument away — same [23 locales](https://github.com/cinar/checker#localized-error-messages) that back `JSONWithLocale`:

```go
pd := errs.ProblemDetailsWithLocale(locales.DeDE)
```

The `reason` strings come back in German; `type`, `status`, and `code` are untouched, since those aren't meant for display.

## Why bother with a standard shape

A hand-rolled error envelope works fine until you have more than one API, or a client library trying to handle errors generically, or an API gateway that wants to do something useful with 4xx bodies without special-casing your service. `application/problem+json` is already understood by tooling in several ecosystems — Spring, ASP.NET Core, and various API gateways emit or consume it natively — so producing it from Go costs you nothing extra (it's the same validation you were already running) and buys you a body shape other tools don't have to be taught.

## Try it

```bash
go get github.com/cinar/checker/v2
```

```go
import checker "github.com/cinar/checker/v2"
```

`ProblemDetails()` and `ProblemDetailsWithLocale()` sit right next to `JSON()` and `JSONWithLocale()` on `CheckErrors` — pick whichever body shape your API needs, from the same validation call. Full details in the [README](https://github.com/cinar/checker#structured-errors).
