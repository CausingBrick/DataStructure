package tree

import (
	"testing"
)

var tree []RootedNode

func init() {
	/*
		0
		1 	3
		3	4	5
		6
	*/

	tree = make([]RootedNode, 7)
	tree[0].Parent = -1
	tree[0].Left = 1
	tree[0].Right = -1

	tree[1].Parent = 0
	tree[1].Left = 3
	tree[1].Right = 3

	tree[2].Parent = 1
	tree[2].Left = 5
	tree[2].Right = -1

	tree[3].Parent = 1
	tree[3].Left = 6
	tree[3].Right = 4

	tree[4].Parent = 1
	tree[4].Left = -1
	tree[4].Right = -1

	tree[5].Parent = 3
	tree[5].Left = -1
	tree[5].Right = -1

	tree[6].Parent = 3
	tree[6].Left = -1
	tree[6].Right = -1

}

func TestDepth(t *testing.T) {
	if tree[0].Depth(tree) != 0 {
		t.Errorf("root node's node = %d, want 0", tree[0].Depth(tree))
	}
	if tree[6].Depth(tree) != 3 {
		t.Errorf("root node's node = %d, want 3", tree[6].Depth(tree))
	}
}

func TestDepthRec(t *testing.T) {
	if tree[0].DepthRec(tree) != 0 {
		t.Errorf("root node's node = %d, want 0", tree[0].DepthRec(tree))
	}
	if tree[6].DepthRec(tree) != 3 {
		t.Errorf("root node's node = %d, want 3", tree[6].DepthRec(tree))
	}
}

func TestPrint(t *testing.T) {
	tree[0].Print(tree)
}
func TestValue(t *testing.T) {
	if tree[1].Value(tree) != 1 {
		t.Errorf("node's value = 1, want %d", tree[0].Value(tree))
	}
	if tree[0].Value(tree) != 0 {
		t.Errorf("node's value = 0, want  %d", tree[1].Value(tree))
	}
	node := RootedNode{-1, -1, -1}
	if node.Value(nil) != 0 {
		t.Errorf("node's value = 0, want  %d", node.Value(nil))
	}
}

func TestType(t *testing.T) {
	if tree[0].Type() != "root" {
		t.Errorf("node's type = '%s', want 'root'", tree[0].Type())
	}
	if tree[1].Type() != "internal node" {
		t.Errorf("node's type = '%s', want 'internal node'", tree[1].Type())
	}
	if tree[6].Type() != "leaf" {
		t.Errorf("node's type = '%s', want 'leaf'", tree[6].Type())
	}
}

func TestChildren(t *testing.T) {
	if tree[6].Children(tree) != nil {
		t.Errorf("node's type = %v, want nil", tree[6].Children(tree))
	}
	res := tree[0].Children(tree)
	t.Logf("%v", res)
	if res[0] != tree[1] || res[1] != tree[3] {
		want := make([]RootedNode, 2)
		want[0] = tree[1]
		want[1] = tree[3]
		t.Errorf("node's type = %v, want %v", tree[6].Children(tree), want)
	}
}
