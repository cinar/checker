// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestNormalizeHTMLEscapeNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)
	type Comment struct {
		Body int `checkers:"html-escape"`
	}

	comment := &Comment{}

	checker.Check(comment)
}

func TestNormalizeHTMLEscape(t *testing.T) {
	type Comment struct {
		Body string `checkers:"html-escape"`
	}

	comment := &Comment{
		Body: "<tag> \"Checker\" & 'Library' </tag>",
	}

	_, valid := checker.Check(comment)
	if !valid {
		t.Fail()
	}

	if comment.Body != "&lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;" {
		t.Fail()
	}
}
