package singlyLinkedList_test

import (
	"testing"

	linkedList "github.com/Nixolay/training/golang/data_structure/singlyLinkedList"
)

func TestNode(t *testing.T) {
	var list linkedList.LinkedList

	const ten, zero = 10, 0

	for i := zero; i < ten; i++ {
		list.Push(i)
	}

	if list.Len() != ten {
		t.Fatal("incorrect list len:", list.Len())
	}
}
