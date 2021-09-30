package binarySearch

import "testing"

var lessInt = func(i, j interface{}) bool { return i.(int) < j.(int) }

func TestInsert(t *testing.T) {
	bst := New()
	for i := 0; i < 10; i++ { //test increase order
		bst.Insert(i, func(i, j interface{}) bool { return i.(int) < j.(int) })
	}
	e := bst.Element
	i := 0
	for e != nil {
		if e.Value.(int) != i {
			t.Errorf("got %d, want %d", bst.Value, i)
		}
		e = e.right
		i++
	}
}

func TestInsert1(t *testing.T) {
	bst := New()
	for i := 10; i > 0; i-- { // test decrease order
		bst.Insert(i, func(i, j interface{}) bool { return i.(int) < j.(int) })
	}
	e := bst.Element
	i := 10
	for e != nil {
		if e.Value.(int) != i {
			t.Errorf("got %d, want %d", bst.Value, i)
		}
		e = e.right
		i--
	}
}

func TestSearch(t *testing.T) {
	bst := New()
	bst.Insert(2, lessInt)
	bst.Insert(1, lessInt)
	bst.Insert(4, lessInt)
	bst.Insert(3, lessInt)
	if bst.Len() != 4 {
		t.Errorf("error len:%d, want: %d", bst.Len(), 4)
	}
	if e := bst.Search(2, lessInt); e == nil {
		t.Errorf("cant find root, want: %d", 2)
	}
	if e := bst.Search(1, lessInt); e == bst.Search(1, lessInt).Left() {
		t.Errorf("cant find root.left, want: %d", 1)
	}
	if e := bst.Search(4, lessInt); e == bst.Search(1, lessInt).Right() {
		t.Errorf("cant find root.right, want: %d", 4)
	}
}
func TestRemove(t *testing.T) {
	bst := New()
	bst.Insert(2, lessInt)
	bst.Insert(1, lessInt)
	bst.Insert(4, lessInt)
	bst.Insert(3, lessInt)

	if bst.Remove(3, lessInt); bst.Search(4, lessInt).Left() != nil {
		t.Errorf("error remove: %d", 3)
	}
	if bst.Remove(4, lessInt); bst.Search(2, lessInt).Right() != nil {
		t.Errorf("error remove: %d", 4)
	}
	if bst.Remove(1, lessInt); bst.Search(2, lessInt).Left() != nil {
		t.Errorf("error remove: %d", 1)
	}
	//test remove 
	if bst.Remove(2, lessInt); bst.Element != nil {
		t.Errorf("error remove: %d", 2)
	}
	// test len
	if bst.Len() != 0 {
		t.Errorf("error len, got: %d, want: %d", bst.Len(), 0)
	}
	// test remove nil
	if bst.Remove(5, lessInt); bst.Element != nil {
		t.Errorf("error remove: %d", 2)
	}

	bst.Insert(2, lessInt)
	bst.Insert(1, lessInt)
	bst.Insert(4, lessInt)
	bst.Insert(3, lessInt)
	if bst.Remove(4, lessInt); bst.Search(2, lessInt).Right() != bst.Search(3, lessInt) {
		t.Errorf("error remove: %d", 4)
	}
}
