// Package v2 is a minimal stand-in for github.com/cinar/checker/v2, just
// enough for checkerlint's own tests to resolve an import of it and detect
// calls to RegisterMaker/RegisterFieldMaker, without pulling in the real
// module (analysistest resolves testdata packages in GOPATH mode, keyed by
// import path, not by module identity).
package v2

// RegisterMaker stands in for the real checker.RegisterMaker.
func RegisterMaker(name string, maker interface{}) {}

// RegisterFieldMaker stands in for the real checker.RegisterFieldMaker.
func RegisterFieldMaker(name string, maker interface{}) {}

// OtherFunc is an unrelated exported function in this package, used to
// exercise the "right package, wrong function name" path in
// recordCustomRegistration.
func OtherFunc(name string) {}

// Named is a stand-in for an exported struct type from this package, used
// to exercise an embedded field qualified by a package selector.
type Named struct{}
