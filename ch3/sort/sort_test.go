package sort

import (
	md "math/rand"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var unsort []int

	for i := 0; i < 200; i++ {
		unsort = append(unsort, md.Int())
	}
	sorted := unsort
	InsertionSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("InsertionSort %v, fact: %v", unsort, sorted)
	}
}
