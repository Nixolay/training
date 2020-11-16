package stack_test

import (
	"testing"

	. "github.com/Nixolay/training/golang/data_structure/stack"
	. "github.com/smartystreets/goconvey/convey"
)

type Data struct {
	brackets string
	ok       bool
	position int
}

func TestBracketsIscorrectly(t *testing.T) {
	rangeData := []Data{
		{brackets: "([](){([])})", ok: true, position: 0},
		{brackets: "()[]}", ok: false, position: 5},
		{brackets: "{{[()]]", ok: false, position: 7},
		{brackets: "{{{[][][]", ok: false, position: 3},
		{brackets: "{*{{}", ok: false, position: 3},
		{brackets: "[[*", ok: false, position: 2},
		{brackets: "{*}", ok: true, position: 0},
		{brackets: "{{", ok: false, position: 2},
		{brackets: "{}", ok: true, position: 0},
		{brackets: "", ok: true, position: 0},
		{brackets: "}", ok: false, position: 1},
		{brackets: "*{}", ok: true, position: 0},
		{brackets: "{{{**[][][]", ok: false, position: 3},
	}

	//nolint:scopelint
	for _, data := range rangeData {
		Convey(data.brackets, t, func() {
			position, ok := BracketsIscorrectly(data.brackets)
			So(position, ShouldEqual, data.position)
			So(ok, ShouldEqual, data.ok)
		})
	}
}
