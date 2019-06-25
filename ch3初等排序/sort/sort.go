package sort

// InsertionSort 插入排序
// a为待排序slice的地址
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

// BubbleSort 冒泡排序
func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

// BubbleSortF flag 冒泡排序
// 改进优化BubbleSort
func BubbleSortF(a []int) {
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				flag = true
			}
		}
	}
}
