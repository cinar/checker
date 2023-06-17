package checker

import (
	"net"
	"reflect"
)

// CheckerIp is the name of the checker.
const CheckerIp = "ip"

// ResultNotIp indicates that the given value is not an IP address.
const ResultNotIp = "NOT_IP"

// IsIp checks if the given value is an IP address.
func IsIp(value string) Result {
	ip := net.ParseIP(value)
	if ip == nil {
		return ResultNotIp
	}

	return ResultValid
}

// makeIp makes a checker function for the ip checker.
func makeIp(_ string) CheckFunc {
	return checkIp
}

// checkIp checks if the given value is an IP address.
func checkIp(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIp(value.String())
}
