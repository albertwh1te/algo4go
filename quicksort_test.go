package algo4go

import (
	"sort"
	"testing"
)

// TestPartition generate test slice and check it's partition
func TestPartition(t *testing.T) {
	var test []int
	times := 200
	for time := 1; time < times; time++ {
		test = randomSlice(7, 103)
		length := len(test)
		pivot := test[length-1]
		start, end := Partition(&test, 0, length-1, length-1)
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
}

// TestQuickSort compare quicksort result with built in Sort
func TestQuickSort(t *testing.T) {
	var test []int
	times := 200
	for time := 1; time < times; time++ {
		test = randomSlice(7, 103)
		QuickSort(&test, 0, len(test)-1)
		if !sort.IntsAreSorted(test) {
			t.Errorf("%v is not sorted", test)
		}
	}
}
