/*
列举组合的递归函数
*/
package main

import "fmt"

var arr = [3]int{}

func main() {
	comb(0)
}

func comb(i int) {
	if i >= len(arr) {
		fmt.Println(arr)
		return
	}
	comb(i + 1)
	arr[i] = 1 //表示i被选中
	comb(i + 1)
	arr[i] = 0 //表示i未被选中
}

/* res:

[0 0 0]
[0 0 1]
[0 1 0]
[0 1 1]
[1 0 0]
[1 0 1]
[1 1 0]
[1 1 1]

*/
