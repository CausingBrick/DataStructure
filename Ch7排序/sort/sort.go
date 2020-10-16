package sort

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

// InsertionSort 插入排序
// a为待排序slice的地址
func InsertionSort(a []int) {
	var j, val int
	for i := 1; i < len(a); i++ {
		val = a[i]
		for j = i; j > 0 && a[j-1] > val; j-- {
			a[j] = a[j-1]
		}
		a[j] = val
	}
}

// BinaryInsertionSort 折半插入排序
func BinaryInsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		val := a[i]
		low, high := 0, i
		//find location high+1 to insert val
		for low < high {
			mid := (low + high) / 2
			if val < a[mid] {
				high = mid
			} else {
				low = mid + 1
			}
		}
		for j := i; j > high; j-- {
			a[j] = a[j-1]
		}
		a[high] = val
	}
}

// BubbleSort 冒泡排序
func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// BubbleSortF flag 冒泡排序
// 改进优化BubbleSort
func BubbleSortF(a []int) {
	for flag := true; flag; {
		flag = false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				flag = true
			}
		}
	}
}

// SelectionSort 选择排序
func SelectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}
}

// ShellSort 希尔排序
// gap 增量为3*gap+1
func ShellSort(a []int) {
	// Generate incremental series for  3*i+1
	var gap = []int{1}
	for i := 0; 3*gap[i]+1 < len(a); i++ {
		gap = append(gap, 3*gap[i]+1)
	}

	for g := len(gap) - 1; g >= 0 && gap[g] > 0; g-- {
		for i := gap[g]; i < len(a); i++ {
			temp := a[i]
			j := i
			for j >= gap[g] && a[j-gap[g]] > temp {
				a[j] = a[j-gap[g]]
				j -= gap[g]
			}
			a[j] = temp
		}
	}
}
