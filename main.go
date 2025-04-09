package main

import (
	"treealgos/internal/structs"
	"treealgos/internal/treetraversal"
)

func main() {
	// addr := new(string)
	// *addr = "localhost:8000"
	// http.HandleFunc("/", stream.Home)
	// http.HandleFunc("/events", stream.SendEvents)
	// http.ListenAndServe(*addr, nil)

	// pointers.PointerBasic()

	// depthSimpleTree := treetraversal.BfsSimple(testdata.SimpleTreeRoot)
	// fmt.Println("Depth for Simple Tree:", depthSimpleTree)

	// always start with 1 as root node
	tree := &structs.MultiChildTreeNode{
		Val:      1,
		Children: []*structs.MultiChildTreeNode{},
	}
	value := 2
	var counter *int = &value
	treeInput := []int{2, 2} // number of children per node from second level onwards
	treetraversal.TreeBuilder(tree, treeInput, counter, 1)

	// depthMultiChildTree := treetraversal.BfsMultiChild(tree)
	// shared.Green("Depth for Multi-Child Tree:", depthMultiChildTree)
}
