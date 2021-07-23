package sort

import "strconv"

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// InsertionSort 插入排序
// arr为待排序slice的地址
func InsertionSort(arr []int) {
	var j, val int
	for i := 1; i < len(arr); i++ {
		val = arr[i]
		for j = i; j > 0 && arr[j-1] > val; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = val
	}
}

// BinaryInsertionSort 折半插入排序
func BinaryInsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		low, high := 0, i
		//find location high+1 to insert val
		for low < high {
			mid := (low + high) / 2
			if val < arr[mid] {
				high = mid
			} else {
				low = mid + 1
			}
		}
		for j := i; j > high; j-- {
			arr[j] = arr[j-1]
		}
		arr[high] = val
	}
}

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// BubbleSortF flag 冒泡排序
// 改进优化BubbleSort
func BubbleSortF(arr []int) {
	for flag := true; flag; {
		flag = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = true
			}
		}
	}
}

// ShellSort 希尔排序
// gap 增量为3*gap+1
func ShellSort(arr []int) {
	// Generate incremental series for  3*i+1
	var gap = []int{1}
	for i := 0; 3*gap[i]+1 < len(arr); i++ {
		gap = append(gap, 3*gap[i]+1)
	}

	for g := len(gap) - 1; g >= 0 && gap[g] > 0; g-- {
		for i := gap[g]; i < len(arr); i++ {
			val := arr[i]
			j := i
			for j >= gap[g] && arr[j-gap[g]] > val {
				arr[j] = arr[j-gap[g]]
				j -= gap[g]
			}
			arr[j] = val
		}
	}
}

// QuickSort 快速排序
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, p, r int) {
	if p < r {
		m := partition(arr, p, r)

		quickSort(arr, p, m-1)

		quickSort(arr, m+1, r)
	}
}

// partition 原址排序返回一个分割
func partition(arr []int, p, r int) int {
	val := arr[r]
	i := p - 1
	for j := p; j < r; j++ {
		if arr[j] <= val {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

// HeapSort 堆排序
func HeapSort(arr []int) {
	heapSize := len(arr)
	buildMAXHeap(arr)
	for i := heapSize - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapSize--
		maxHeapify(arr, 0, heapSize)
	}
}

// buildMAXHeap 将数组a做最大堆化处理
func buildMAXHeap(arr []int) {
	heapSize := len(arr)
	for i := heapSize >> 1; i >= 0; i-- {
		maxHeapify(arr, i, heapSize)
	}
}

// maxHeapify 维护最大堆的性质
func maxHeapify(arr []int, i, heapSize int) {
	left, right := i<<1+1, i<<1+2
	largest := i
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		maxHeapify(arr, largest, heapSize)
	}
}

// MergeSort 归并排序
func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) >> 1
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left, mid, right int) {
	leftArr, rightArr := make([]int, mid+1-left), make([]int, right-mid)
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])
	lSize, rSize := len(leftArr), len(rightArr)
	i, j := 0, 0
	for i < lSize && j < rSize {
		if leftArr[i] <= rightArr[j] {
			arr[left+i+j] = leftArr[i]
			i++
		} else {
			arr[left+i+j] = rightArr[j]
			j++
		}
	}
	for i < lSize {
		arr[left+i+j] = leftArr[i]
		i++
	}
	for j < rSize {
		arr[left+i+j] = rightArr[j]
		j++
	}
}

// RadixSort 基数排序
func RadixSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	radixSort(arr, len(strconv.Itoa(max)))
}
func radixSort(arr []int, maxDigit int) {
	mod, div := 10, 1
	for i := 0; i < maxDigit; i++ {
		counter := make([][]int, 10)
		//push data sorted by digit to bucket
		for j := 0; j < len(arr); j++ {
			bucket := (arr[j] % mod) / div
			if len(counter[bucket]) == 0 {
				counter[bucket] = make([]int, 0)
			}
			counter[bucket] = append(counter[bucket], arr[j])
		}
		//sort data from bucket to arr
		for j, pos := 0, 0; j < len(counter); j++ {
			for k := 0; k < len(counter[j]); k++ {
				arr[pos] = counter[j][k]
				pos++
			}
		}
		div *= 10
		mod *= 10
	}
}
