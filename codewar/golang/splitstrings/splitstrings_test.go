package splitstrings_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/splitstrings"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, Solution("abc"), []string{"ab", "c_"})
	require.Equal(t, Solution("abcdef"), []string{"ab", "cd", "ef"})
}
