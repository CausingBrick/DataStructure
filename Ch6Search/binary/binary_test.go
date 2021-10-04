package binary

import (
	md "math/rand"
	"testing"
	"time"

	"github.com/CausingBrick/DataStructure/Ch7Search/sort"
)

func TestSearch(t *testing.T) {
	r := md.New(md.NewSource(time.Now().UnixNano()))
	var sequence1 []int
	for i := 0; i < 10; i++ {
		sequence1 = append(sequence1, r.Int())
	}
	sort.QuickSort(sequence1)
	key := sequence1[2]
	if Search(key, sequence1) != 2 {
		t.Errorf("fail: the index does not match")
	}
	t.Log(sequence1, key, Search(key, sequence1))
}
