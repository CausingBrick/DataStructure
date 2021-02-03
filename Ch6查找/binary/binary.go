package binary

//NOTFOUND is the not found code.
const NOTFOUND = -404

// Search 二分搜索
func Search(key int, seq []int) int {
	left, right := 0, len(seq)
	for left < right {
		mid := (left + right) >> 1
		if key == seq[mid] {
			return mid
		} else if key < seq[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	return NOTFOUND
}
