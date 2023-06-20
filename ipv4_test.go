package checker

import "testing"

func TestIsIPV4Invalid(t *testing.T) {
	if IsIPV4("900.800.200.100") == ResultValid {
		t.Fail()
	}
}

func TestIsIPV4InvalidV6(t *testing.T) {
	if IsIPV4("2001:db8::68") == ResultValid {
		t.Fail()
	}
}

func TestIsIPV4Valid(t *testing.T) {
	if IsIPV4("192.168.1.1") != ResultValid {
		t.Fail()
	}
}

func TestCheckIPV4NonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Request struct {
		RemoteIP int `checkers:"ipv4"`
	}

	request := &Request{}

	Check(request)
}

func TestCheckIPV4Invalid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ipv4"`
	}

	request := &Request{
		RemoteIP: "900.800.200.100",
	}

	_, valid := Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIPV4Valid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ipv4"`
	}

	request := &Request{
		RemoteIP: "192.168.1.1",
	}

	_, valid := Check(request)
	if !valid {
		t.Fail()
	}
}
