package main

var arr = []int{13, 89, 456, 32, 24, 45, 67, 43, 34, 29, 92}

func main() {
	println(findMAX(0, len(arr)-1))
}

func findMAX(f, l int) int {
	if f == l {
		return arr[f]
	}
	return max(findMAX(f, (l+f)/2), findMAX((f+l+1)/2, l))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
