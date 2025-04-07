package main

import (
	"fmt"
	"treealgos/internal/treetraversal"
	"treealgos/test/testdata"
)

func main() {
	// pointer basics
	// pointers.PointerBasic()

	depthSimpleTree := treetraversal.BfsSimple(testdata.SimpleTreeRoot)
	fmt.Println("Depth for Simple Tree:", depthSimpleTree)

	// depthMultiChildTree := treetraversal.BfsMultiChild(testdata.MultiChildTreeRoot)
	// fmt.Println("Depth for Multi-Child Tree:", depthMultiChildTree)

	// treeInput := []int{1, 3, 2}
	// treeBuilder(treeInput)
}

// func treeBuilder(nodesPerLevel []int) {

// 	builtTree := &structs.MultiChildTreeNode{}
// 	builtTree.Val = 1

// 	var currentNode structs.MultiChildTreeNode

// 	for i, numberOfNodesPerParent := range nodesPerLevel {
// 		fmt.Printf("Current depth %v\n", i)
// 		fmt.Printf("Nodes in level %v\n\n", numberOfNodesPerParent)

// 	}
// }
