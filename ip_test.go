package checker

import "testing"

func TestIsIpInvalid(t *testing.T) {
	if IsIp("900.800.200.100") == ResultValid {
		t.Fail()
	}
}

func TestIsIpValid(t *testing.T) {
	if IsIp("2001:db8::68") != ResultValid {
		t.Fail()
	}
}

func TestCheckIpNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Request struct {
		RemoteIp int `checkers:"ip"`
	}

	request := &Request{}

	Check(request)
}

func TestCheckIpInvalid(t *testing.T) {
	type Request struct {
		RemoteIp string `checkers:"ip"`
	}

	request := &Request{
		RemoteIp: "900.800.200.100",
	}

	_, valid := Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIpValid(t *testing.T) {
	type Request struct {
		RemoteIp string `checkers:"ip"`
	}

	request := &Request{
		RemoteIp: "192.168.1.1",
	}

	_, valid := Check(request)
	if !valid {
		t.Fail()
	}
}
