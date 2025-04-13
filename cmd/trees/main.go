package main

import (
	"treealgos/graph/tool"
	"treealgos/internal/shared"
	"treealgos/internal/treetraversal"
	"treealgos/types/trees"
)

func main() {
	// always start with 1 as root node
	tree := &trees.MultiChildTreeNode{
		Val:      1,
		Children: []*trees.MultiChildTreeNode{},
	}
	value := 2 // need to refactor to remove dependency on value variable in TreeBuilder
	var counter *int = &value
	treeInput := []int{3, 2, 2} // number of children per node from second level onwards
	treetraversal.TreeBuilder(tree, treeInput, counter, 1)

	f := tool.CreateDotFile("testtree", "testtree")
	depth := treetraversal.BfsMultiChild(tree, f)
	shared.Green("Depth: %v\n", depth)
	tool.CloseDotFile(f)
	tool.CreateGraph(f, "gif", false)

}
