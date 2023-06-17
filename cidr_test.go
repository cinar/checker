package checker

import "testing"

func TestIsCidrInvalid(t *testing.T) {
	if IsCidr("900.800.200.100//24") == ResultValid {
		t.Fail()
	}
}

func TestIsCidrValid(t *testing.T) {
	if IsCidr("2001:db8::/32") != ResultValid {
		t.Fail()
	}
}

func TestCheckCidrNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Network struct {
		Subnet int `checkers:"cidr"`
	}

	network := &Network{}

	Check(network)
}

func TestCheckCidrInvalid(t *testing.T) {
	type Network struct {
		Subnet string `checkers:"cidr"`
	}

	network := &Network{
		Subnet: "900.800.200.100//24",
	}

	_, valid := Check(network)
	if valid {
		t.Fail()
	}
}

func TestCheckCidrValid(t *testing.T) {
	type Network struct {
		Subnet string `checkers:"cidr"`
	}

	network := &Network{
		Subnet: "192.0.2.0/24",
	}

	_, valid := Check(network)
	if !valid {
		t.Fail()
	}
}
