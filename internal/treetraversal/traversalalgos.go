package treetraversal

import (
	"fmt"
)

func BfsSimple(root *BiTreeNode) int {
	if root == nil {
		return 0
	}

	nodeList := []*BiTreeNode{root}
	depth := 0

	for len(nodeList) > 0 {

		for _, node := range nodeList {
			// fmt.Println("Before Slicing", nodeList)
			nodeList = nodeList[1:]
			// fmt.Println("After Slicing", nodeList)

			// fmt.Println("Current node value: ", node.Val)

			if node.Left != nil {
				// fmt.Println("Appending left node: ", node.Left)
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				// fmt.Println("Appending right node: ", node.Right)
				nodeList = append(nodeList, node.Right)
			}
			// fmt.Println()
		}

		depth++
	}

	return depth
}

func BfsMultiChild(root *MultiChildTreeNode) int {
	if root == nil {
		return 0
	}

	nodeList := []*MultiChildTreeNode{root}
	depth := 0

	for len(nodeList) > 0 {

		for _, node := range nodeList {
			fmt.Println("Before Slicing", nodeList)
			nodeList = nodeList[1:]
			fmt.Println("After Slicing", nodeList)

			fmt.Println("Current node value: ", node.Val)

			if node.Children != nil {
				fmt.Println("Appending children nodes: ", node.Children)
				nodeList = append(nodeList, node.Children...)
			}
			fmt.Println()
		}

		depth++
	}

	return depth
}
