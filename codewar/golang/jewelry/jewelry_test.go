package jewelry_test

import (
	"testing"

	"github.com/Nixolay/training/codewar/golang/jewelry"
	"github.com/stretchr/testify/require"
)

/*
Даны две строки строчных латинских символов: строка J и строка S. Символы, входящие в строку J, — «драгоценности», входящие в строку S — «камни». Нужно определить, какое количество символов из S одновременно являются «драгоценностями». Проще говоря, нужно проверить, какое количество символов из S входит в J.
*/

func TestJewelry(t *testing.T) {
	data := []struct {
		j, s string
		r    int
	}{
		{
			j: "ab",
			s: "aabbccd",
			r: 4,
		},
		{
			j: "cd",
			s: "aabbccd",
			r: 3,
		},
	}

	for _, td := range data {
		require.Equal(t, jewelry.Jewelry(td.j, td.s), td.r)
	}
}
