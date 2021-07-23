package main

import (
	"container/list"
)

type element struct {
	// Left and right pointers in binary tree.
	left, right *element
	//if value is ture represent thread, or child
	ltag, rtag bool

	// The value stored with this element.
	Value interface{}
}

func preOrderThreading(e *element) {
	if e == nil {
		return
	}
	s := list.New()
	s.PushBack(e)
	var pre *element = nil
	for s.Len() != 0 {
		cur := s.Remove(s.Back()).(*element)

		if cur.right != nil {
			cur.rtag = false
			s.PushBack(cur.right)
		} else {
			cur.rtag = true
			if next := s.Back(); next != nil {
				cur.right = next.Value.(*element)
			} else {
				cur.right = nil
			}
		}
		if cur.left != nil {
			cur.ltag = false
			s.PushBack(cur.left)
		} else {
			cur.ltag = true
			cur.left = pre
		}
		pre = cur
	}
}

func inOrderThreading(e *element) {
	s := list.New()
	var pre *element = nil
	cur := e
	for cur != nil || s.Len() != 0 {
		if cur != nil {
			s.PushBack(cur)
			cur = cur.left
		} else {
			cur = s.Remove(s.Back()).(*element)
			if left := cur.left; left != nil {
				cur.ltag = false
			} else {
				cur.ltag = true
				cur.left = pre
			}
			right := cur.right
			if right != nil {
				cur.rtag = false

			} else {
				cur.rtag = true
				//Determine whether the successor has a node
				if next := s.Back(); next != nil {
					cur.right = next.Value.(*element)
				} else {
					cur.right = nil
				}
			}
			cur = right
		}
	}
}
