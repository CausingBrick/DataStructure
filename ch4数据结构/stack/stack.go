package stack

import (
	"fmt"
)

// Stack is the common type for stack.
type Stack struct {
	// len int
	// cap int
	stack []int
	point int
}

// Initialize stack.
func (s *Stack) Initialize(cap int) {
	s.stack = make([]int, 0, cap)
	s.point = 0
}

func (s *Stack) Pop(x int) {
	s.stack[s.point] = x
	if s.point < cap(s.stack) {
		s.point++
	} else {
		fmt.Errorf("Full of stack")
	}
}
