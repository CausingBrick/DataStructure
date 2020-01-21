package main

import "fmt"

var data []int

func fibo(n int) int {
	if n == 0 || n == 1 {
		if data[n] == 0 {
			data[n] = 1
		}
		return data[n]
	}
	if data[n] == 0 {
		data[n] = fibo(n-1) + fibo(n-2)
	}
	return data[n]
}

func main() {
	n := 3
	data = make([]int, n+1)
	fmt.Println(fibo(n))
}
