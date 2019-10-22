package main

import "fmt"

func main() {
	var as = []int{6, 2, 3, 8, 1, 4, 7, 5}
	quickSort(as, 0, len(as)-1)
	fmt.Println(as)
}

func quickSort(arr []int, p, r int) {
	if p < r {
		m := partition(arr, p, r)
		quickSort(arr, p, m-1)
		quickSort(arr, m+1, r)
	}
}

func partition(arr []int, p, r int) int {
	x := arr[r]
	i := p - 1
	for j := p; j < r; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
