package numbersearch_test

import (
	"reflect"
	"testing"

	search "github.com/Nixolay/training/golang/array_operations/number_search"
)

func TestFindSumTwoNumbers(t *testing.T) {
	data := []struct {
		arr      []int
		k        int
		expected []int
	}{
		{[]int{2, 5, 7, 3, 4}, 6, []int{2, 4}},
		{[]int{1, 3, 4, 5, 6}, 2, nil},
		{[]int{3, 5, 6}, 3, nil},
	}

	for _, v := range data {
		actual := search.FindSumTwoNumbers(v.arr, v.k)
		if !reflect.DeepEqual(v.expected, actual) {
			t.Fatalf("expected: %v equals actual: %v", v.expected, actual)
		}
	}
}

func TestFindNumberInSortedArrays(t *testing.T) {
	data := []struct {
		x, y, z  []uint
		expected int
	}{
		{
			[]uint{0, 2, 4, 5, 6},
			[]uint{1, 6, 10, 123},
			[]uint{0, 3, 5, 6, 7, 8, 9},
			6,
		},
		{
			[]uint{0, 2, 4, 5, 6},
			[]uint{1, 5, 10, 123},
			[]uint{0, 3, 5, 7, 8, 9},
			5,
		},
		{
			[]uint{0, 2, 4, 5},
			[]uint{1, 6, 10, 123},
			[]uint{0, 3, 5, 6, 7, 8, 9},
			-1,
		},
		{
			[]uint{0, 1, 2, 3, 4, 5},
			[]uint{6, 7, 8, 9, 10, 11},
			[]uint{12, 13, 14, 15, 16, 17, 18, 19, 20},
			-1,
		},
	}

	for _, v := range data {
		if actual := search.FindNumberInSortedArrays(v.x, v.y, v.z); actual != v.expected {
			t.Fatalf("expected: %v equals actual: %v", v.expected, actual)
		}
	}
}
