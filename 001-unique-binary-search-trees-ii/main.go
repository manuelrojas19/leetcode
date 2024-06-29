/**
 * Given an integer n,
 * return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n.
 * Return the answer in any order.
 * See: https://leetcode.com/problems/unique-binary-search-trees-ii/
 */
package main

import (
	"fmt"
)

// TreeNode represents a node in a binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// GenerateTrees generates all structurally unique BSTs for values from start to end.
func GenerateTrees(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{}
	}

	memory := make(map[string][]*TreeNode)
	return helper(start, end, memory)
}

func helper(start, end int, memory map[string][]*TreeNode) []*TreeNode {
	key := fmt.Sprintf("%d-%d", start, end)

	if val, exists := memory[key]; exists {
		return val
	}

	var nodes []*TreeNode

	if start > end {
		nodes = append(nodes, nil)
	} else {
		for root := start; root <= end; root++ {
			leftNodes := helper(start, root-1, memory)
			rightNodes := helper(root+1, end, memory)

			for _, leftNode := range leftNodes {
				for _, rightNode := range rightNodes {
					node := &TreeNode{Val: root}
					node.Left = leftNode
					node.Right = rightNode
					nodes = append(nodes, node)
				}
			}
		}
	}

	memory[key] = nodes
	return nodes
}

func main() {
	n := 3
	bsts := GenerateTrees(1, n)
	for _, bst := range bsts {
		printTree(bst)
		fmt.Println()
	}
}

// Helper function to print the tree
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
