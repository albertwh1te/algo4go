package algo4go

import (
	"sort"
	"testing"
)

// testPartition generate test slice and check it's partition
func testPartition(t *testing.T) {
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

// testQuickSort compare quicksort result with built in Sort
func testQuickSort(t *testing.T) {
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

// TestQuickSort is main function for this test
func TestQuickSort(t *testing.T) {
	t.Run("test partition", func(t *testing.T) {
		testPartition(t)
	})
	t.Run("test quicksort", func(t *testing.T) {
		testPartition(t)
	})
}
