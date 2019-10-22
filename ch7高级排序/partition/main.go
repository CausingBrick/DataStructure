package main

func main() {
	var as = []int{6, 2, 3, 8, 1, 4, 7, 5}
	mid := partition(as, 0, len(as)-1)

	for i := 0; i < len(as); i++ {
		if i == mid {

			print("|", as[i], "| ")

		} else {
			print(as[i], " ")
		}
	}
}

func partition(arr []int, p, r int) int {
	x := arr[r]
	i := p - 1
	for j := p; j < r; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

/*
2 3 1 4 |5| 8 7 6
*/
