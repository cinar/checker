// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checkergen

import (
	"fmt"
	"strconv"
	"strings"
)

// sibling resolves a sibling field's Go access expression (e.g.
// "v.Password"), for a field-relative checker's target field name. It
// returns an error if no field with that Go name exists on the struct
// currently being generated. Deliberately doesn't also resolve the
// sibling's json-tag-facing name: a field-relative checker's error message
// uses the field name exactly as written in the tag (the Go field name
// used to look it up), same as CheckStruct/IsEqField/IsRequiredIf/
// IsRequiredUnless do at runtime -- it's the raw params/name text callers
// below already have, not something sibling needs to look up.
type sibling func(name string) (expr string, err error)

// callSpec turns one checkers-tag token (already split into name and
// params) into a Go source expression of type checker.CheckFunc[fieldType],
// suitable for splicing into a checker.Check(fieldExpr, <expr>, ...) call.
// fieldType is the field's Go type exactly as it should appear in
// generated source (e.g. "string", "int"); see the package doc for the
// (deliberately narrow, for v1) set of types this ever gets called with.
type callSpec func(params, fieldType string, sib sibling) (string, error)

// bare returns a callSpec for a checker/normalizer whose exported function
// already matches checker.CheckFunc[T] exactly (single T argument, (T,
// error) result) and takes no parameters from the tag: the bare qualified
// identifier is used directly, relying on Go's type inference from
// checker.Check's other arguments to instantiate any generic type
// parameter. Panics if params is non-empty, mirroring how a runtime
// make<Name> ignores or rejects unexpected parameters.
func bare(funcName string) callSpec {
	return func(params, _ string, _ sibling) (string, error) {
		return "checker." + funcName, nil
	}
}

// closure returns a callSpec that wraps funcName in a closure literal,
// splicing in args (already-rendered Go argument expressions, in call
// order) built from the tag's params.
func closure(funcName string, args func(params string) ([]string, error)) callSpec {
	return func(params, fieldType string, _ sibling) (string, error) {
		argExprs, err := args(params)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(
			"func(value %s) (%s, error) { return checker.%s(%s) }",
			fieldType, fieldType, funcName, strings.Join(argExprs, ", "),
		), nil
	}
}

// curried returns a callSpec for one of the generic "maker" functions
// (MinLen, MaxLen, Len, Default) that take the tag's parameter directly and
// return a checker.CheckFunc[T] themselves: T is spliced in explicitly,
// since (unlike the checkers wrapped by closure/bare) none of these
// functions' own arguments are of type T for Go to infer it from.
func curried(funcName string, arg func(params, fieldType string) (string, error)) callSpec {
	return func(params, fieldType string, _ sibling) (string, error) {
		argExpr, err := arg(params, fieldType)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("checker.%s[%s](%s)", funcName, fieldType, argExpr), nil
	}
}

// requireType wraps spec so it fails with a clear checkergen-level error if
// fieldType isn't exactly want, instead of letting a mismatched call
// through to fail as a confusing compile error in the generated file. Used
// for the handful of checkers (finite, int, multiple-of) whose exported
// function is hardcoded to float64 rather than generic over T, unlike
// every other numeric checker (gt, gte, lt, lte, ...), which accepts any
// cmp.Ordered T.
func requireType(checkerName, want string, spec callSpec) callSpec {
	return func(params, fieldType string, sib sibling) (string, error) {
		if fieldType != want {
			return "", fmt.Errorf("%s requires a %s field, but the field's type is %s", checkerName, want, fieldType)
		}

		return spec(params, fieldType, sib)
	}
}

// splitParam splits params on ":" into exactly two non-empty parts,
// mirroring the strings.Cut+panic pattern every two-part runtime maker
// (after, after-field, required-if, ...) already uses, so a malformed tag
// fails at generate time with the same clarity it would at runtime.
func splitParam(checkerName, params string) (first, second string) {
	first, second, found := strings.Cut(params, ":")
	if !found || first == "" || second == "" {
		panic(fmt.Sprintf("%s requires two ':'-separated parameters, got %q", checkerName, params))
	}

	return first, second
}

// numberLiteral returns params verbatim, after confirming it parses as a
// number: an untyped Go numeric literal spliced into a call adapts to
// whatever numeric field type it's assigned/compared against, so the
// source text itself doesn't need to change per field type.
func numberLiteral(checkerName, params string) (string, error) {
	if _, err := strconv.ParseFloat(params, 64); err != nil {
		return "", fmt.Errorf("%s: parameter %q is not a number", checkerName, params)
	}

	return params, nil
}

// intLiteral returns params verbatim, after confirming it parses as an
// integer.
func intLiteral(checkerName, params string) (string, error) {
	if _, err := strconv.Atoi(params); err != nil {
		return "", fmt.Errorf("%s: parameter %q is not an integer", checkerName, params)
	}

	return params, nil
}

// callSpecs maps every checker/normalizer name checkergen knows how to
// generate to its callSpec. A name with no entry here fails generation for
// the struct it's on with a clear diagnostic (see Generate), rather than
// silently producing incomplete validation -- extending this table for an
// unmapped checker, including a custom one registered via
// checker.RegisterMaker, is the intended way to teach checkergen about it.
var callSpecs = map[string]callSpec{
	// Normalizers: bare, no parameters, string only.
	"trim":            bare("TrimSpace"),
	"trim-left":       bare("TrimLeft"),
	"trim-right":      bare("TrimRight"),
	"lower":           bare("Lower"),
	"upper":           bare("Upper"),
	"title":           bare("Title"),
	"html-escape":     bare("HTMLEscape"),
	"html-unescape":   bare("HTMLUnescape"),
	"url-escape":      bare("URLEscape"),
	"url-unescape":    bare("URLUnescape"),
	"strip-invisible": bare("StripInvisible"),

	// Checkers: bare, no parameters.
	"required":          bare("Required"),
	"finite":            requireType("finite", "float64", bare("IsFinite")),
	"int":               requireType("int", "float64", bare("IsInt")),
	"positive":          bare("IsPositive"),
	"negative":          bare("IsNegative"),
	"nonnegative":       bare("IsNonnegative"),
	"ascii":             bare("IsASCII"),
	"alpha":             bare("IsAlpha"),
	"alphanumeric":      bare("IsAlphanumeric"),
	"base64":            bare("IsBase64"),
	"base64-url":        bare("IsBase64URL"),
	"cidr":              bare("IsCIDR"),
	"digits":            bare("IsDigits"),
	"email":             bare("IsEmail"),
	"eoa":               bare("IsEOA"),
	"fqdn":              bare("IsFQDN"),
	"hex":               bare("IsHex"),
	"hex-color":         bare("IsHexColor"),
	"iban":              bare("IsIBAN"),
	"ip":                bare("IsIP"),
	"ipv4":              bare("IsIPv4"),
	"ipv6":              bare("IsIPv6"),
	"isbn":              bare("IsISBN"),
	"iso3166-1-alpha-2": bare("IsISO31661Alpha2"),
	"iso3166-1-alpha-3": bare("IsISO31661Alpha3"),
	"iso639-1":          bare("IsISO6391"),
	"jwt":               bare("IsJWT"),
	"luhn":              bare("IsLUHN"),
	"mac":               bare("IsMAC"),
	"mongo-id":          bare("IsMongoID"),
	"numeric":           bare("IsNumeric"),
	"semver":            bare("IsSemver"),
	"slug":              bare("IsSlug"),
	"ulid":              bare("IsULID"),
	"url":               bare("IsURL"),
	"uuid":              bare("IsUUID"),
	"e164":              bare("IsE164"),
	"credit-card":       bare("IsAnyCreditCard"),

	// Checkers needing one parameter, spliced ahead of value.
	"contains":    closure("IsContains", oneStringArgBefore("contains")),
	"starts-with": closure("IsStartsWith", oneStringArgBefore("starts-with")),
	"ends-with":   closure("IsEndsWith", oneStringArgBefore("ends-with")),
	"regexp":      closure("IsRegexp", oneStringArgBefore("regexp")),
	"time":        closure("IsTime", oneStringArgBefore("time")),

	"hash":        closure("IsHash", oneStringArgBefore("hash")),
	"postal-code": closure("IsPostalCode", oneStringArgBefore("postal-code")),

	// Checkers needing one parameter, spliced after value.
	"eq":          closure("IsEq", oneNumberOrStringArgAfter("eq")),
	"ne":          closure("IsNe", oneNumberOrStringArgAfter("ne")),
	"gt":          closure("IsGt", oneNumberArgAfter("gt")),
	"gte":         closure("IsGte", oneNumberArgAfter("gte")),
	"lt":          closure("IsLt", oneNumberArgAfter("lt")),
	"lte":         closure("IsLte", oneNumberArgAfter("lte")),
	"multiple-of": requireType("multiple-of", "float64", closure("IsMultipleOf", oneNumberArgAfter("multiple-of"))),

	"oneof": closure("IsOneOf", oneOfArgs),

	// Checkers needing two ":"-separated parameters.
	"after":  closure("IsAfter", twoStringArgsBefore("after")),
	"before": closure("IsBefore", twoStringArgsBefore("before")),

	// Curried "maker" generics: T spliced in explicitly.
	"min-len": curried("MinLen", intArg("min-len")),
	"max-len": curried("MaxLen", intArg("max-len")),
	"len":     curried("Len", intArg("len")),
	"default": curried("Default", defaultArg),
}

// fieldCallSpecs maps the five field-relative checker names to their
// callSpec, kept separate from callSpecs since they're the only ones that
// use the sib parameter to resolve a sibling field's access expression and
// tag-facing name instead of a literal from params.
var fieldCallSpecs = map[string]callSpec{
	"eq-field": func(params, fieldType string, sib sibling) (string, error) {
		expr, err := sib(params)
		if err != nil {
			return "", err
		}

		// The error's field name is params (the sibling's Go field name,
		// exactly as written in the tag), not its json-tag-facing name --
		// matching IsEqField/checkEqField's own runtime behavior.
		return fmt.Sprintf(
			"func(value %s) (%s, error) { return checker.IsEqField(value, %s, %s) }",
			fieldType, fieldType, expr, strconv.Quote(params),
		), nil
	},

	"required-if": func(params, fieldType string, sib sibling) (string, error) {
		name, expected := splitParam("required-if", params)

		expr, err := sib(name)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(
			"func(value %s) (%s, error) { return checker.IsRequiredIf(value, %s, %s) }",
			fieldType, fieldType, expr, strconv.Quote(expected),
		), nil
	},

	"required-unless": func(params, fieldType string, sib sibling) (string, error) {
		name, expected := splitParam("required-unless", params)

		expr, err := sib(name)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(
			"func(value %s) (%s, error) { return checker.IsRequiredUnless(value, %s, %s) }",
			fieldType, fieldType, expr, strconv.Quote(expected),
		), nil
	},

	"after-field": func(params, fieldType string, sib sibling) (string, error) {
		layout, name := splitParam("after-field", params)

		expr, err := sib(name)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(
			"func(value %s) (%s, error) { return checker.IsAfterField(%s, %s, value) }",
			fieldType, fieldType, strconv.Quote(layout), expr,
		), nil
	},

	"before-field": func(params, fieldType string, sib sibling) (string, error) {
		layout, name := splitParam("before-field", params)

		expr, err := sib(name)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(
			"func(value %s) (%s, error) { return checker.IsBeforeField(%s, %s, value) }",
			fieldType, fieldType, strconv.Quote(layout), expr,
		), nil
	},
}

// oneStringArgBefore returns an args func for a checker whose Go signature
// is func(param, value string) (string, error): the tag's whole params
// string is quoted and passed ahead of "value".
func oneStringArgBefore(string) func(string) ([]string, error) {
	return func(params string) ([]string, error) {
		return []string{strconv.Quote(params), "value"}, nil
	}
}

// twoStringArgsBefore returns an args func for a checker whose Go
// signature is func(a, b, value string) (string, error): the tag's params
// splits on ":" into two string literals, both passed ahead of "value".
func twoStringArgsBefore(checkerName string) func(string) ([]string, error) {
	return func(params string) ([]string, error) {
		a, b := splitParam(checkerName, params)
		return []string{strconv.Quote(a), strconv.Quote(b), "value"}, nil
	}
}

// oneNumberArgAfter returns an args func for a checker whose Go signature
// is func(value, n T) (T, error): "value" first, then the tag's params
// spliced in verbatim as an untyped numeric literal.
func oneNumberArgAfter(checkerName string) func(string) ([]string, error) {
	return func(params string) ([]string, error) {
		lit, err := numberLiteral(checkerName, params)
		if err != nil {
			return nil, err
		}

		return []string{"value", lit}, nil
	}
}

// oneNumberOrStringArgAfter is like oneNumberArgAfter, but for eq/ne, whose
// parameter may be a string, a bool, or a number depending on the field's
// type: a valid number is spliced in as a numeric literal (so it works for
// a numeric field), otherwise it's quoted as a string literal (so it works
// for a string field). Go rejects the mismatched case (e.g. a quoted
// string literal compared against an int field) at compile time, same as
// any other type error in generated code.
func oneNumberOrStringArgAfter(checkerName string) func(string) ([]string, error) {
	return func(params string) ([]string, error) {
		if lit, err := numberLiteral(checkerName, params); err == nil {
			return []string{"value", lit}, nil
		}

		return []string{"value", strconv.Quote(params)}, nil
	}
}

// oneOfArgs builds the arguments for IsOneOf(value T, allowed ...T): value
// first, then the tag's comma-separated allowed list, each element quoted.
// Numeric allowed values are unquoted instead, using the same
// number-or-string heuristic as oneNumberOrStringArgAfter.
func oneOfArgs(params string) ([]string, error) {
	args := []string{"value"}

	for _, allowed := range strings.Split(params, ",") {
		if lit, err := numberLiteral("oneof", allowed); err == nil {
			args = append(args, lit)
		} else {
			args = append(args, strconv.Quote(allowed))
		}
	}

	return args, nil
}

// intArg returns an arg func for a curried maker taking a single integer
// parameter (min-len, max-len, len).
func intArg(checkerName string) func(params, fieldType string) (string, error) {
	return func(params, _ string) (string, error) {
		return intLiteral(checkerName, params)
	}
}

// defaultArg builds the fallback argument for Default[T](fallback T): a
// quoted string literal for a string field, an unquoted (bool/numeric)
// literal otherwise.
func defaultArg(params, fieldType string) (string, error) {
	if fieldType == "string" {
		return strconv.Quote(params), nil
	}

	return params, nil
}
