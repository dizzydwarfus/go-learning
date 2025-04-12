package main

import (
	"encoding/json"
	"treealgos/internal/shared"
	"treealgos/internal/treetraversal"
)

func main() {
	// always start with 1 as root node
	tree := &treetraversal.MultiChildTreeNode{
		Val:      1,
		Children: []*treetraversal.MultiChildTreeNode{},
	}
	value := 2
	var counter *int = &value
	treeInput := []int{2, 2, 2, 2, 2} // number of children per node from second level onwards
	treetraversal.TreeBuilder(tree, treeInput, counter, 1)

	jsonBytes, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		panic(err)
	}

	shared.Green("Resulting Tree:\n--------------------------\n%v\n", string(jsonBytes))

	// depthSimpleTree := treetraversal.BfsSimple(testdata.SimpleTreeRoot)
	// fmt.Println("Depth for Simple Tree:", depthSimpleTree)

	// depthMultiChildTree := treetraversal.BfsMultiChild(tree)
	// shared.Green("Depth for Multi-Child Tree:", depthMultiChildTree)
}
