package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestNormalizeHTMLUnescapeNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Comment struct {
		Body int `checkers:"html-unescape"`
	}

	comment := &Comment{}

	checker.Check(comment)
}

func TestNormalizeHTMLUnescape(t *testing.T) {
	type Comment struct {
		Body string `checkers:"html-unescape"`
	}

	comment := &Comment{
		Body: "&lt;tag&gt; &#34;Checker&#34; &amp; &#39;Library&#39; &lt;/tag&gt;",
	}

	_, valid := checker.Check(comment)
	if !valid {
		t.Fail()
	}

	if comment.Body != "<tag> \"Checker\" & 'Library' </tag>" {
		t.Fail()
	}
}
