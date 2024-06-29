/**
* Given the root of a binary tree,
* Return the inorder traversal of its nodes' values.
* See: https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 */
package main

// TreeNode represents a node in a binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	helper(root, &res)
	return res
}

// helper function to recursively perform inorder traversal.
func helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	helper(root.Left, res)
	*res = append(*res, root.Val)
	helper(root.Right, res)
}
