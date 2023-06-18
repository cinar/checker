package checker

import "testing"

func TestCheckFdqnWithoutTld(t *testing.T) {
	if IsFqdn("abcd") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnShortTld(t *testing.T) {
	if IsFqdn("abcd.c") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnNumericTld(t *testing.T) {
	if IsFqdn("abcd.1234") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnLong(t *testing.T) {
	if IsFqdn("abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd.com") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnInvalidCharacters(t *testing.T) {
	if IsFqdn("ab_cd.com") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnStaringWithHyphen(t *testing.T) {
	if IsFqdn("-abcd.com") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnStaringEndingWithHyphen(t *testing.T) {
	if IsFqdn("abcd-.com") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnStartingWithDot(t *testing.T) {
	if IsFqdn(".abcd.com") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFdqnEndingWithDot(t *testing.T) {
	if IsFqdn("abcd.com.") != ResultNotFqdn {
		t.Fail()
	}
}

func TestCheckFqdnNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Request struct {
		Domain int `checkers:"fqdn"`
	}

	request := &Request{}

	Check(request)
}

func TestCheckFqdnValid(t *testing.T) {
	type Request struct {
		Domain string `checkers:"fqdn"`
	}

	request := &Request{
		Domain: "zdo.com",
	}

	_, valid := Check(request)
	if !valid {
		t.Fail()
	}
}
