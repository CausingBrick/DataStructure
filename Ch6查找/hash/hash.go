package hash

import (
	"fmt"
	"os"
)

func h1(key, m int) int {
	return key % m
}

func h2(key, m int) int {
	return 1 + key%m
}

func H(key, m, i int) int {
	return (h1(key, m) + i*h2(key, m)) % m
}

func Insert(seq []int, key int) int {
	i := 0
	for true {
		j := H(key, len(seq), i)
		if j >= len(seq) {
			fmt.Errorf("hash value out of index")
			os.Exit(0)
		}
		if seq[j] == 0 {
			seq[j] = key
			return j
		} else {
			i++
		}
	}
	return 0
}

//Search  ruturn the index of the key in seq.
func Search(seq []int, key int) int {
	i := 0
	for true {
		j := H(key, len(seq), i)
		if seq[j] == key {
			return j
		} else if seq[j] == 0 || i >= len(seq) {
			return 0
		} else {
			i++
		}
	}
	return 0
}
