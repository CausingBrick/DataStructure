package hash

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	seq1 := [10]int{23, 12, 22, 4, 28, 7, 6, 9, 11, 3}
	var seq2 [10]bool
	for i := 0; i < 10; i++ {
		x := 1
		j := 0
		for true {
			j = Hash(seq1[i], x, len(seq2)) - 1
			fmt.Println(j, len(seq2))
			if x > len(seq2) {
				t.Fatal("get value from Hash function error.")
			} else if seq2[j] == false {
				break
			}
			x++
		}
		seq2[j] = true
	}
	for _, val := range seq2 {
		if !val {
			t.Error("there are also gap values not filled")
		}
	}
}

func SearchTest(t *testing.T) {
	seq1 := [10]int{23, 12, 22, 4, 28, 7, 6, 9, 11, 3}
	seq2 := []int{}
	for i := 0; i < 10; i++ {
		x := 1
		j := 0
		for true {
			j = Hash(seq1[i], x, len(seq2))
			if seq2[j] == 0 {
				break
			} else if x > len(seq2) {
				t.Fatal("get value from Hash function error.")
			}
			x++
		}
	}
}
