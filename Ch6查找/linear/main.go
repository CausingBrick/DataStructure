package main

import (
	"fmt"
)

func linearSearch(a []int, key int) int {
	i := 0
	a = append(a, key)
	for a[i] != key {
		i++
	}
	if i == len(a)-1 {
		return -1
	}
	return i
}

func main() {
	var (
		S, T []int
		n    int
		num  int
	)
	fmt.Scanln(&n)
	S = make([]int, n, n+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &S[i])
	}

	fmt.Scanln(&n)
	T = make([]int, n, n+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &T[i])
	}
	fmt.Println(T)
	for _, v := range S {
		if linearSearch(T, v) != -1 {
			num++
		}
	}
	fmt.Println(num)
}
