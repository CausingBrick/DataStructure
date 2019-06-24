package sort

// InsertionSort 插入排序
// a为待排序slice的地址 ,n为长度
func InsertionSort(a []int) {
	var j, val int
	for i := 1; i < len(a); i++ {
		val = a[i]
		for j = i - 1; j >= 0 && a[j] > val; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = val
	}
}

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
