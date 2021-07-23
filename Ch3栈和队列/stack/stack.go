package main

import "fmt"

// StackArr is the common type for stack.
type StackArr struct {
	stack []interface{}
}

// Pop return the top value of stack.
// The length of stack will decrease.
func (s *StackArr) Pop() interface{} {
	if len(s.stack) < 1 {
		return nil
	}
	popVal := s.stack[len(s.stack)-1]
	temp := make([]interface{}, len(s.stack)-1)
	copy(temp, s.stack)
	s.stack = temp
	return popVal
}

// Push will increase one unit of stack.
// The x will be the top of stack.
func (s *StackArr) Push(x ...interface{}) {
	for i := 0; i < len(x); i++ {
		s.stack = append(s.stack, x[i])
	}
}

// IsEmpty return the status for stack if it is empty.
func (s *StackArr) IsEmpty() bool {
	return len(s.stack) == 0
}

// Init a stack with cap 10.
func (s *StackArr) Init() {
	s.stack = make([]interface{}, 0, 10)
}

// Len return the length of stack.
func (s *StackArr) Len() int {
	return len(s.stack)
}

// Show return the length of stack.
func (s *StackArr) Show() {
	for _, val := range s.stack {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

// Element is the element of stack.
type Element struct {
	next, prev *Element
	Value      interface{}
}

// Stack represents a stack implement by double linked list.
type Stack struct {
	// The head stores the tail with head.next and top with head.prev.
	head Element
	// Current stack lenght exclude (this) sintinel element.
	len int
}

// Init initializes stack or clears s.
func (s *Stack) Init() *Stack {
	s.head.next = &s.head
	s.head.prev = &s.head
	s.len = 0
	return s
}

// Push pushs a element to stack.
func (s *Stack) Push(v ...interface{}) {
	for i := 0; i < len(v); i++ {
		e := &Element{&s.head, s.head.prev, v[i]}
		s.head.prev.next = e
		s.head.prev = e
		s.len++
	}
}

// Pop removes the top element of stack and return e.
func (s *Stack) Pop() *Element {
	if s.len == 0 {
		return nil
	}
	current := s.head.prev
	s.head.prev = current.prev
	s.head.prev.next = &s.head
	current.prev, current.next = nil, nil
	s.len--
	return current
}

// Show prints all the element's value.
func (s *Stack) Show() {
	if s.len == 0 {
		return
	}
	for e := s.head.next; e != &s.head; e = e.next {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

// IsEmpty return if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func main() {
	// Test stack implemented by array
	var s1 StackArr
	s1.Push(1, "me", 5, 6, 7)
	s1.Show() //1 me 5 6 7
	s1.Pop()
	s1.Show() //1 me 5 6
	s1.Push("u")
	s1.Show()                                     //1 me 5 6 u
	fmt.Println("Is stack null? :", s1.IsEmpty()) //Is stack null? : false
	s1.Init()
	fmt.Println("Is stack null? :", s1.IsEmpty()) //Is stack null? : true

	// Test stack implemented by list
	var s2 = new(Stack).Init()
	s2.Push(1, "me", 5, 6, 7)
	s2.Show() //1 me 5 6 7
	s2.Pop()
	s2.Show() //1 me 5 6
	s2.Push("u")
	s2.Show()                                     //1 me 5 6 u
	fmt.Println("Is stack null? :", s2.IsEmpty()) //Is stack null? : false
	s2.Init()
	fmt.Println("Is stack null? :", s2.IsEmpty()) //Is stack null? : true

}
