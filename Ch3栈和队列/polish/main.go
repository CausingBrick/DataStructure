// This package is the excise for polish reprentation.

package main

import (
	"fmt"
	"os"
	"strconv"

	s "github.com/CausingBrick/DataStructure/Ch3栈和队列/stack"
)

func main() {
	infix := os.Args[1:]
	if len(infix) == 0 {
		return
	}

	var stack s.Stack
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
