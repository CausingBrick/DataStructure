package binary

import (
	md "math/rand"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	r := md.New(md.NewSource(time.Now().UnixNano()))
	var sequence1 []int
	for i := 0; i < 10; i++ {
		sequence1 = append(sequence1, r.Int())
	}
	key := sequence1[5]
	if Search(key, sequence1) != 5 {
		t.Errorf("fail: the index does not match")
	}
	t.Log(sequence1, key, Search(key, sequence1))
}
