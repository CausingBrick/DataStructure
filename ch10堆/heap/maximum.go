package main

import "fmt"

func main() {
	var a = make([]int, 10)
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	fmt.Println(a)
	for i := len(a) / 2; i >= 0; i-- {
		MaxHeapify(a, i)
	}
	fmt.Println(a)
}

// MaxHeapify 数组最大堆化处理.i为传入的根节点
func MaxHeapify(a []int, i int) {
	l, r := 2*i, 2*i+1
	var largest int
	if l < len(a) && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		MaxHeapify(a, largest)
	}

}
