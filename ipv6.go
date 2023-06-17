package checker

import (
	"net"
	"reflect"
)

// CheckerIpV6 is the name of the checker.
const CheckerIpV6 = "ipv6"

// ResultNotIpV6 indicates that the given value is not an IPv6 address.
const ResultNotIpV6 = "NOT_IP_V6"

// IsIpV6 checks if the given value is an IPv6 address.
func IsIpV6(value string) Result {
	ip := net.ParseIP(value)
	if ip == nil {
		return ResultNotIpV6
	}

	if ip.To4() != nil {
		return ResultNotIpV6
	}

	return ResultValid
}

// makeIpV6 makes a checker function for the ipV6 checker.
func makeIpV6(_ string) CheckFunc {
	return checkIpV6
}

// checkIpV6 checks if the given value is an IPv6 address.
func checkIpV6(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIpV6(value.String())
}
