package tool

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"treealgos/internal/shared"
	"treealgos/types/trees"
)

func CreateGraph(f *os.File) {
	_, fileName := filepath.Split(f.Name())
	parsedFileName := strings.Split(fileName, ".")[0]

	shared.Check(os.MkdirAll("graph/svg", os.ModePerm))
	outputFileName := fmt.Sprintf("graph/svg/%v.svg", parsedFileName)

	createDot := exec.Command("dot", "-Tsvg", "-o", outputFileName, f.Name())
	// shared.Faint("%v\n", createDot.String())

	var stderr bytes.Buffer
	createDot.Stderr = &stderr

	if err := createDot.Run(); err != nil {
		shared.Red("dot command failed: %v\nstderr: %s", err, stderr.String())
	}
}

// every line should be either
// 1 "{label}" [color={color}];
// 2 "{parentLabel}" -- "{childLabel}"
// arrow type: --, ->
// arrow styling: [penwidth = 5 fontsize = 28 fontcolor = "black", label = "test"]
func AddNode(f *os.File, node *trees.MultiChildTreeNode) {
	// take value of node as label
	content := fmt.Appendf([]byte{}, "\"%v\";\n", node.Val)

	// for each children add link to current node
	for _, child := range node.Children {
		if len(node.Children) > 0 {
			content = fmt.Appendf([]byte(content), "\"%v\" -- \"%v\"\n", node.Val, child.Val)
		}
	}
	n, err := f.Write(content)
	shared.Check(err)
	shared.Faint("Written %v bytes for node.Val=%v with %v children\n", n, node.Val, len(node.Children))
}

func CreateDotFile(f, g string) *os.File {
	f = filepath.Join("graph", "dot", f)
	fileName := fmt.Sprintf("%v.dot", f)

	startGraph := fmt.Appendf([]byte{}, "graph %v {\n", g)

	if file, err := os.Create(fileName); err == nil {
		shared.Green("Successfully created %v\n", fileName)

		n, err := file.Write(startGraph)
		shared.Check(err)
		shared.Faint("Written %v bytes\n", n)
		return file
	} else {
		log.Fatal(shared.Sred("Failed to create %v: %v", fileName, err))
		return nil
	}
}

func CloseDotFile(f *os.File) {
	f.WriteString("}\n")
	shared.Yellow("Finishing .dot file creation, closing %v\n", f.Name())
	err := f.Close()
	if err != nil {
		log.Fatal(shared.Sred("Failed to close: %v\n", err))
	}
}
