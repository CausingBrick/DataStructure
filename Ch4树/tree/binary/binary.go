package main

import (
	"container/list"
	"fmt"
)

// var NOMOREVALUE = errors.New("no more value")

// Element is an element of a binary tree.
type Element struct {
	// Left and right pointers in binary tree.
	left, right *Element

	// The tree to which this element belongs.
	tree *BinTree

	// The value stored with this element.
	Value interface{}
}

// Left returns the left child or nil.
func (e *Element) Left() *Element {
	if e.left == nil {
		return nil
	}
	return e.left
}

// Right returns the right child or nil.
func (e *Element) Right() *Element {
	if e.right == nil {
		return nil
	}
	return e.right
}

// BinTree represents a binary tree.
// The zero value for tree is an empty tree ready to use.
type BinTree struct {
	root *Element
	// tail *Element
	// len  int
}

// Init initializes a tree.
func (b *BinTree) Init() *BinTree {
	// b.root.left = nil
	// b.root.right = nil
	// b.tail = b.root
	// b.len = 0
	return b
}

// New return a empty tree
func New() *BinTree { return new(BinTree).Init() }

func (b *BinTree) getLevelTraseValues(val ...interface{}) {
	vlen := len(val)
	if vlen == 0 {
		return
	}

	b.root = &Element{nil, nil, b, val[0]}
	if vlen == 1 {
		return
	}

	var queue = list.New()
	queue.PushBack(b.root)
	for e, i := b.root, 1; ; i += 2 {
		e = queue.Remove(queue.Front()).(*Element)

		if vlen < i {
			e.left = &Element{nil, nil, b, val[i]}
			queue.PushBack(e.left)
		}
		if vlen < i+1 {
			e.right = &Element{nil, nil, b, val[i+1]}
			queue.PushBack(e.right)
		}
		if queue.Len() == 0 {
			return
		}
	}
}

// Depth return the depth
func (b *BinTree) Depth() int {
	var dp int
	for e := b.root; e != nil; e = e.left {
		dp++
	}
	return dp
}

// ActionInTraverse contains the funcion for elelment's operiaton.
type ActionInTraverse func(e *Element)

func preorderTraverse(e *Element, action ActionInTraverse) {
	if e == nil {
		return
	}
	//do something
	action(e)
	preorderTraverse(e.left, action)
	preorderTraverse(e.right, action)
}

func preorderTraverseIteration(e *Element, action ActionInTraverse) {
	if e == nil {
		return
	}
	s := list.New()
	s.PushBack(e)
	for s.Len() != 0 {
		top := s.Remove(s.Back()).(*Element)
		action(top)
		if top.right != nil {
			s.PushBack(top.right)
		}
		if top.left != nil {
			s.PushBack(top.left)
		}
	}
}

func inorderTraverse(e *Element, action ActionInTraverse) {
	if e == nil {
		return
	}
	inorderTraverse(e.left, action)
	//do something
	action(e)
	inorderTraverse(e.right, action)
}

func inorderTraverseIteration(e *Element, action ActionInTraverse) {
	p := e
	s := list.New()
	for p != nil || s.Len() != 0 {
		if p != nil {
			s.PushBack(p)
			p = p.left
		} else {
			q := s.Remove(s.Back()).(*Element)
			action(q)
			p = q.right
		}
	}
}

func postorderTraverse(e *Element, action ActionInTraverse) {
	if e == nil {
		return
	}
	postorderTraverse(e.left, action)
	postorderTraverse(e.right, action)
	//do something
	action(e)
}

// use double stack
func postorderTraverseIteration1(e *Element, action ActionInTraverse) {
	if e == nil {
		return
	}
	s := list.New()
	r := list.New()
	s.PushBack(e)
	for s.Len() != 0 {
		cur := s.Remove(s.Back()).(*Element)
		r.PushBack(cur)
		if cur.left != nil {
			s.PushBack(cur.left)
		}
		if cur.right != nil {
			s.PushBack(cur.right)
		}
	}
	for r.Len() > 0 {
		action(r.Remove(r.Back()).(*Element))
	}
}

// use singel stack
func postorderTraverseIteration2(e *Element, action ActionInTraverse) {
	cur := e
	var pre *Element
	s := list.New()
	for cur != nil || s.Len() != 0 {
		for cur != nil {
			s.PushBack(cur)
			cur = cur.left
		}
		if s.Len() != 0 {
			cur = s.Back().Value.(*Element)
			if cur.right == nil || pre == cur.right {
				action(cur)
				s.Remove(s.Back())
				pre = cur
				cur = nil
			} else {
				cur = cur.right
			}
		}
	}
}
func levelorderTraverse(b *BinTree, action ActionInTraverse) {
	if b.root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(b.root)
	for e := b.root; queue.Len() > 0; {
		e = queue.Remove(queue.Front()).(*Element)
		action(e)
		if e.left != nil {
			queue.PushBack(e.left)
		}
		if e.right != nil {
			queue.PushBack(e.right)
		}
	}
}



func main() {
	t1 := New()
	t1.root = &Element{nil, nil, t1, 1}
	t1.root.left = &Element{nil, nil, t1, 2}
	t1.root.left.left = &Element{nil, nil, t1, 3}
	t1.root.left.right = &Element{nil, nil, t1, 4}
	t1.root.right = &Element{nil, nil, t1, 5}
	t1.root.right.left = &Element{nil, nil, t1, 6}
	t1.root.right.right = &Element{nil, nil, t1, 7}

	printTree := func(e *Element) {
		fmt.Printf("%v ", e.Value)
	}
	preorderTraverse(t1.root, printTree)
	fmt.Println()
	preorderTraverseIteration(t1.root, printTree)
	fmt.Println()

	inorderTraverse(t1.root, printTree)
	fmt.Println()
	inorderTraverseIteration(t1.root, printTree)
	fmt.Println()

	postorderTraverse(t1.root, printTree)
	fmt.Println()
	postorderTraverseIteration1(t1.root, printTree)
	fmt.Println()
	postorderTraverseIteration2(t1.root, printTree)
	fmt.Println()

	fmt.Println(t1.Depth())
	levelorderTraverse(t1, printTree)

}
