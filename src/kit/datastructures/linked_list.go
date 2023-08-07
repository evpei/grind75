package datastructure

type LinkedListNode[T any] struct {
	Val  T
	Next *LinkedListNode[T]
}

type LinkedList[T any] struct {
	Head *LinkedListNode[T]
}

func (list LinkedList[T]) New(vals ...T) LinkedList[T] {
	for _, val := range vals {
		list.Add(val)
	}
	return list
}

// Adds a value to the end of the linked list
func (list *LinkedList[T]) Add(val T) {
	if list.Head == nil {
		list.Head = &LinkedListNode[T]{Val: val}
		return
	}
	node := list.Head

	for node.Next != nil {
		node = node.Next
	}
	next := LinkedListNode[T]{Val: val}
	node.Next = &next
}
