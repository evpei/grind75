package mergetwo

import (
	ds "github.com/evpei/grind-75/src/kit/datastructures"
)

// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.

func Solve(list1 ds.LinkedList[int], list2 ds.LinkedList[int]) ds.LinkedList[int] {
	// iterate as long as either list 1 or two have node
	// appends the smaller value to the result list
	node1 := list1.Head
	node2 := list2.Head

	mergedList := ds.LinkedList[int]{}.New()

	for node1 != nil || node2 != nil {
		if node1 == nil {
			mergedList.Add(node2.Val)
			node2 = node2.Next
		} else if node2 == nil {
			mergedList.Add(node1.Val)
			node1 = node1.Next
		} else if node1.Val == node2.Val {
			mergedList.Add(node1.Val)
			node1 = node1.Next
			mergedList.Add(node2.Val)
			node2 = node2.Next
		} else if node1.Val < node2.Val {
			mergedList.Add(node1.Val)
			node1 = node1.Next
		} else {
			mergedList.Add(node2.Val)
			node2 = node2.Next
		}
	}
	return mergedList
}
