package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
	Учитывая целочисленный массив nums, найдите подмассив с
	наибольшей суммой и возвращает ее сумму.

Пример 1:
	Ввод: числа = [-2,1,-3,4,-1,2,1,-5,4]
	Выходные данные:6
	Объяснение: Подмассив [4,-1,2,1] имеет наибольшую сумму 6.
Пример 2:
	Входные данные: nums = [1]
	Вывод:1
	Объяснение: Подмассив [1] имеет наибольшую сумму 1.
Пример 3:
	Входные данные: числа = [5,4,-1,7,8]
	Вывод:23
	Объяснение: Подмассив [5,4,-1,7,8] имеет наибольшую сумму 23.

Ограничения:
	1 <= nums.length <= 105
	-104 <= nums[i] <= 104


Продолжение:
	Если вы разобрались с O(n) решением, попробуйте написать другое решение,
	используя подход разделяй и властвуй, который является более тонким.
*/

func TestMaximumSubarray(t *testing.T) {
	t.Parallel()

	type testData struct {
		nums     []int
		expected int
	}

	data := []testData{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, // [4,-1,2,1]
		{[]int{1}, 1},               // [1]
		{[]int{5, 4, -1, 7, 8}, 23}, // [5,4,-1,7,8]
	}

	for _, v := range data {
		require.Equal(t, v.expected, maxSubArray(v.nums))
	}
}

func maxSubArray(nums []int) int {
	currentMax, globalMax := nums[0], nums[0]

	for iter := range nums {
		if currentMax+nums[iter] > nums[iter] && iter > 0 {
			currentMax += nums[iter] // Обновляем текущую максимальную сумму
		} else {
			currentMax = nums[iter] // либо начинаем новый подмассив с текущего элемента
		}

		if currentMax > globalMax {
			globalMax = currentMax // Обновляем глобальную максимальную сумму
		}
	}

	return globalMax
}
