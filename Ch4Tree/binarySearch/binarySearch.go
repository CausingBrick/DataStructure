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
	for e != nil && e.Value.(int) != v.(int) {
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
	// find e and parent p
	for e != nil && e.Value != v {
		p = e
		if less(v, e.Value) {
			e = e.left
		} else {
			e = e.right
		}
	}
	// remove e
	switch {
	case e == nil:
		// empty tree
		break
	case e == b.Element:
		b.Element = nil
	case e.left != nil && e.right != nil:
		// left and right subtrees exist
		s, ps := e, e
		// find the max element s in left subtree
		for s.left != nil && s.right != nil {
			ps = s
			if s.right != nil {
				s = s.right
			} else {
				s = s.left
			}
		}
		// remove s form subtree
		if s == ps.left {
			ps.left = nil
		} else {
			ps.right = nil
		}
		// replace e with s
		s.left, s.right = e.left, e.right
		if e == p.left {
			p.left = s
		} else {
			p.right = s
		}

	case e.left != nil:
		if p.left == e {
			p.left = e.left
		} else {
			p.right = e.left
		}
	case e.right != nil:
		if p.left == e {
			p.left = e.right
		} else {
			p.right = e.right
		}
	default:
		if p.left == e {
			p.left = nil
		} else {
			p.right = nil
		}

	}
	if e != nil {
		e.left, e.right = nil, nil // avoid memory leaks
		b.len--
	}
	return e
}
