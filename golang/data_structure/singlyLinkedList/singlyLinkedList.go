package singlyLinkedList

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

// GetSlice will return slice of values
func (list LinkedList) GetSlice() []int {
	s := make([]int, 0, list.len)

	node := list.head
	for node != nil {
		s = append(s, node.value)
		node = node.next
	}

	return s
}

func (linkedList LinkedList) Len() int {
	return linkedList.len
}

func (list *LinkedList) Push(v int) {
	if list.head == nil {
		list.len++
		list.head = &Node{value: v}
		return
	}

	list.len = 1

	node := list.head
	for node.next != nil {
		list.len++
		node = node.next
	}

	list.len++
	node.next = &Node{value: v}
}

func (list *LinkedList) Pop() int {
	if list.head == nil {
		return 0
	}

	defer func() { list.len-- }()

	if list.head.next == nil {
		v := list.head.value
		list.head = nil
		return v
	}

	tail := list.head
	var prevTail *Node

	for tail.next != nil {
		prevTail = tail
		tail = tail.next
	}

	prevTail.next = nil

	return tail.value
}

func (list *LinkedList) Get(iter int) int {
	node := list.head
	for i := 0; i < iter; i++ {
		node = node.next
	}

	return node.value
}

func (list *LinkedList) Remove(iter int) {
	node := list.head
	for i := 0; i < iter || node.next != nil; i++ {
		node = node.next
	}

	if node.next.next == nil {
		return
	}

	node.next = node.next.next
}
