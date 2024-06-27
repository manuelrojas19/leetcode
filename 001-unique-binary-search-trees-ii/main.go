package main

import (
	"fmt"
)

// Node represents a node in a binary search tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// GenerateTrees generates all structurally unique BSTs for values from start to end.
func GenerateTrees(start, end int) []*Node {
	if start > end {
		return []*Node{}
	}

	memory := make(map[string][]*Node)
	return generateTrees(start, end, memory)
}

func generateTrees(start, end int, memory map[string][]*Node) []*Node {
	key := fmt.Sprintf("%d-%d", start, end)

	if val, exists := memory[key]; exists {
		return val
	}

	var nodes []*Node

	if start > end {
		nodes = append(nodes, nil)
	} else {
		for root := start; root <= end; root++ {
			leftNodes := generateTrees(start, root-1, memory)
			rightNodes := generateTrees(root+1, end, memory)

			for _, leftNode := range leftNodes {
				for _, rightNode := range rightNodes {
					node := &Node{Value: root}
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
func printTree(root *Node) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Value)
	printTree(root.Left)
	printTree(root.Right)
}
