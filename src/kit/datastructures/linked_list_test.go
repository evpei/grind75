package datastructure

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := LinkedList[int]{}.New(1, 2)
	node := list.Head
	if node.Val != 1 {
		t.Errorf("Expected 1, but got %d", node.Val)
	}
	node = node.Next
	if node.Val != 2 {
		t.Errorf("Expected 2, but got %d", node.Val)
	}
	if node.Next != nil {
		t.Errorf("Expected 2 nodes, but got more")
	}
}

func TestAppendToLinkedList(t *testing.T) {
	list := LinkedList[int]{}.New(1, 2, 3)
	list.Add(4)
	node := list.Head

	for i := 1; i <= 4; i++ {
		if node.Val != i {
			t.Errorf("Expected %d, but got %d", i, node.Val)
		}
		node = node.Next
	}
}

func TestNewEmpty(t *testing.T) {
	list := LinkedList[int]{}
	if list.Head != nil {
		t.Errorf("Head is not nil")
	}
}
