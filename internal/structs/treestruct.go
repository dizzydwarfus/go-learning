package structs

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
	Val      int
	Children []*MultiChildTreeNode
}

func (t *MultiChildTreeNode) String() string {
	if t == nil {
		return "nil"
	}
	return fmt.Sprintf("{Val: %d, Children: %v}", t.Val, t.Children)
}
