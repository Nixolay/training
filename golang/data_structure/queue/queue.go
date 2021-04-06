// Package queue ...
package queue

const zero = 0

// Queue сама очередь.
type Queue []int

// Empty проверка на пустоту.
func (q Queue) Empty() bool {
	return len(q) == zero
}

// Push положить элемент в очередь.
func (q *Queue) Push(v int) {
	(*q) = append((*q), v)
}

// Pop вытаскивает элемент из очереди.
func (q *Queue) Pop() int {
	v := (*q)[zero]
	(*q) = (*q)[1:len(*q)]

	return v
}
