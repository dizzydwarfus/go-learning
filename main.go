package main

import (
	"fmt"
	"net/http"
	"treealgos/internal/stream"
	"treealgos/internal/structs"
	"treealgos/internal/treetraversal"
)

func main() {
	addr := new(string)
	*addr = "localhost:8000"
	http.HandleFunc("/", stream.Home)
	http.HandleFunc("/events", stream.SendEvents)
	http.ListenAndServe(*addr, nil)

	// pointers.PointerBasic()

	// depthSimpleTree := treetraversal.BfsSimple(testdata.SimpleTreeRoot)
	// fmt.Println("Depth for Simple Tree:", depthSimpleTree)

	// always start with 1 as root node
	root := &structs.MultiChildTreeNode{
		Val:      1,
		Children: []*structs.MultiChildTreeNode{},
	}
	currentCounter := root.Val

	treeInput := []int{1, 3, 2} // number of children per node from second level onwards
	builtTree := treetraversal.TreeBuilder(root, treeInput, currentCounter, 1)
	fmt.Printf("Resulting Tree: %v\n", builtTree)

	depthMultiChildTree := treetraversal.BfsMultiChild(builtTree)
	fmt.Println("Depth for Multi-Child Tree:", depthMultiChildTree)
}
