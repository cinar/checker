package checker

import "testing"

func TestIsMacInvalid(t *testing.T) {
	if IsMac("1234") == ResultValid {
		t.Fail()
	}
}

func TestIsMacValid(t *testing.T) {
	if IsMac("00:00:5e:00:53:01") != ResultValid {
		t.Fail()
	}
}

func TestCheckMacNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Network struct {
		HardwareAddress int `checkers:"mac"`
	}

	network := &Network{}

	Check(network)
}

func TestCheckMacInvalid(t *testing.T) {
	type Network struct {
		HardwareAddress string `checkers:"mac"`
	}

	network := &Network{
		HardwareAddress: "1234",
	}

	_, valid := Check(network)
	if valid {
		t.Fail()
	}
}

func TestCheckMacValid(t *testing.T) {
	type Network struct {
		HardwareAddress string `checkers:"mac"`
	}

	network := &Network{
		HardwareAddress: "00:00:5e:00:53:01",
	}

	_, valid := Check(network)
	if !valid {
		t.Fail()
	}
}
