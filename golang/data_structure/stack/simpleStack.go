// Package stack пример стека.
package stack

// SimpleStack ...
type SimpleStack []int

// Empty проверка на пустоту.
func (s SimpleStack) Empty() bool {
	return len(s) == zero
}

// Push положить элемент.
func (s *SimpleStack) Push(v int) {
	(*s) = append((*s), v)
}

// Pop вытаскивает элемент.
func (s *SimpleStack) Pop() int {
	v := (*s)[len(*s)-one]
	(*s) = (*s)[:len(*s)-one]

	return v
}
