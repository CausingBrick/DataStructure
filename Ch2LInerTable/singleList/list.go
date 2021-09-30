package singleList

type Element struct {
	next  *Element
	Value interface{}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	return e.next
}

type List struct {
	root *Element
	len  int
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.root = nil
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *List { return new(List).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {

	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v ...interface{}) {
	for i := range v {
		e := &Element{l.Front(), v[i]}
		l.root = e
		l.len++
	}
}

// // PushBack inserts a new element e with value v at the back of list l and returns e.
// func (l *List) PushBack(v interface{}) *Element {
// 	l.lazyInit()
// 	return l.insertValue(v, l.root.prev)
// }
