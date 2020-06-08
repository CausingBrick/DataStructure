// This package is the excise for polish reprentation.

package main

import (
	"fmt"
	"os"
	"strconv"
)

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
func main() {
	infix := os.Args[1:]
	if len(infix) == 0 {
		return
	}
	var stack StackArr
	for i := 0; i < len(infix); i++ {

		switch infix[i] {
		case "+":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			stack.Push(a + b)
		case "-":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			stack.Push(a - b)
		case "*":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			stack.Push(a * b)
		case "/":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			stack.Push(a / b)
		default:
			a, _ := strconv.Atoi(infix[i])
			stack.Push(a)
		}
	}
	fmt.Println("result: ", stack.Pop())
}

/*
!+ textoutput

$ go run main.go 1 2 +  3 4  - -

result:  -2

!- textoutput

*/
