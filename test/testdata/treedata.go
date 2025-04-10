package testdata

import "treealgos/internal/treetraversal"

var (
	SimpleTreeRoot = &treetraversal.BiTreeNode{
		Val: 1,
		Left: &treetraversal.BiTreeNode{
			Val: 2,
			Left: &treetraversal.BiTreeNode{
				Val: 4,
			},
			Right: &treetraversal.BiTreeNode{
				Val: 5,
			},
		},
		Right: &treetraversal.BiTreeNode{
			Val: 3,
		},
	}
	MultiChildTreeRoot = &treetraversal.MultiChildTreeNode{
		Val: 1,
		Children: []*treetraversal.MultiChildTreeNode{
			{
				Val: 2,
				Children: []*treetraversal.MultiChildTreeNode{
					&treetraversal.MultiChildTreeNode{Val: 6, Children: []*treetraversal.MultiChildTreeNode{
						&treetraversal.MultiChildTreeNode{Val: 10, Children: []*treetraversal.MultiChildTreeNode{}},
					}},
				},
			},
			{
				Val: 3,
				Children: []*treetraversal.MultiChildTreeNode{
					&treetraversal.MultiChildTreeNode{Val: 7, Children: []*treetraversal.MultiChildTreeNode{
						&treetraversal.MultiChildTreeNode{Val: 11, Children: []*treetraversal.MultiChildTreeNode{}},
					}},
				},
			},
			{
				Val: 4,
				Children: []*treetraversal.MultiChildTreeNode{
					&treetraversal.MultiChildTreeNode{Val: 8, Children: []*treetraversal.MultiChildTreeNode{
						&treetraversal.MultiChildTreeNode{Val: 12, Children: []*treetraversal.MultiChildTreeNode{}},
					}},
				},
			},
			{
				Val: 5,
				Children: []*treetraversal.MultiChildTreeNode{
					&treetraversal.MultiChildTreeNode{Val: 9, Children: []*treetraversal.MultiChildTreeNode{
						&treetraversal.MultiChildTreeNode{Val: 13, Children: []*treetraversal.MultiChildTreeNode{}},
					}},
				},
			},
		},
	}
)
