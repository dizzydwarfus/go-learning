package tool

import (
	"fmt"
	"log"
	"os"
	"treealgos/internal/shared"
	"treealgos/internal/treetraversal"
)

// every line should be either
// 1 "{label}" [color={color}];
// 2 "{parentLabel}" -- "{childLabel}"
// arrow type: --, ->
// arrow styling: [penwidth = 5 fontsize = 28 fontcolor = "black", label = "test"]
func AddNode(f *os.File, node *treetraversal.MultiChildTreeNode) {

}

func CreateGraph(f, g string) *os.File {
	fileName := fmt.Sprintf("%v.dot", f)

	startGraph := fmt.Appendf([]byte{}, "diagraph %v {\n}", g)

	if file, err := os.Create(fileName); err == nil {
		shared.Green("Successfully created %v\n", fileName)
		defer file.Close()

		n, err := file.Write(startGraph)
		shared.Check(err)
		shared.Faint("Written %v bytes\n", n)
		return file
	} else {
		log.Fatal(shared.Sred("Failed to create %v: %v", fileName, err))
		return nil
	}
}
