package main

import "fmt"

func main() {
	var as = []int{6, 2, 3, 8, 1, 4, 6, 5, 7, 5}
	fmt.Println(mergeSort(as, 0, len(as)-1))
}

// 返回逆序
func mergeSort(arr []int, l, r int) int {
	if l < r {
		mid := (l + r) / 2
		v1 := mergeSort(arr, l, mid)
		v2 := mergeSort(arr, mid+1, r)
		v3 := merge(arr, l, mid, r)
		return v1 + v2 + v3
	}
	return 0
}

// 返回逆序
func merge(arr []int, l, mid, r int) int {
	al := make([]int, mid-l+1)
	ar := make([]int, r-mid)

	copy(al, arr[l:mid+1])
	copy(ar, arr[mid+1:r+1])
	rev := 0
	var i, j int
	for i < len(al) && j < len(ar) {
		if al[i] <= ar[j] {
			arr[l+j+i] = al[i]
			i++
		} else {
			arr[l+j+i] = ar[j]
			j++
			rev++
		}
	}
	for i < len(al) {
		arr[l+j+i] = al[i]
		i++
	}
	for j < len(ar) {
		arr[l+j+i] = ar[j]
		j++
		rev++
	}
	return rev
}

/*
res:
15
*/
