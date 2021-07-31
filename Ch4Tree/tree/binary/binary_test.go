package binary

import (
	"container/list"
	"testing"
)

// Breadth-first tree creation
func createTree(values ...interface{}) *BinTree {
	if len(values) == 0 {
		return nil
	}
	b := New(values[0])
	q := list.New()
	q.PushBack(&b.Element)
	for i := 1; i < len(values); {
		e := q.Remove(q.Front()).(*Element)
		e.left = NewElement(values[i])
		q.PushBack(e.left)
		i++
		if i < len(values) {
			e.right = NewElement(values[i])
			q.PushBack(e.right)
			i++
		}
	}
	return b
}

func isEqual(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestDepth(t *testing.T) {
	tb := []struct {
		arr   []interface{}
		depth int
	}{
		{[]interface{}{1}, 1},
		{[]interface{}{1, 2}, 2},
		{[]interface{}{1, 2, 3}, 2},
		{[]interface{}{1, 2, 3, 4}, 3},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 3},
	}
	var bt *BinTree
	for _, v := range tb {
		bt = createTree(v.arr...)
		if d := bt.Depth(); d != v.depth {
			t.Errorf("got %d, want %d", d, v.depth)
		}
	}
}

func TestPreorderTraverse(t *testing.T) {
	tb := []struct {
		input  []interface{}
		result []interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 2, 3, 4}, []interface{}{1, 2, 4, 3}},
		{[]interface{}{"a", "b", "c", "d", "e"}, []interface{}{"a", "b", "d", "e", "c"}},
	}
	var bt *BinTree
	for _, v := range tb {
		bt = createTree(v.input...)
		output := make([]interface{}, len(v.input))
		i := 0
		bt.PreorderTraverse(func(e *Element) {
			output[i] = e.Value
			i++
		})
		if !isEqual(output, v.result) {
			t.Errorf("got %v, want %v", output, v.result)
		}
	}
}

func TestPreorderTraverseIteration(t *testing.T) {
	tb := []struct {
		input  []interface{}
		result []interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 2, 3, 4}, []interface{}{1, 2, 4, 3}},
		{[]interface{}{"a", "b", "c", "d", "e"}, []interface{}{"a", "b", "d", "e", "c"}},
	}
	var bt *BinTree
	for _, v := range tb {
		bt = createTree(v.input...)
		output := make([]interface{}, len(v.input))
		i := 0
		bt.PreorderTraverseIteration(func(e *Element) {
			output[i] = e.Value
			i++
		})
		if !isEqual(output, v.result) {
			t.Errorf("got %v, want %v", output, v.result)
		}
	}
}

func TestInorderTraverse(t *testing.T) {
	tb := []struct {
		input  []interface{}
		result []interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 2, 3, 4}, []interface{}{4, 2, 1, 3}},
		{[]interface{}{"a", "b", "c", "d", "e"}, []interface{}{"d", "b", "e", "a", "c"}},
	}
	var bt *BinTree
	for _, v := range tb {
		bt = createTree(v.input...)
		output := make([]interface{}, len(v.input))
		i := 0
		bt.InorderTraverse(func(e *Element) {
			output[i] = e.Value
			i++
		})
		if !isEqual(output, v.result) {
			t.Errorf("got %v, want %v", output, v.result)
		}
	}
}

func TestInorderTraverseIteration(t *testing.T) {
	tb := []struct {
		input  []interface{}
		result []interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 2, 3, 4}, []interface{}{4, 2, 1, 3}},
		{[]interface{}{"a", "b", "c", "d", "e"}, []interface{}{"d", "b", "e", "a", "c"}},
	}
	var bt *BinTree
	for _, v := range tb {
		bt = createTree(v.input...)
		output := make([]interface{}, len(v.input))
		i := 0
		bt.InorderTraverseIteration(func(e *Element) {
			output[i] = e.Value
			i++
		})
		if !isEqual(output, v.result) {
			t.Errorf("got %v, want %v", output, v.result)
		}
	}
}

func TestPostorderTraverse(t *testing.T) {
	tb := []struct {
		input  []interface{}
		result []interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 2, 3, 4}, []interface{}{4, 2, 3, 1}},
		{[]interface{}{"a", "b", "c", "d", "e"}, []interface{}{"d", "e", "b", "c", "a"}},
	}
	var bt *BinTree
	for _, v := range tb {
		bt = createTree(v.input...)
		output := make([]interface{}, len(v.input))
		i := 0
		bt.PostorderTraverse(func(e *Element) {
			output[i] = e.Value
			i++
		})
		if !isEqual(output, v.result) {
			t.Errorf("got %v, want %v", output, v.result)
		}
	}
}

func TestPostorderTraverseIter(t *testing.T) {
	tb := []struct {
		input  []interface{}
		result []interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 2, 3, 4}, []interface{}{4, 2, 3, 1}},
		{[]interface{}{"a", "b", "c", "d", "e"}, []interface{}{"d", "e", "b", "c", "a"}},
	}
	var bt *BinTree
	for _, v := range tb {
		bt = createTree(v.input...)
		output := make([]interface{}, len(v.input))
		i := 0
		bt.PostorderTraverseIteration(func(e *Element) {
			output[i] = e.Value
			i++
		})
		if !isEqual(output, v.result) {
			t.Errorf("got %v, want %v", output, v.result)
		}
	}
}

func TestLevelorderTraverse(t *testing.T) {
	tb := []struct {
		input  []interface{}
		result []interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 2, 3, 4}, []interface{}{1, 2, 3, 4}},
		{[]interface{}{"a", "b", "c", "d", "e"}, []interface{}{"a", "b", "c", "d", "e"}},
	}
	var bt *BinTree
	for _, v := range tb {
		bt = createTree(v.input...)
		output := make([]interface{}, len(v.input))
		i := 0
		bt.LevelorderTraverse(func(e *Element) {
			output[i] = e.Value
			i++
		})
		if !isEqual(output, v.result) {
			t.Errorf("got %v, want %v", output, v.result)
		}
	}
}
