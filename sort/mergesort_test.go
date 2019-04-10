package sort

import (
	"sort"
	"testing"

	"github.com/MarkWh1te/algo4go/util"
)

func testMergeSort(t *testing.T) {
	var test []int
	times := 200
	for time := 1; time < times; time++ {
		test = util.RandomSlice(7, 1030)
		mergeSorted := MergeSort(&test, 0, len(test)-1)
		if !sort.IntsAreSorted(mergeSorted) {
			t.Errorf("%v is not sorted", test)
		}
	}
}

// TestInsertionSort is main function for this test
func TestMergeSort(t *testing.T) {
	t.Run("test mergesort sort", testMergeSort)
}
