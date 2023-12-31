package checker

import (
	"reflect"
	"regexp"
	"strings"
)

// CheckerEmail is the name of the checker.
const CheckerEmail = "email"

// ResultNotEmail indicates that the given string is not a valid email.
const ResultNotEmail = "NOT_EMAIL"

// ipV6Prefix is the IPv6 prefix for the domain.
const ipV6Prefix = "[IPv6:"

// notQuotedChars is the valid not quoted characters.
var notQuotedChars = regexp.MustCompile("[a-zA-Z0-9!#$%&'*\\+\\-/=?^_`{|}~]")

// IsEmail checks if the given string is an email address.
func IsEmail(email string) Result {
	atIndex := strings.LastIndex(email, "@")
	if atIndex == -1 || atIndex == len(email)-1 {
		return ResultNotEmail
	}

	domain := email[atIndex+1:]
	if isValidEmailDomain(domain) != ResultValid {
		return ResultNotEmail
	}

	return isValidEmailUser(email[:atIndex])
}

// makeEmail makes a checker function for the email checker.
func makeEmail(_ string) CheckFunc {
	return checkEmail
}

// checkEmail checks if the given string is an email address.
func checkEmail(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsEmail(value.String())
}

// isValidEmailDomain checks if the email domain is a IPv4 or IPv6 address, or a FQDN.
func isValidEmailDomain(domain string) Result {
	if len(domain) > 255 {
		return ResultNotEmail
	}

	if domain[0] == '[' {
		if strings.HasPrefix(domain, ipV6Prefix) {
			// postmaster@[IPv6:2001:0db8:85a3:0000:0000:8a2e:0370:7334]
			return IsIPV6(domain[len(ipV6Prefix) : len(domain)-1])
		}

		// postmaster@[123.123.123.123]
		return IsIPV4(domain[1 : len(domain)-1])
	}

	return IsFqdn(domain)
}

// isValidEmailUser checks if the email user is valid.
func isValidEmailUser(user string) Result {
	// Cannot be empty user
	if user == "" || len(user) > 64 {
		return ResultNotEmail
	}

	// Cannot start or end with dot
	if user[0] == '.' || user[len(user)-1] == '.' {
		return ResultNotEmail
	}

	return isValidEmailUserCharacters(user)
}

// isValidEmailUserCharacters if the email user characters are valid.
func isValidEmailUserCharacters(user string) Result {
	quoted := false
	start := true
	prev := ' '

	for _, c := range user {
		// Cannot have a double dot unless quoted
		if !quoted && c == '.' && prev == '.' {
			return ResultNotEmail
		}

		if start {
			start = false

			if c == '"' {
				quoted = true
				prev = c
				continue
			}
		}

		if !quoted {
			if c == '.' {
				start = true
			} else if !notQuotedChars.MatchString(string(c)) {
				return ResultNotEmail
			}
		} else {
			if c == '"' && prev != '\\' {
				quoted = false
			}
		}

		prev = c
	}

	return ResultValid
}
