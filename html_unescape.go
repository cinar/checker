package checker

import (
	"html"
	"reflect"
)

// NormalizerHTMLUnescape is the name of the normalizer.
const NormalizerHTMLUnescape = "html-unescape"

// makeHTMLUnescape makes a normalizer function for the HTML unscape normalizer.
func makeHTMLUnescape(_ string) CheckFunc {
	return normalizeHTMLUnescape
}

// normalizeHTMLUnescape applies HTML unescaping to special characters.
// Uses html.UnescapeString for the actual unescape operation.
func normalizeHTMLUnescape(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(html.UnescapeString(value.String()))

	return ResultValid
}
