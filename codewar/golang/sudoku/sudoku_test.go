package sudoku_test

import (
	"testing"

	. "github.com/Nixolay/training/codewar/golang/sudoku"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSudoku(t *testing.T) {
	testTrue := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	testFalse := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 0, 3, 4, 8},
		{1, 0, 0, 3, 4, 2, 5, 6, 0},
		{8, 5, 9, 7, 6, 1, 0, 2, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 0, 1, 5, 3, 7, 2, 1, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 0, 0, 4, 8, 1, 1, 7, 9},
	}

	testFalse2 := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{2, 3, 1, 5, 6, 4, 8, 9, 7},
		{3, 1, 2, 6, 4, 5, 9, 7, 8},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{5, 6, 4, 8, 9, 7, 2, 3, 1},
		{6, 4, 5, 9, 7, 8, 3, 1, 2},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{8, 9, 7, 2, 3, 1, 5, 6, 4},
		{9, 7, 8, 3, 1, 2, 6, 4, 5},
	}

	Convey("should test that the solution returns the correct value", t, func() {
		So(ValidateSolution(testTrue), ShouldBeTrue)
		So(ValidateSolution(testFalse), ShouldBeFalse)
		So(ValidateSolution(testFalse2), ShouldBeFalse)
	})
}
