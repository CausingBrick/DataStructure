package tree

import (
	"fmt"
)

//BinaryNode 二叉树节点结构
type BinaryNode struct {
	Parent int
	Left   int
	Right  int
}

// DepthRec 返回树的深度, 递归
func (r *BinaryNode) DepthRec(tree []BinaryNode) int {
	if r.Parent == -1 {
		return 0
	}
	return tree[r.Parent].DepthRec(tree) + 1
}

// Depth 求节点在树的深度
func (r *BinaryNode) Depth(tree []BinaryNode) int {
	depth := 0
	node := *r
	for node.Parent != -1 {
		depth++
		node = tree[node.Parent]
	}
	return depth
}

// Height return the heigth of the node in the tree
func (r *BinaryNode) Height(tree []BinaryNode) int {
	//leaf node
	if r.Left == -1 && r.Right == -1 {
		return 0
	}
	// right single subtree
	if r.Left == -1 {
		return tree[r.Right].Height(tree) + 1
	}
	// left single subtree
	if r.Right == -1 {
		return tree[r.Left].Height(tree) + 1
	}

	return max(tree[r.Right].Height(tree)+1, tree[r.Left].Height(tree)+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Value return the node value in the tree.
func (r *BinaryNode) Value(tree []BinaryNode) int {
	if r.Parent != -1 {
		return tree[r.Parent].Left
	}
	if r.Left != -1 {
		return tree[r.Left].Parent
	}
	if r.Right != -1 {
		return tree[r.Right].Parent
	}
	return 0
}

// Degrees return the degree of the node.
func (r *BinaryNode) Degrees(tree []BinaryNode) int {
	if r.Left == -1 && r.Right == -1 {
		return 0
	}
	if r.Left != -1 && r.Right != -1 {
		return 2
	}
	return 1
}

// Type return the kind of node in the tree.
func (r *BinaryNode) Type() string {
	if r.Parent == -1 {
		return "root"
	}
	if r.Left == -1 && r.Right == -1 {
		return "leaf"
	}
	return "internal node"
}

// Brother return the index of brother node.
func (r *BinaryNode) Brother(tree []BinaryNode) int {
	if r.Parent == -1 {
		return -1
	}
	if tree[r.Parent].Left == r.Value(tree) {
		return tree[r.Parent].Right
	}
	return tree[r.Parent].Left
}

// Print 输出树的信息
func (r *BinaryNode) Print(tree []BinaryNode) {
	node := *r
	if node.Left != -1 {
		fmt.Printf("node %d: parent = %d, brother = %d, degree = %d, depth = %d, height = %d, type:%s\n",
			node.Value(tree),
			node.Parent,
			node.Brother(tree),
			node.Degrees(tree),
			node.DepthRec(tree),
			node.Height(tree),
			node.Type())
	}
}
