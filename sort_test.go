package algo4go

import (
	"fmt"
	"sort"
	"testing"
)

// generate random slice
func randomSlice(upper int) []int {
	length := random(10, upper)
	// use slice because length is not a constant
	result := make([]int, length)
	for index := 0; index < length; index++ {
		result[index] = random(1, upper)
	}
	return result
}

// TestPartition generate test slice and check it's partition
func TestPartition(t *testing.T) {
	const UPPER int = 1000
	test := randomSlice(UPPER)
	length := len(test)
	pivot := test[length-1]
	fmt.Println("pivot: ", pivot)
	fmt.Println("origin arr: ", test)
	start, end := Partition(&test, 0, length-1, length-1)
	fmt.Println("start index: ", start, " end index: ", end)
	fmt.Println("new arr: ", test)
	for index := 0; index < length; index++ {
		if test[index] > pivot {
			if index < end {
				t.Errorf("big index check error")
			}
		} else if test[index] < pivot {
			if index > start {
				t.Errorf("small index check error")
			}
		}
	}
}

// TestQuickSort compare quicksort result with built in Sort
func TestQuickSort(t *testing.T) {
	const UPPER int = 13
	test := randomSlice(UPPER)
	QuickSort(&test, 0, len(test)-1)
	// sort.Ints(test)
	if !sort.IntsAreSorted(test) {
		t.Errorf("%v is not sorted", test)
	}
}
