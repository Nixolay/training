package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
https://leetcode.com/problems/trapping-rain-water/description/
Учитывая n неотрицательные целые числа, представляющие карту высот,
где ширина каждой полосы равна 1, вычислите, сколько воды она может улавливать после дождя.

Пример 1:
	Ввод: height = [0,1,0,2,1,0,1,3,2,1,2,1]
	Выход:6
	Пояснение:
		Приведенная выше карта высот (черная секция) представлена массивом [0,1,0,2,1,0,1,3,2,1,2,1].
		В этом случае улавливается 6 единиц дождевой воды (синяя секция).

	Пример 2:
	Ввод: height = [4,2,0,3,2,5]
	Результат: 9


Ограничения:
	n == height.length
	1 <= n <= 2 * 104
	0 <= height[i] <= 105
*/

func TestTrappingRainWater(t *testing.T) {
	t.Parallel()

	type testData struct {
		nums     []int
		expected int
	}

	data := []testData{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	for _, v := range data {
		require.Equal(t, v.expected, trap(v.nums))
	}
}

/*
	Для решения задачи, где требуется вычислить количество воды, задерживаемой между столбцами,
представляющими карту высот, можно использовать подход двух указателей.
Этот подход эффективен и работает за время O(n), где n - количество элементов в массиве heights.

Основная идея заключается в следующем:
1. Используем два указателя, left и right, которые начинаются с краев массива.
2. Заводим переменные left_max и right_max для отслеживания максимальной высоты
	на левой и правой сторонах от текущего положения указателей.
3. Идем по массиву с помощью указателей:
	- Если height[left] < height[right], это означает,
		что текущий элемент слева может быть ограничителем для воды справа,
		поэтому смотрим, сколько воды можно задержать между left и left_max.
	- Аналогично, если height[left] >= height[right],
		текущий элемент справа может быть ограничителем для воды слева.
4.Перемещаем указатели в направлении увеличения максимальной высоты
	(left_max или right_max), пока не пройдем весь массив.
*/

func trap(height []int) int {
	left, right := 0, len(height)-1
	left_max, right_max := 0, 0
	water := 0
	for left <= right {
		if height[left] < height[right] {
			if height[left] >= left_max {
				left_max = height[left]
			} else {
				water += left_max - height[left]
			}
			left++
		} else {
			if height[right] >= right_max {
				right_max = height[right]
			} else {
				water += right_max - height[right]
			}
			right--
		}
	}
	return water
}
