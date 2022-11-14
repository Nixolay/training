package deduplicate_test

import (
	"testing"

	"github.com/Nixolay/training/codewar/golang/deduplicate"
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
		require.Equal(t, td.r, deduplicate.Deduplicate(td.d))
	}
}
