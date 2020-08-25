package ipvalidator

import (
	"net"
	"strconv"
	"strings"
)

//nolint
func IsValidIP(ip string) bool {
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

// nolint
func Is_valid_ip(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}

	return true
}
