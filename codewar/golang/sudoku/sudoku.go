/*
Судоку Фон

Судоку - игра, в которую играют по сетке 9х9.
Цель игры - заполнить все ячейки сетки цифрами от 1 до 9,
чтобы каждый столбец, каждая строка и каждая из девяти подрешеток 3x3
(также называемых блоками) содержали все цифры от 1 до 9
(Более подробная информация на: http://en.wikipedia.org/wiki/Sudoku)
Судоку Решение Валидатор

Напишите функцию validSolution/ValidateSolution/valid_solution(),
которая принимает двумерный массив, представляющий доску судоку,
и возвращает true, если оно является допустимым решением, или false
в противном случае.
Ячейки доски судоку также могут содержать нули,
которые будут представлять пустые ячейки.
Доски, содержащие один или несколько нулей,
считаются недействительными решениями.

Доска всегда 9 ячеек на 9 ячеек, и каждая ячейка содержит только целые числа от 0 до 9.
*/
package sudoku

import (
	"fmt"
	"sort"
)

func ValidateSolution(m [][]int) bool {
	for i := range m {
		data := make([]int, 9)
		copy(data, m[i])

		sort.Ints(data)
		if fmt.Sprint(data) != "[1 2 3 4 5 6 7 8 9]" {
			return false
		}

		for j := 0; j < 9; j++ {
			data[j] = m[j][i]
		}

		sort.Ints(data)
		if fmt.Sprint(data) != "[1 2 3 4 5 6 7 8 9]" {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			data := append([]int{}, m[i][j:j+3]...)
			data = append(data, m[i+1][j:j+3]...)
			data = append(data, m[i+2][j:j+3]...)

			sort.Ints(data)
			if fmt.Sprint(data) != "[1 2 3 4 5 6 7 8 9]" {
				return false
			}
		}
	}

	return true
}
