/**
* You are given the heads of two sorted linked lists list1 and list2.
* Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
* Return the head of the merged linked list.
* See: https://leetcode.com/problems/merge-two-sorted-lists/description/
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			list1.Next = mergeTwoListsRecursive(list1.Next, list2)
			return list1
		}
		list2.Next = mergeTwoListsRecursive(list2.Next, list1)
		return list2
	}

	if list1 == nil {
		return list2
	}
	return list1
}

func mergeTwoListsIterative(list1 *ListNode, list2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	currentNode := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}
		currentNode = currentNode.Next
	}

	// Append the remaining nodes of list1 or list2 (if any)
	if list1 != nil {
		currentNode.Next = list1
	} else {
		currentNode.Next = list2
	}

	// Return the head of the merged list (skip the dummy node)
	return dummyHead.Next
}
