package binarySearch

type Element struct {
	Value       interface{}
	left, right *Element
}

func (e *Element) Left() *Element {
	return e.left
}
func (e *Element) Right() *Element {
	return e.right
}

type BinarySearch struct {
	*Element
	len int
}

func New() *BinarySearch {
	return new(BinarySearch)
}
func (b *BinarySearch) Len() int {
	return b.len
}

func (b *BinarySearch) Insert(v interface{}, less func(i, j interface{}) bool) *Element {
	if b.Element == nil {
		b.Element = &Element{Value: v}
		b.len++
		return b.Element
	}
	// find the position to be inserted
	p, e := b.Element, b.Element
	for e != nil {
		p = e
		switch {
		case less(v, e.Value):
			e = e.left
		default:
			e = e.right
		}
	}
	// insert
	e = &Element{Value: v}
	if less(v, p.Value) {
		p.left = e
	} else {
		p.right = e
	}
	b.len++
	return e
}

func (b *BinarySearch) Search(v interface{}, less func(i, j interface{}) bool) *Element {
	e := b.Element
	for e != nil && e.Value != v {
		if less(v, e.Value) {
			e = e.left
		} else {
			e = e.right
		}
	}
	return e
}

func (b *BinarySearch) Remove(v interface{}, less func(i, j interface{}) bool) *Element {
	p, e := b.Element, b.Element
	// find e and its parent p
	for e != nil && e.Value != v {
		p = e
		if less(v, e.Value) {
			e = e.left
		} else {
			e = e.right
		}
	}
	// remove e
	// empty tree or e not found.
	switch {
	case e == nil:
		break
	// e is leaf node.
	case e.left == nil && e.right == nil:
		// just root node exists.
		if b.Element == e {
			b.Element = nil
			break
		}
		if p.left == e {
			p.left = nil
		} else {
			p.right = nil
		}
	// just left node exists.
	case e.left != nil && e.right == nil:
		if p.left == e {
			p.left = e.left
		} else {
			p.right = e.left
		}
	// just right node exists.
	case e.left == nil && e.right != nil:
		if p.left == e {
			p.left = e.right
		} else {
			p.right = e.right
		}
	// left and right subtrees are existed.
	default:
		// use the pre node in inorder traverse to replace the e.
		pre, ppre := e.left, e
		for pre.right != nil {
			ppre = pre
			pre = pre.right
		}
		if pre != ppre {
			if ppre.left == pre {
				ppre.left = nil
			} else {
				ppre.right = nil
			}
		}
		if p.left == e {
			p.left = pre
		} else {
			p.right = pre
		}
	}
	if e != nil {
		e.left, e.right = nil, nil // avoid memory leaks
		b.len--
	}
	return e
}
