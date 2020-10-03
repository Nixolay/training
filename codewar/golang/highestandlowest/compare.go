package compare

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	low, big := 0, 0

	for i, n := range strings.Split(in, " ") {
		number, _ := strconv.Atoi(n)

		if i == 0 {
			low, big = number, number

			continue
		}

		if number > big {
			big = number
		}

		if number < low {
			low = number
		}
	}

	return fmt.Sprintf("%d %d", big, low)
}
