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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		}
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}

	if list1 == nil {
		return list2
	}
	return list1
}
