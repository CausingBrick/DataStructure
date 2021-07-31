package binary

import (
	"container/list"
)

// Element is an element of a binary tree.
type Element struct {
	// Left and right pointers in binary tree.
	left, right *Element

	// The value stored with this element.
	Value interface{}
}

func NewElement(v interface{}) *Element {
	e := new(Element)
	e.Value = v
	return e
}

// Left returns the left child or nil.
func (e *Element) Left() *Element {
	return e.left
}

// Right returns the right child or nil.
func (e *Element) Right() *Element {
	return e.right
}

// BinTree represents a binary tree.
// The zero value for tree is an empty tree ready to use.
type BinTree struct {
	Element
}

// Init initializes a tree.
func (b *BinTree) Init() *BinTree {
	b.left = nil
	b.right = nil
	return b
}

// New returns an empty tree
func New(v interface{}) *BinTree {
	bt := new(BinTree).Init()
	bt.Value = v
	return bt
}

// Depth return the depth
func (b *BinTree) Depth() int {
	return depth(&b.Element)
}

func depth(e *Element) int {
	if e == nil {
		return 0
	}
	if l, r := depth(e.left), depth(e.right); l > r {
		return l + 1
	} else {
		return r + 1
	}
}

// PreorderTraverse
func (b *BinTree) PreorderTraverse(visit func(e *Element)) {
	preorderTraverse(&b.Element, visit)
}

func preorderTraverse(e *Element, visit func(e *Element)) {
	if e == nil {
		return
	}
	visit(e)
	preorderTraverse(e.left, visit)
	preorderTraverse(e.right, visit)
}

func (b *BinTree) PreorderTraverseIteration(visit func(e *Element)) {
	preorderTraverseIteration(&b.Element, visit)
}

func preorderTraverseIteration(e *Element, visit func(e *Element)) {
	if e == nil {
		return
	}
	s := list.New()
	s.PushBack(e)
	for s.Len() != 0 {
		top := s.Remove(s.Back()).(*Element)
		visit(top)
		if top.right != nil {
			s.PushBack(top.right)
		}
		if top.left != nil {
			s.PushBack(top.left)
		}
	}
}

func (b *BinTree) InorderTraverse(visit func(e *Element)) {
	inorderTraverse(&b.Element, visit)
}

func inorderTraverse(e *Element, visit func(e *Element)) {
	if e == nil {
		return
	}
	inorderTraverse(e.left, visit)
	visit(e)
	inorderTraverse(e.right, visit)
}

func (b *BinTree) InorderTraverseIteration(visit func(e *Element)) {
	inorderTraverseIteration(&b.Element, visit)
}

func inorderTraverseIteration(e *Element, visit func(e *Element)) {
	s := list.New()
	for {
		if e != nil {
			s.PushBack(e)
			e = e.left
		} else if s.Len() != 0 {
			e = s.Remove(s.Back()).(*Element)
			visit(e)
			e = e.right
		} else {
			break
		}
	}
}

func (b *BinTree) PostorderTraverse(visit func(e *Element)) {
	postorderTraverse(&b.Element, visit)
}

func postorderTraverse(e *Element, visit func(e *Element)) {
	if e == nil {
		return
	}
	postorderTraverse(e.left, visit)
	postorderTraverse(e.right, visit)
	visit(e)
}

func (b *BinTree) PostorderTraverseIteration(visit func(e *Element)) {
	postorderTraverseIteration(&b.Element, visit)
}

func postorderTraverseIteration(e *Element, visit func(e *Element)) {
	if e == nil {
		return
	}
	s := list.New()
	s.PushBack(e)
	for e.right != nil || e.left != nil {
		e = s.Back().Value.(*Element)
		if e.right != nil {
			s.PushBack(e.right)
		}
		if e.left != nil {
			s.PushBack(e.left)
		}
	}
	for s.Len() != 0 {
		visit(s.Remove(s.Back()).(*Element))
	}
}
func (b *BinTree) LevelorderTraverse(visit func(e *Element)) {
	levelorderTraverse(&b.Element, visit)
}
func levelorderTraverse(e *Element, visit func(e *Element)) {
	if e == nil {
		return
	}
	q := list.New()
	q.PushBack(e)
	for q.Len() != 0 {
		t := q.Remove(q.Front()).(*Element)
		visit(t)
		if t.left != nil {
			q.PushBack(t.left)
		}
		if t.right != nil {
			q.PushBack(t.right)
		}
	}
}
