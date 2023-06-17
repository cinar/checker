package checker

import (
	"net"
	"reflect"
)

// CheckerIpV4 is the name of the checker.
const CheckerIpV4 = "ipv4"

// ResultNotIpV4 indicates that the given value is not an IPv4 address.
const ResultNotIpV4 = "NOT_IP_V4"

// IsIpV4 checks if the given value is an IPv4 address.
func IsIpV4(value string) Result {
	ip := net.ParseIP(value)
	if ip == nil {
		return ResultNotIpV4
	}

	if ip.To4() == nil {
		return ResultNotIpV4
	}

	return ResultValid
}

// makeIpV4 makes a checker function for the ipV4 checker.
func makeIpV4(_ string) CheckFunc {
	return checkIpV4
}

// checkIpV4 checks if the given value is an IPv4 address.
func checkIpV4(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIpV4(value.String())
}
