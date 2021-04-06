// Package singlylinkedlist пример односвязного списка.
package singlylinkedlist

// Node represents a node of linked list.
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list.
type LinkedList struct {
	head *Node
	len  int
}

// GetSlice will return slice of values.
func (linkedList LinkedList) GetSlice() []int {
	s := make([]int, 0, linkedList.len)

	node := linkedList.head
	for node != nil {
		s = append(s, node.value)
		node = node.next
	}

	return s
}

// Len получить длину связного списка.
func (linkedList LinkedList) Len() int {
	return linkedList.len
}

// Push добавить элемент из связного списка.
//nolint:wsl
func (linkedList *LinkedList) Push(v int) {
	if linkedList.head == nil {
		linkedList.len++
		linkedList.head = &Node{value: v}

		return
	}

	linkedList.len = 1

	node := linkedList.head
	for node.next != nil {
		linkedList.len++
		node = node.next
	}

	linkedList.len++
	node.next = &Node{value: v}
}

// Pop вытащить элемент из связного списка.
//nolint:wsl
func (linkedList *LinkedList) Pop() int {
	if linkedList.head == nil {
		return 0
	}

	defer func() { linkedList.len-- }()

	if linkedList.head.next == nil {
		v := linkedList.head.value
		linkedList.head = nil

		return v
	}

	tail := linkedList.head
	var prevTail *Node

	for tail.next != nil {
		prevTail = tail
		tail = tail.next
	}

	prevTail.next = nil

	return tail.value
}

// Get получение элемента из связного списка.
func (linkedList *LinkedList) Get(iter int) int {
	node := linkedList.head
	for i := 0; i < iter; i++ {
		node = node.next
	}

	return node.value
}

// Remove удаление элемента из связного списка.
func (linkedList *LinkedList) Remove(iter int) {
	node := linkedList.head
	for i := 0; i < iter || node.next != nil; i++ {
		node = node.next
	}

	if node.next.next == nil {
		return
	}

	node.next = node.next.next
}
