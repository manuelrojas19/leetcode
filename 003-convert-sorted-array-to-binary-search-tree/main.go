/**
* Given an integer array nums where the elements are sorted in ascending order,
* Convert it to a height-balanced binary search tree.
* See: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 */
package main

// TreeNode represents a node in a binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

// sortedArrayToBSTHelper is a recursive helper function to convert a range of the array to a BST.
func sortedArrayToBSTHelper(nums []int, beg, end int) *TreeNode {
	if beg > end {
		return nil
	}
	mid := (beg + end) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBSTHelper(nums, beg, mid-1)
	root.Right = sortedArrayToBSTHelper(nums, mid+1, end)
	return root
}
