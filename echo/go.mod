module github.com/cinar/checker/v2/echo

go 1.25.0

require (
	github.com/cinar/checker/v2 v2.0.1
	github.com/labstack/echo/v4 v4.15.4
)

require (
	github.com/labstack/gommon v0.5.0 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.38.0 // indirect
)

// Always build against the sibling checker module in this repository during
// development and CI. This has no effect on external consumers of this
// module: go get resolves the require directive above via the module proxy,
// since replace directives in a dependency's own go.mod are ignored.
replace github.com/cinar/checker/v2 => ../
