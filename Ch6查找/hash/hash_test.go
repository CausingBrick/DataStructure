package hash

import (
	"testing"
)

var seq1 = []int{23, 12, 22, 4, 28, 7, 6, 9, 11}
var seq2 = make([]int, 10)

func TestInsert(t *testing.T) {
	for i := 0; i < len(seq1); i++ {
		Insert(seq2, seq1[i])
	}

}

func TestSearch(t *testing.T) {
	for i := 1; i < 9; i++ {
		if seq2[Search(seq2, seq1[i])] == 0 {
			t.Error("this slice cannot contain 0", seq2)
		}
	}
}
