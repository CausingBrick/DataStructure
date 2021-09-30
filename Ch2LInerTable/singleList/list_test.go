package singleList

import (
	"testing"
)

func TestPushFront(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	l := New()
	l.PushFront(v)
	for e, i := l.Front(), 0; e != nil; e = e.Next() {
		// do something with e.Value
		if i != l.Len() || e.Value != v[i] {
			t.Errorf("error")
		}
	}

}
