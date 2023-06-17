package checker

import "testing"

func TestIsIpV4Invalid(t *testing.T) {
	if IsIpV4("900.800.200.100") == ResultValid {
		t.Fail()
	}
}

func TestIsIpV4InvalidV6(t *testing.T) {
	if IsIpV4("2001:db8::68") == ResultValid {
		t.Fail()
	}
}

func TestIsIpV4Valid(t *testing.T) {
	if IsIpV4("192.168.1.1") != ResultValid {
		t.Fail()
	}
}

func TestCheckIpV4NonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Request struct {
		RemoteIp int `checkers:"ipv4"`
	}

	request := &Request{}

	Check(request)
}

func TestCheckIpV4Invalid(t *testing.T) {
	type Request struct {
		RemoteIp string `checkers:"ipv4"`
	}

	request := &Request{
		RemoteIp: "900.800.200.100",
	}

	_, valid := Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIpV4Valid(t *testing.T) {
	type Request struct {
		RemoteIp string `checkers:"ipv4"`
	}

	request := &Request{
		RemoteIp: "192.168.1.1",
	}

	_, valid := Check(request)
	if !valid {
		t.Fail()
	}
}
