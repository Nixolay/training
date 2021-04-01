package stack

type SimpleStack []int

func (s SimpleStack) Empty() bool {
	return len(s) == zero
}

func (s *SimpleStack) Push(v int) {
	(*s) = append((*s), v)
}

func (s *SimpleStack) Pop() int {
	v := (*s)[len(*s)-one]
	(*s) = (*s)[:len(*s)-one]
	return v
}
