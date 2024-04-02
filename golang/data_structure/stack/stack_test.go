package stack_test

import (
	"testing"

	"github.com/Nixolay/training/golang/data_structure/stack"
	"github.com/stretchr/testify/require"
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

	for _, data := range rangeData {
		t.Run(data.brackets, func(t *testing.T) {
			position, ok := stack.BracketsIscorrectly(data.brackets)
			require.Equal(t, position, data.position)
			require.Equal(t, ok, data.ok)
		})
	}
}
