package hash

// Hash get a hash key for index of value.
// key is the element.
// len is the length of elements.
func Hash(key, i, len int) int {
	return (h1(key, len) + i*h2(key, i, len)%len)
}
func h1(key, len int) int {
	return key % len
}

func h2(key, i, len int) int {
	return 1 + (key % (len - 1))
}

//Search  ruturn the index of the key in seq.
func Search(seq []int, key int) int {
	i := 0
	for true {
		j := Hash(key, i, len(seq))
		if seq[j] == key {
			return j
		} else if seq[j] == 0 || i >= len(seq) {
			return 404
		} else {
			i++
		}
	}
	return 404
}
