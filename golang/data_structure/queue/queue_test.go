package queue_test

import (
	"testing"

	"github.com/Nixolay/training/golang/data_structure/queue"
)

//nolint:paralleltest
func TestQueue(t *testing.T) {
	var q queue.Queue

	for i := range [10]int{} {
		q.Push(i)
	}

	for i := range q {
		if v := q.Pop(); i != v {
			t.Fatalf("i: %d is not equal v: %d", i, v)
		}
	}
}
