package sort

import (
	"math/rand"
	"testing"
)

//Length of unsorted array
var arrLen = 200
var unsort []int

func init() {
	for i := 0; i < arrLen; i++ {
		unsort = append(unsort, rand.Int()%arrLen)
	}
}
func TestIsSorted(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	if isSorted(sorted) {
		t.Logf("%v", sorted)
		t.Errorf("isSorted %v, In fact: %v", unsort, sorted)
	}
}

func TestInsertionSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	InsertionSort(sorted)
	if !isSorted(sorted) {
		t.Logf("%v", sorted)
		t.Errorf("InsertionSort %v, In fact: %v", unsort, sorted)
	}
}

func TestBinaryInsertionSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	BinaryInsertionSort(sorted)
	if !isSorted(sorted) {
		t.Logf("%v", sorted)
		t.Errorf("BinaryInsertionSort %v, In fact: %v", unsort, sorted)
	}
}

func TestBubbleSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	BubbleSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("BubbleSort %v, In fact: %v", unsort, sorted)
	}
}

func TestBubbleSortF(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	BubbleSortF(sorted)
	if !isSorted(sorted) {
		t.Errorf("InsertionSort %v, In fact: %v", unsort, sorted)
	}
}

func TestShellSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	ShellSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("ShellSort err: In fact: %v", sorted)
	}
}

func TestQuickSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	QuickSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("QuickSort err: In fact: %v", sorted)
	}
}

func TestSelectionSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	SelectionSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("SelectionSort %v, In fact: %v", unsort, sorted)
	}
}
func TestHeapSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	HeapSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("HeapSort %v, In fact: %v", unsort, sorted)
	}
}
func TestMergeSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	MergeSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("MergeSort %v, In fact: %v", unsort, sorted)
	}
}
func TestRadixSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	RadixSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("RadixSort %v, In fact: %v", unsort, sorted)
	}
}
