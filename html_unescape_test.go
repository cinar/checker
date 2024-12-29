// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestHTMLUnescape(t *testing.T) {
	input := "&lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;"
	expected := "<tag> \"Checker\" & 'Library' </tag>"

	actual, err := v2.HTMLUnescape(input)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestReflectHTMLUnescape(t *testing.T) {
	type Comment struct {
		Body string `checkers:"html-unescape"`
	}

	comment := &Comment{
		Body: "&lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;",
	}

	expected := "<tag> \"Checker\" & 'Library' </tag>"

	errs, ok := v2.CheckStruct(comment)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if comment.Body != expected {
		t.Fatalf("actual %s expected %s", comment.Body, expected)
	}
}
