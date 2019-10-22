package main

import "fmt"

var arr = []int{6, 2, 3, 8, 5, 1, 7, 4, 0}

func main() {
	mergeSort(0, len(arr)-1)
}

func mergeSort(l, r int) {
	if l < r {
		mid := (l + r) / 2
		mergeSort(l, mid)
		mergeSort(mid+1, r)
		merge(l, mid, r)
	}
}

func merge(l, mid, r int) {
	al := make([]int, mid-l+1)
	ar := make([]int, r-mid)

	copy(al, arr[l:mid+1])
	copy(ar, arr[mid+1:r+1])

	var i, j int
	for i < len(al) && j < len(ar) {
		if al[i] <= ar[j] {
			arr[l+j+i] = al[i]
			i++
		} else {
			arr[l+j+i] = ar[j]
			j++
		}
	}
	for i < len(al) {
		arr[l+j+i] = al[i]
		i++
	}
	for j < len(ar) {
		arr[l+j+i] = ar[j]
		j++
	}
	fmt.Println(arr)
}

/*
[2 6 3 8 5 1 7 4 0]
[2 3 6 8 5 1 7 4 0]
[2 3 6 5 8 1 7 4 0]
[2 3 5 6 8 1 7 4 0]
[2 3 5 6 8 1 7 4 0]
[2 3 5 6 8 1 7 0 4]
[2 3 5 6 8 0 1 4 7]
[0 1 2 3 4 5 6 7 8]
*/
