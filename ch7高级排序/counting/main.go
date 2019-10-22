package main

import "fmt"

func main() {
	var as = []int{6, 2, 3, 8, 1, 4, 6, 5, 7, 5}
	coutingSort(as)
	fmt.Println(as)
}

func coutingSort(arr []int) []int {
	temp := make([]int, findMaX(arr)+1)
	res := arr[:]
	for _, v := range arr {
		temp[v]++
	}
	k := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			res[k] = i
			k++
		}
	}
	return res
}

func findMaX(arr []int) int {
	var m int
	for _, v := range arr {
		if m < v {
			m = v
		}
	}
	return m
}

/*
res:
[1 2 3 4 5 5 6 6 7 8]

*/
