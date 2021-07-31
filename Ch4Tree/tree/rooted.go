package tree

import (
	"fmt"
)

//RootedNode 左子右兄弟表示法结构
type RootedNode struct {
	Parent int
	Left   int
	Right  int
}

// DepthRec 返回树的深度, 递归
func (r *RootedNode) DepthRec(tree []RootedNode) int {
	if r.Parent == -1 {
		return 0
	}
	return tree[r.Parent].DepthRec(tree) + 1
}

// Depth 求节点在树的深度
func (r *RootedNode) Depth(tree []RootedNode) int {
	depth := 0
	node := *r
	for node.Parent != -1 {
		depth++
		node = tree[node.Parent]
	}
	return depth
}

// Print 输出树的信息
func (r *RootedNode) Print(tree []RootedNode) {
	node := *r
	if node.Left != -1 {
		fmt.Printf("node %d: parent = %d, depth = %d, %s, %v",
			node.Value(tree),
			node.Left,
			node.DepthRec(tree),
			node.Type(),
			node.Children(tree))
	}
}

// Value return the node value in the tree.
func (r *RootedNode) Value(tree []RootedNode) int {
	if r.Parent != -1 {
		return tree[r.Parent].Left
	}
	if r.Left != -1 {
		return tree[r.Left].Parent
	}
	return 0
}

// Type return the kind of node in the tree.
func (r *RootedNode) Type() string {
	if r.Parent == -1 {
		return "root"
	}
	if r.Left == -1 {
		return "leaf"
	}
	return "internal node"
}

// Children return the children of the node.
func (r *RootedNode) Children(tree []RootedNode) []RootedNode {
	if r.Left == -1 {
		return nil
	}
	var (
		p   RootedNode
		res []RootedNode
	)
	p = tree[r.Left]
	for p.Right != -1 {
		res = append(res, p)
		p = tree[p.Right]
	}
	return res
}
