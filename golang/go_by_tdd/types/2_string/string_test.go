package string_test

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestArray(t *testing.T) {
	t.Parallel()

	arr := [8]int64{1, 2, 3}
	fmt.Printf("%#v\n", arr)
}

func TestSlice(t *testing.T) {
	var slice0 []int                     // len=0, cap=0
	slice1 := []int{6, 1, 2}             // len=3,cap=3
	slice2 := make([]int, 0, 3)          // len=0,cap=3
	slice3 := append([]int{}, slice2...) // len=0, cap=0
	slice4 := append(slice1, 3)          // len=4, cap=4

	println("clice0", len(slice0), cap(slice0))
	println("clice1", len(slice1), cap(slice1))
	println("clice2", len(slice2), cap(slice2))
	println("clice3", len(slice3), cap(slice3))
	println("clice4", len(slice4), cap(slice4))
}

func TestString(t *testing.T) {
	sample := "Привет мир! 你好世界!"
	println(utf8.RuneCountInString(sample))
	println(len([]rune(sample)))
	println(len(sample))

	for _, data := range sample {
		PrintType(data)
	}
}

func PrintType(data interface{}) {
	switch v := data.(type) {
	case rune:
		println("rune:", v)
	default:
		println("not rune:", v)
	}
}
