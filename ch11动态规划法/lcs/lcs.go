package main

import "fmt"

func main() {
	x, y := "hello", "hloe"
	n := lcs(x, y)
	fmt.Println(n)
}

func lcs(x, y string) (length int) {
	if len(x) == 0 || len(y) == 0 {
		return 0
	}
	var c [][]int
	temp := make([]int, len(x)+1)
	for i := len(y); i >= 0; i-- {
		c = append(c, temp)
	}
	for i := 1; i < len(x); i++ {
		for j := 1; j < len(y); j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = c[j-1][i-1] + 1
			} else {
				c[i][j] = max(c[i-1][j], c[i][j-1])
			}
			length = max(length, c[i][j])
		}
	}
	return length
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
