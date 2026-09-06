---
title: The Go Validator Blind Spot: Invisible Characters and Unicode Lookalikes
published: true
description: A username that reads "admin" on every screen you check it on can be storing something else entirely -- and it'll sail past a reserved-name check that's doing exactly what it was told. Most Go validation stacks have no way to catch it. Here's why, and two small normalizers that close the gap.
tags: go, golang, security, unicode
canonical_url: https://dev.to/onurcinar/the-go-validator-blind-spot-invisible-characters-and-unicode-lookalikes-38j5
---

Here's a username: `‮nimda‬`. Depending on how your browser renders bidirectional text, that might already display as `admin` — which is the entire point. Read it in a user list, a support ticket, or an audit log, and every person who looks at it will tell you the account is named `admin`. Now compare it to the literal string, the way a reserved-name check normally works:

```go
spoofed := "‮nimda‬" // renders as "admin" wherever bidi text is displayed

fmt.Println(spoofed == "admin") // false
fmt.Println(len([]rune(spoofed))) // 7, not 5
```

It isn't `"admin"`. It's the five letters `n`, `i`, `m`, `d`, `a` — in that literal order — wrapped in two invisible bidirectional control characters, U+202E (right-to-left override) and U+202C (pop directional formatting). Unicode reserves those for rendering Arabic and Hebrew correctly. Wrapped around plain Latin letters instead, they just flip the visual order: `nimda`, laid out right-to-left, reads `admin`.

A reserved-username check that rejects the exact string `"admin"` lets this straight through — correctly, by its own logic, since the string genuinely isn't `"admin"`. Nothing here is a bug in `checker`, or in any other Go validation library, or even in Unicode. The check did exactly what it was told. The problem is one level up: it was never asked whether the string secretly contains characters that don't correspond to anything the comparison sees but everyone looking at a screen does.

## This is not a new attack

The trick behind that username has a name: in November 2021, researchers at Cambridge published [Trojan Source (CVE-2021-42574)](https://trojansource.codes/), showing that the same bidirectional control characters can be slipped into source code to make it *compile* one way and *display* another way to a human reviewer. The characters are identical; only the target changes. Anything a person reads back instead of a compiler — a username, a chat message, a support ticket, a PR title — is just as susceptible.

A narrower, older trick uses zero-width characters instead of bidi ones: insert a zero-width space or zero-width joiner into a blocklisted word and a naive filter no longer matches it, even though the word displays unchanged. It's a standard way spam and moderation filters get evaded, and it works against a validation struct tag exactly as well as it works against a regex.

## The other half of the problem: it doesn't need to be invisible at all

Zero-width characters are the sharp case, but plain Unicode gives you the same problem with characters you *can* see. `"ALICE"`, `"ＡＬＩＣＥ"` (fullwidth Latin letters), and a version using ligatures or Roman numerals in place of ordinary letters are three different byte sequences that a human reading them would call the same word. A uniqueness check, a keyword filter, or a duplicate-handle guard that only compares strings byte-for-byte treats all three as unrelated.

This isn't hypothetical, either — fullwidth-character substitution is a known way to slip a blocked command past a keyword-based filter (`ｒｍ　－ｒｆ　／` reads as `rm -rf /` to anyone who looks at it, and normalizes right back to it), and several platforms have had to retrofit Unicode normalization after discovering it let people register visually-identical duplicate usernames.

None of this is because validation is careless. `required`, `min-len`, an exact-match reserved-name rule — every one of them is a byte-level check, and byte-level checks were never asked a different question: does this string contain characters that don't correspond to anything the person reading it actually typed? That's a normalization problem, not a validation one. It needs its own step in the pipeline, before anything else looks at the value.

## Two checkers, two different problems

`checker` added two normalizers for this, deliberately kept separate because they solve different problems and have different costs.

**`strip-invisible`** (core module, no dependency) removes zero-width space, zero-width non-joiner, zero-width joiner, word joiner, the zero-width no-break space (BOM), and the bidirectional embedding/override/isolate controls behind Trojan Source:

```go
type Handle struct {
	Name string `checkers:"trim strip-invisible required min-len:3"`
}

h := &Handle{Name: "adm‌in"}
checker.CheckStruct(h)
// h.Name is now "admin" -- five runes, matches the literal string.

unmasked, _ := checker.StripInvisible("‮nimda‬")
fmt.Println(unmasked)          // nimda
fmt.Println(unmasked == "admin") // false -- and now it doesn't even display like it
```

Run the username from the opening through it and the disguise falls apart: `strip-invisible` doesn't know or care what the override characters were trying to accomplish, it just removes them — leaving the honest `nimda`, which no longer reads as `admin` in any renderer.

**`nfkc`** (an opt-in module — it needs `golang.org/x/text/unicode/norm`, so it's kept out of the dependency-free core) applies [Unicode Normalization Form KC](https://unicode.org/reports/tr15/): it folds compatibility characters — fullwidth forms, ligatures, and similar stylistic variants — into their canonical equivalents:

```go
import _ "github.com/cinar/checker/v2/nfkc"

type Handle struct {
	Name string `checkers:"trim nfkc required min-len:3"`
}

a := &Handle{Name: "ＡＬＩＣＥ"}
checker.CheckStruct(a)
// a.Name is now "ALICE"

waf, _ := nfkc.Normalize("ｒｍ　－ｒｆ　／")
// waf is "rm -rf /"
```

Both are one word in a `checkers` tag. Neither needs a second pass over your validation logic — they sit in the same pipeline as `trim`, `lower`, and `required`, because normalizers and checkers are the same function type in this library and can be freely mixed.

## What this does *not* fix

Be precise about what NFKC actually covers, because it's narrower than "Unicode spoofing" in general. NFKC only folds characters that are *compatibility-equivalent* under the Unicode standard — the same letter, rendered differently. It does nothing for **homoglyphs**: characters that are visually similar but come from a different script entirely, like a Cyrillic `а` (U+0430) standing in for a Latin `a` (U+0061). Those two are not NFKC-equivalent, so `paypal.com` spelled with a Cyrillic `а` normalizes right back to itself — still spoofed.

Catching that class of lookalike needs a *confusables* check — [Unicode UTS #39](https://unicode.org/reports/tr39/)'s skeleton algorithm, which maps visually similar characters across scripts to a shared representative form before comparing. It's a genuinely different, heavier piece of machinery than either normalizer here, and `checker` doesn't currently provide it. If cross-script impersonation is part of your threat model — a login handle, a domain-like identifier — `strip-invisible` plus `nfkc` narrows the problem but doesn't close it.

## Where to apply this, and where not to

Both normalizers are opinionated about *removing or rewriting* characters, which means they're right for the fields where an invisible or compatibility character is never legitimate — a username, a handle, a search keyword, an API key — and wrong for general free-text content. Zero-width joiner is load-bearing in emoji sequences; zero-width non-joiner has real uses in Persian and other scripts. Running `strip-invisible` over a chat message body or a comment field will quietly corrupt text some of your users actually rely on. Point these at identifiers, not prose.

```go
type Registration struct {
	Handle string `checkers:"trim strip-invisible nfkc required min-len:3 alphanumeric"`
	Bio    string `checkers:"trim max-len:500"` // free text -- neither normalizer applies here
}
```

`strip-invisible` ships in [Checker](https://github.com/cinar/checker)'s core module; `nfkc` is a separate, opt-in module at [`github.com/cinar/checker/v2/nfkc`](https://github.com/cinar/checker/tree/main/nfkc), so the dependency it needs never touches a codebase that doesn't ask for it.
