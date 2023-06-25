package checker

import (
	"html"
	"reflect"
)

// NormalizerHTMLEscape is the name of the normalizer.
const NormalizerHTMLEscape = "html-escape"

// makeHTMLEscape makes a normalizer function for the HTML escape normalizer.
func makeHTMLEscape(_ string) CheckFunc {
	return normalizeHTMLEscape
}

// normalizeHTMLEscape applies HTML escaping to special characters.
// Uses html.EscapeString for the actual escape operation.
func normalizeHTMLEscape(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(html.EscapeString(value.String()))

	return ResultValid
}
