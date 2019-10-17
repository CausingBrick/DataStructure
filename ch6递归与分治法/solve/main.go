package main

var arr = []int{13, 89, 456, 32, 24, 45, 67, 43, 34, 29, 92}

func main() {
	println(solve(0, arr[9]+arr[1]))
}

func solve(i, m int) bool {
	if i >= len(arr) || m < 0 {
		return false
	}
	if m == 0 {
		return true
	}

	return solve(i+1, m) || solve(i+1, m-arr[i])
}
