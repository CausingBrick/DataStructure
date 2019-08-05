package binary

//NOTFOUND is the not found code.
const NOTFOUND = -404

// Search 二分搜索
func Search(key int, seq []int) int {
	left, right := 0, len(seq)
	mid := (left + right) / 2
	for left < right {
		if key == seq[mid] {
			return mid
		} else if key < mid {
			right = mid
		} else {
			left = mid
		}
	}
	return NOTFOUND
}
