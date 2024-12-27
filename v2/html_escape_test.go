// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestHTMLEscape(t *testing.T) {
	input := "<tag> \"Checker\" & 'Library' </tag>"
	expected := "&lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;"

	actual, err := v2.HTMLEscape(input)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestReflectHTMLEscape(t *testing.T) {
	type Comment struct {
		Body string `checkers:"html-escape"`
	}

	comment := &Comment{
		Body: "<tag> \"Checker\" & 'Library' </tag>",
	}

	expected := "&lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;"

	errs, ok := v2.CheckStruct(comment)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if comment.Body != expected {
		t.Fatalf("actual %s expected %s", comment.Body, expected)
	}
}
