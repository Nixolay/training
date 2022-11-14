package numberofunits

import "strconv"

func LongestSequenceOfUnits(arr []int64) int {
	item := 0
	max := 0

	for i, n := range arr {
		count := 0

		for _, e := range strconv.FormatInt(n, 2) {
			print(string(e))
			if e == '1' {
				count++
			}
		}
		println()

		if count > max {
			max = count
			item = i
		}
	}
	println("-----------------")

	return item
}
