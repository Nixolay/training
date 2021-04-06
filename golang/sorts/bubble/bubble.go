// Package bubble ...
package bubble

// Sort осуществляет сортировку чисел пузырьком.
func Sort(numbers []int) {
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}
}
