// Package queue ...
package queue

const zero = 0

type Queue []int

func (q Queue) Empty() bool {
	return len(q) == zero
}

func (q *Queue) Push(v int) {
	(*q) = append((*q), v)
}

func (q *Queue) Pop() int {
	v := (*q)[zero]
	(*q) = (*q)[1:len(*q)]

	return v
}
