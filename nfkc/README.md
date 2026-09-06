# Checker NFKC Normalizer

This package registers an `nfkc` [Checker](https://github.com/cinar/checker) checkers-tag normalizer applying [Unicode Normalization Form KC](https://unicode.org/reports/tr15/) (NFKC): it composes combining character sequences and replaces compatibility characters — fullwidth digits, ligatures, and many other stylistic variants a naive keyword filter, uniqueness check, or handle comparison would otherwise treat as distinct — with their canonical equivalents.

It's a separate module, not part of the core `checker` package, because NFKC normalization needs [`golang.org/x/text/unicode/norm`](https://pkg.go.dev/golang.org/x/text/unicode/norm). Keeping it isolated means the core module's zero-dependency promise holds regardless of whether a caller opts into this normalizer — the same reasoning behind the `gin`/`echo`/`nethttp`/`fiber` adapter modules, just for a checker instead of a framework binding.

For zero-width and bidirectional spoofing characters (the "Trojan Source" technique), see the core module's [`strip-invisible`](../README.md#normalizers-provided) normalizer instead — that one needs no external dependency and ships in the core module directly.

## Install

```bash
go get github.com/cinar/checker/v2/nfkc
```

## Usage

Importing this package registers the `nfkc` normalizer through [`RegisterMaker`](https://pkg.go.dev/github.com/cinar/checker/v2#RegisterMaker) in its `init` function — a blank import is enough if you only use it through a struct tag:

```golang
import (
	checker "github.com/cinar/checker/v2"
	_ "github.com/cinar/checker/v2/nfkc"
)

type Handle struct {
	Name string `checkers:"trim nfkc required min-len:3"`
}

handle := &Handle{
	Name: "ＡＬＩＣＥ",
}

errs, ok := checker.CheckStruct(handle)
// handle.Name is now "ALICE"
```

Call `Normalize` directly for a one-off value instead of a struct field:

```golang
normalized, _ := nfkc.Normalize("１２３")
// normalized is "123"
```

## License

This library is free to use, modify, and distribute under the terms of the MIT license found in the [LICENSE](../LICENSE) file of the parent [Checker](https://github.com/cinar/checker) repository.
