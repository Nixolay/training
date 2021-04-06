// Package numbersearch ...
package numbersearch

// FindSumTwoNumbers ...
func FindSumTwoNumbers(arr []int, number int) []int {
	vMap := make(map[int]struct{}, len(arr))
	for _, v := range arr {
		if _, ok := vMap[number-v]; ok {
			return []int{number - v, v}
		}

		vMap[v] = struct{}{}
	}

	return nil
}

// FindNumberInSortedArrays ...
func FindNumberInSortedArrays(x, y, z []uint) int {
	var nX, nY, nZ int

	for i := len(x) + len(y) + len(z); i > 0; i-- {
		if nX < len(x)-1 && x[nX] < y[nY] {
			nX++
		}

		if nY < len(y)-1 && y[nY] < z[nZ] {
			nY++
		}

		if nZ < len(z)-1 && z[nZ] < x[nX] {
			nZ++
		}

		if x[nX] == y[nY] && y[nY] == z[nZ] {
			return int(x[nX])
		}
	}

	return -1
}
