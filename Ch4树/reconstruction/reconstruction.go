package reconstruction

import "fmt"

// Reconstruction return the post order by given pre and inorder.
func Reconstruction(pre, in []int) {
	if len(pre) != len((in)) {
		fmt.Printf("error pre or in")
	}
	var i int
	reconstruction(0, len(pre), pre, in, &i)
}

func reconstruction(left, right int, pre, in []int, i *int) {
	if left >= right {
		return
	}
	midVal := pre[*i]
	*i++
	m := index(in, midVal)
	reconstruction(left, m, pre, in, i)
	reconstruction(m+1, right, pre, in, i)
	fmt.Println(midVal)
}

func index(a []int, v int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == v {
			return i
		}
	}
	return -1
}
