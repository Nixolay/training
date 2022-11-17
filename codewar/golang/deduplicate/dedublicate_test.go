package deduplicate_test

import (
	"testing"

	dd "github.com/Nixolay/training/codewar/golang/deduplicate"
	"github.com/stretchr/testify/require"
)

/*
Дан упорядоченный по неубыванию массив целых 32-разрядных чисел. Требуется удалить из него все повторения.

Желательно получить решение, которое не считывает входной файл целиком в память, т.е., использует лишь константный объем памяти в процессе работы.
*/

func TestDeduplicate(t *testing.T) {
	data := []struct {
		d []int
		r []int
	}{
		{
			d: []int{5, 2, 4, 8, 8, 8},
			r: []int{5, 2, 4, 8},
		},
		{
			d: []int{5, 2, 2, 2, 8, 8},
			r: []int{5, 2, 8},
		},
	}

	for _, td := range data {
		require.Equal(t, td.r, dd.Deduplicate(td.d))
	}
}

// Написать функцию, которая на вход принимает голову
// односвязного списка (можно считать, что в списке хранятся int-ы)
// и удаляет из него все дубликаты, оставляя только первое вхождение.

func TestDeduplicateList(t *testing.T) {
	data := []struct {
		list, result dd.List
	}{
		{
			list:   dd.List{Value: 3, Next: &dd.List{2, &dd.List{2, nil}}},
			result: dd.List{Value: 3, Next: &dd.List{2, nil}},
		},
		{
			list:   dd.List{5, &dd.List{5, &dd.List{2, &dd.List{2, &dd.List{2, nil}}}}},
			result: dd.List{Value: 5, Next: &dd.List{2, nil}},
		},
		{
			list:   dd.List{5, &dd.List{5, &dd.List{2, &dd.List{2, &dd.List{5, nil}}}}},
			result: dd.List{Value: 5, Next: &dd.List{2, nil}},
		},
	}

	for _, td := range data {
		dd.DeduplicateList(&td.list)
		require.Equal(t, td.result, td.list)
	}
}
