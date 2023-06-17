package checker

import "testing"

func TestIsIpV6Invalid(t *testing.T) {
	if IsIpV6("900.800.200.100") == ResultValid {
		t.Fail()
	}
}

func TestIsIpV6Valid(t *testing.T) {
	if IsIpV6("2001:db8::68") != ResultValid {
		t.Fail()
	}
}

func TestCheckIpV6NonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Request struct {
		RemoteIp int `checkers:"ipv6"`
	}

	request := &Request{}

	Check(request)
}

func TestCheckIpV6Invalid(t *testing.T) {
	type Request struct {
		RemoteIp string `checkers:"ipv6"`
	}

	request := &Request{
		RemoteIp: "900.800.200.100",
	}

	_, valid := Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIpV6Valid(t *testing.T) {
	type Request struct {
		RemoteIp string `checkers:"ipv6"`
	}

	request := &Request{
		RemoteIp: "2001:db8::68",
	}

	_, valid := Check(request)
	if !valid {
		t.Fail()
	}
}
