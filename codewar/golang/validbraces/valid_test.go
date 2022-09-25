package validbraces_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/validbraces"
	"github.com/stretchr/testify/require"
)

func TestValidBraces(t *testing.T) {
	require.True(t, ValidBraces("(){}[]"))
	require.True(t, ValidBraces("([{}])"))
	require.False(t, ValidBraces("(}"))
	require.False(t, ValidBraces("[(])"))
	require.False(t, ValidBraces("[({)](]"))
}
