package treetraversal

import (
	"fmt"
	"treealgos/internal/structs"
)

func TreeBuilder(node *structs.MultiChildTreeNode, nodesPerLevel []int, currentCounter int, depth int) *structs.MultiChildTreeNode {
	if len(nodesPerLevel) == 0 {
		return node
	}
	fmt.Printf("Starting Tree: %v\n\n", node)

	for range nodesPerLevel[0] {
		fmt.Printf("Current Counter %v\n", currentCounter)
		fmt.Printf("Current Level %v\n", depth)
		fmt.Printf("# of Nodes in this level: %v\n", nodesPerLevel[0])

		node.Children = append(node.Children, &structs.MultiChildTreeNode{Val: currentCounter, Children: []*structs.MultiChildTreeNode{}})
		currentCounter++
		node = TreeBuilder(node, nodesPerLevel[1:], currentCounter, depth+1)
	}
	fmt.Printf("nodesPerlevel sliced: %v\n", nodesPerLevel[1:])
	fmt.Printf("Resulting Tree: %v\n\n", node)
	return node
	// newNode := &structs.MultiChildTreeNode{Val: currentCounter, Children: []*structs.MultiChildTreeNode{}}
}
