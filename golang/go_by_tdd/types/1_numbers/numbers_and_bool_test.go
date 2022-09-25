package basetypes_test

import (
	"testing"
)

func TestInt(t *testing.T) {
	t.Parallel()

	data := []any{} // заполните целочисленными переменными.
	for _, v := range data {
		if !isInt(v) {
			t.Fatal("item is not int:", v)
		}
	}
}

func TestUint(t *testing.T) {
	t.Parallel()

	data := []any{} // заполните беззнаковыми целыми.
	for _, v := range data {
		if !isUint(v) {
			t.Fatal("item is not uint:", v)
		}
	}
}

func TestFloat(t *testing.T) {
	t.Parallel()

	data := []any{} // заполните числами с плавающей запятой.
	for _, v := range data {
		if !isFloat(v) {
			t.Fatal("item is not float:", v)
		}
	}
}

func isInt(i interface{}) bool {
	switch i.(type) {
	case int, int8, int16, int32, int64:
		return true
	}

	return false
}

func isUint(i interface{}) bool {
	switch i.(type) {
	case uint, uint8, uint16, uint32, uint64:
		return true
	}

	return false
}

func isFloat(i interface{}) bool {
	switch i.(type) {
	case float32, float64:
		return true
	}

	return false
}
