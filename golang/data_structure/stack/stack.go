// Package stack practice stack
package stack

import (
	"strings"
)

const (
	zero = iota
	one
)

// BracketsIscorrectly checks for correct brackets.
func BracketsIscorrectly(brackets string) (int, bool) {
	var position, stack Stack

	bracketsMap := map[rune]rune{'[': ']', '(': ')', '{': '}'}

	for i, r := range brackets {
		if _, ok := bracketsMap[r]; ok {
			position.push(rune(i + one))
			stack.push(r)

			continue
		}

		//nolint:gocritic
		if !strings.Contains("]})", string(r)) {
			continue
		}

		if !stack.empty() {
			top, ok := stack.pop()

			if end := bracketsMap[top]; ok && end == r {
				position.pop()

				continue
			}
		}

		position.push(rune(i + one))
		stack.push(r)

		break
	}

	p, _ := position.pop()

	return int(p), false || stack.empty()
}

// Stack implements the underlying data structure.
type Stack struct {
	buf []rune
}

func (stack *Stack) push(item rune) {
	stack.buf = append(stack.buf, item)
}

func (stack *Stack) pop() (rune, bool) {
	if len(stack.buf) > zero {
		out := stack.buf[len(stack.buf)-one]
		stack.buf = stack.buf[:len(stack.buf)-one]

		return out, true
	}

	return zero, false
}

func (stack *Stack) empty() bool {
	return stack.len() == zero
}

func (stack *Stack) len() int {
	return len(stack.buf)
}
