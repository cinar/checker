package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestCheckFdqnWithoutTld(t *testing.T) {
	if checker.IsFqdn("abcd") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnShortTld(t *testing.T) {
	if checker.IsFqdn("abcd.c") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnNumericTld(t *testing.T) {
	if checker.IsFqdn("abcd.1234") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnLong(t *testing.T) {
	if checker.IsFqdn("abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd.com") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnInvalidCharacters(t *testing.T) {
	if checker.IsFqdn("ab_cd.com") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnStaringWithHyphen(t *testing.T) {
	if checker.IsFqdn("-abcd.com") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnStaringEndingWithHyphen(t *testing.T) {
	if checker.IsFqdn("abcd-.com") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnStartingWithDot(t *testing.T) {
	if checker.IsFqdn(".abcd.com") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnEndingWithDot(t *testing.T) {
	if checker.IsFqdn("abcd.com.") != checker.ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFqdnNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Request struct {
		Domain int `checkers:"fqdn"`
	}

	request := &Request{}

	checker.Check(request)
}

func TestCheckFqdnValid(t *testing.T) {
	type Request struct {
		Domain string `checkers:"fqdn"`
	}

	request := &Request{
		Domain: "zdo.com",
	}

	_, valid := checker.Check(request)
	if !valid {
		t.Fail()
	}
}
