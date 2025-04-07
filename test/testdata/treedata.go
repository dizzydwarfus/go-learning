package testdata

import "treealgos/internal/structs"

var (
	SimpleTreeRoot = &structs.BiTreeNode{
		Val: 1,
		Left: &structs.BiTreeNode{
			Val: 2,
			Left: &structs.BiTreeNode{
				Val: 4,
			},
			Right: &structs.BiTreeNode{
				Val: 5,
			},
		},
		Right: &structs.BiTreeNode{
			Val: 3,
		},
	}
	MultiChildTreeRoot = &structs.MultiChildTreeNode{
		Val: 1,
		Children: []*structs.MultiChildTreeNode{
			{
				Val: 2,
				Children: []*structs.MultiChildTreeNode{
					&structs.MultiChildTreeNode{Val: 6, Children: []*structs.MultiChildTreeNode{
						&structs.MultiChildTreeNode{Val: 10, Children: []*structs.MultiChildTreeNode{}},
					}},
				},
			},
			{
				Val: 3,
				Children: []*structs.MultiChildTreeNode{
					&structs.MultiChildTreeNode{Val: 7, Children: []*structs.MultiChildTreeNode{
						&structs.MultiChildTreeNode{Val: 11, Children: []*structs.MultiChildTreeNode{}},
					}},
				},
			},
			{
				Val: 4,
				Children: []*structs.MultiChildTreeNode{
					&structs.MultiChildTreeNode{Val: 8, Children: []*structs.MultiChildTreeNode{
						&structs.MultiChildTreeNode{Val: 12, Children: []*structs.MultiChildTreeNode{}},
					}},
				},
			},
			{
				Val: 5,
				Children: []*structs.MultiChildTreeNode{
					&structs.MultiChildTreeNode{Val: 9, Children: []*structs.MultiChildTreeNode{
						&structs.MultiChildTreeNode{Val: 13, Children: []*structs.MultiChildTreeNode{}},
					}},
				},
			},
		},
	}
)
