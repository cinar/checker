package checker

import "testing"

func TestIsIPInvalid(t *testing.T) {
	if IsIP("900.800.200.100") == ResultValid {
		t.Fail()
	}
}

func TestIsIPValid(t *testing.T) {
	if IsIP("2001:db8::68") != ResultValid {
		t.Fail()
	}
}

func TestCheckIpNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Request struct {
		RemoteIP int `checkers:"ip"`
	}

	request := &Request{}

	Check(request)
}

func TestCheckIpInvalid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ip"`
	}

	request := &Request{
		RemoteIP: "900.800.200.100",
	}

	_, valid := Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIPValid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ip"`
	}

	request := &Request{
		RemoteIP: "192.168.1.1",
	}

	_, valid := Check(request)
	if !valid {
		t.Fail()
	}
}
