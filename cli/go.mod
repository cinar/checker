module github.com/cinar/checker/v2/cli

go 1.25.0

require github.com/cinar/checker/v2 v2.0.1

// Always build against the sibling checker module in this repository during
// development and CI. This has no effect on external consumers of this
// module: go get resolves the require directive above via the module proxy,
// since replace directives in a dependency's own go.mod are ignored.
replace github.com/cinar/checker/v2 => ../
