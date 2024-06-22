package leetcode

/*
Учитывая массив целых чисел nums и целое число target, верните индексы двух чисел таким образом, чтобы они в сумме равнялись target.
Вы можете предположить, что каждый ввод будет иметь ровно одно решение, и вы не можете использовать один и тот же элемент дважды.
Вы можете вернуть ответ в любом порядке.

Пример 1:
Входные данные: числа = [2,7,11,15], цель = 9
Выходные данные: [0,1]
Объяснение: Поскольку nums[0] + nums[1] == 9, мы возвращаем [0, 1].

Пример 2:
Входные данные: числа = [3,2,4], цель = 6
Выходные данные: [1,2]

Пример 3:
Входные данные: числа = [3,3], цель = 6
Выходные данные: [0,1]
*/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	type testData struct {
		nums     []int
		target   int
		expected []int
	}

	data := []testData{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, v := range data {
		require.Equal(t, v.expected, TwoSum(v.nums, v.target))
	}
}

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
