package checker

import "testing"

func TestIsIPV6Invalid(t *testing.T) {
	if IsIPV6("900.800.200.100") == ResultValid {
		t.Fail()
	}
}

func TestIsIPV6InvalidV4(t *testing.T) {
	if IsIPV6("192.168.1.1") == ResultValid {
		t.Fail()
	}
}

func TestIsIPV6Valid(t *testing.T) {
	if IsIPV6("2001:db8::68") != ResultValid {
		t.Fail()
	}
}

func TestCheckIPV6NonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Request struct {
		RemoteIP int `checkers:"ipv6"`
	}

	request := &Request{}

	Check(request)
}

func TestCheckIPV6Invalid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ipv6"`
	}

	request := &Request{
		RemoteIP: "900.800.200.100",
	}

	_, valid := Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIPV6Valid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ipv6"`
	}

	request := &Request{
		RemoteIP: "2001:db8::68",
	}

	_, valid := Check(request)
	if !valid {
		t.Fail()
	}
}
