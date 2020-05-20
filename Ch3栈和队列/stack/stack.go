// Package stack Encapsulate a type int stack made up by slices
package stack

// Stack is the common type for stack.
type Stack struct {
	stack []interface{}
}

// Pop return the top value of stack.
// The length of stack will decrease.
func (s *Stack) Pop() interface{} {
	// fmt.Println("len:", len(s.stack))
	popVal := s.stack[len(s.stack)-1]
	temp := make([]interface{}, len(s.stack)-1)
	copy(temp, s.stack)
	s.stack = temp
	return popVal
}

// Push will increase one unit of stack.
// The x will be the top of stack.
func (s *Stack) Push(x interface{}) {
	s.stack = append(s.stack, x)
}

// IsEmpty return the status for stack if it is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
