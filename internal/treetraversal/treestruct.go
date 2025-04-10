package treetraversal

import "fmt"

type BiTreeNode struct {
	Val   int
	Left  *BiTreeNode
	Right *BiTreeNode
}

func (t *BiTreeNode) String() string {
	if t == nil {
		return "nil"
	}
	return fmt.Sprintf("{Val: %d, Left: %v, Right: %v}", t.Val, t.Left, t.Right)
}

type MultiChildTreeNode struct {
	Val      int                   `json:"val"`
	Children []*MultiChildTreeNode `json:"children"`
}

func (t *MultiChildTreeNode) String() string {
	if t == nil {
		return "nil"
	}
	return fmt.Sprintf("{Val: %d, Children: %v}", t.Val, t.Children)
}
