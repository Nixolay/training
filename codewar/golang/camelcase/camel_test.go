package camel_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/camelcase"
	"github.com/stretchr/testify/require"
)

func TestToCamelCase(t *testing.T) {
	t.Run("should handle basic cases", func(t *testing.T) {
		require.Equal(t, ToCamelCase(""), "")
		require.Equal(t, ToCamelCase("The_Stealth_Warrior"), "TheStealthWarrior")
		require.Equal(t, ToCamelCase("The_Stealth_Warrior"), "TheStealthWarrior")
		require.Equal(t, ToCamelCase("the-stealth-Warrior"), "theStealthWarrior")
		require.Equal(t, ToCamelCase("the stealth warrior"), "theStealthWarrior")
	})
}
