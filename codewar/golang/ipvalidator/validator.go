// Package ipvalidator checked valid ip
package ipvalidator

import (
	"net"
	"strconv"
	"strings"
)

//nolint:gomnd
// IsValidIP checked valid ip.
func IsValidIP2(ip string) bool {
	numbers, four := strings.Split(ip, "."), 4
	if len(numbers) < four {
		return false
	}

	for _, v := range numbers {
		number, err := strconv.Atoi(v)
		if err != nil || number > 255 || number < 0 {
			return false
		}

		if len(v) < 2 {
			continue
		}

		if v[0] == '0' {
			return false
		}
	}

	return true
}

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
