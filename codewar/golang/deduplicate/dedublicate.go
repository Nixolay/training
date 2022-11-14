package deduplicate

func Deduplicate(data []int) []int {
	arr := make([]int, 0, len(data))

	for i := range data {
		if i != 0 && data[i-1] == data[i] {
			continue
		}

		arr = append(arr, data[i])
	}

	return arr
}
