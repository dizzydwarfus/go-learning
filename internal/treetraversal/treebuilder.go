package treetraversal

import (
	"encoding/json"
	"treealgos/internal/shared"
	"treealgos/types/trees"
)

func TreeBuilder(node *trees.MultiChildTreeNode, nodesPerLevel []int, currentCounter *int, depth int) {
	if depth == len(nodesPerLevel)+1 {
		return
	}

	if *currentCounter == node.Val+1 {
		shared.Magenta("Starting TreeBuilder: %v\n", node)
	} else {
		shared.Cyan("Recursively building tree: %v\n", node)
	}

	shared.Yellow("Appending Children\n")
	shared.Yellow("--------------------------\n")
	for range nodesPerLevel[depth-1] { // append nodesPerLevel as number of children
		shared.Faint("Current Counter %v\n", *currentCounter)
		shared.Faint("Current Level %v\n", depth)
		shared.Faint("# of Nodes in this level: %v\n\n", nodesPerLevel[depth-1])

		node.Children = append(node.Children, &trees.MultiChildTreeNode{Val: *currentCounter, Children: []*trees.MultiChildTreeNode{}})
		*currentCounter++
	}

	for _, child := range node.Children { // for each children call TreeBuilder again
		TreeBuilder(child, nodesPerLevel, currentCounter, depth+1)
		// shared.Yellow("Resulting Tree: %v\n\n", node)
	}
}

func ShowJSONTree(tree *trees.MultiChildTreeNode) []byte {
	jsonBytes, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		panic(err)
	}
	shared.Green("Resulting Tree:\n--------------------------\n%v\n", string(jsonBytes))
	return jsonBytes
}
