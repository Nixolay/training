package numberofunits_test

import (
	"testing"

	"github.com/Nixolay/training/codewar/golang/numberofunits"
	"github.com/stretchr/testify/require"
)

func TestLongestSequenceOfUnits(t *testing.T) {
	data := []struct {
		d []int64
		r int
	}{
		{
			d: []int64{5, 1, 0, 1, 0, 1},
			r: 0,
		},
		{
			d: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			r: 6,
		},
	}

	for _, td := range data {
		require.Equal(t, numberofunits.LongestSequenceOfUnits(td.d), td.r)
	}
}
