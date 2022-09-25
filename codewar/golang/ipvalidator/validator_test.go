package ipvalidator_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/ipvalidator"
	"github.com/stretchr/testify/require"
)

func TestIPValidator(t *testing.T) {
	require.True(t, IsValidIP("12.255.56.1"))
	require.False(t, IsValidIP(""))
	require.False(t, IsValidIP("abc.def.ghi.jkl"))
	require.False(t, IsValidIP("123.456.789.0"))
	require.False(t, IsValidIP("12.34.56"))
	require.False(t, IsValidIP("12.34.56 .1"))
	require.False(t, IsValidIP("12.34.56.-1"))
	require.False(t, IsValidIP("123.045.067.089"))
	require.True(t, IsValidIP("127.1.1.0"))
	require.True(t, IsValidIP("0.0.0.0"))
	require.True(t, IsValidIP("0.34.82.53"))
	require.False(t, IsValidIP("192.168.1.300"))
}
