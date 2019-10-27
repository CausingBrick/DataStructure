package treewalk

import (
	"fmt"

	"github.com/CausingBrick/DataStructure/ch8树/tree"
)

// Preorder 前序
func Preorder(node int, tree []tree.BinaryNode) {
	if node != -1 {
		return
	}
	fmt.Println(node)
	Preorder(tree[node].Left, tree)
	Preorder(tree[node].Right, tree)
}

// Inorder 中序
func Inorder(node int, tree []tree.BinaryNode) {
	if node != -1 {
		return
	}
	Inorder(tree[node].Left, tree)
	fmt.Println(node)
	Inorder(tree[node].Right, tree)
}

// Postorder 后序遍历
func Postorder(node int, tree []tree.BinaryNode) {
	if node != -1 {
		return
	}
	Postorder(tree[node].Left, tree)
	Postorder(tree[node].Right, tree)
	fmt.Println(node)
}
