package ex_test

import (
	"fmt"

	. "github.com/Nixolay/training/golang/example_test"
)

func ExampleDouble() {
	fmt.Println(Double(1))
	fmt.Println(Double(2))
	// Output:
	// 2
	// 4
}
