package main

import (
	"time"
	"treealgos/graph/tool"
	"treealgos/internal/shared"
	"treealgos/internal/treetraversal"
	"treealgos/types/trees"
)

func main() {
	// always start with 1 as root node
	start := time.Now()
	tree := &trees.MultiChildTreeNode{
		Val:       1,
		Children:  []*trees.MultiChildTreeNode{},
		IsVisited: false,
	}
	value := 2 // need to refactor to remove dependency on value variable in TreeBuilder
	var counter *int = &value
	treeInput := []int{2, 2, 4} // number of children per node from second level onwards
	treetraversal.TreeBuilder(tree, treeInput, counter, 1)
	shared.Yellow("TreeBuilder BFS took: %v\n", time.Since(start))

	treeGraph := tool.CreateDotFile("testtree", "testtree")

	start = time.Now()
	bfsDepth := treetraversal.BfsMultiChild(tree, treeGraph)
	shared.Yellow("BFS took: %v\n", time.Since(start))

	tool.CloseDotFile(treeGraph)
	tool.CreateGraph(treeGraph, "gif", false)

	start = time.Now()
	dfsDepth := 0
	dfsMaxDepth := 0
	treetraversal.DfsMultiChild(tree, &dfsDepth, &dfsMaxDepth)
	shared.Yellow("DFS took: %v\n", time.Since(start))

	shared.Green("BFS Depth: %v\n", bfsDepth)
	shared.Green("DFS Depth: %v\n", dfsMaxDepth)
}
