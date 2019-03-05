package algo4go

import (
	"sort"
	"testing"
)

func testInsertionSort(t *testing.T) {
	var test []int
	times := 200
	for time := 1; time < times; time++ {
		test = randomSlice(7, 103)
		InsertionSort(&test, 0, len(test)-1)
		if !sort.IntsAreSorted(test) {
			t.Errorf("%v is not sorted", test)
		}
	}
}

// TestInsertionSort is main function for this test
func TestInsertionSort(t *testing.T) {
	t.Run("test insertion sort", testInsertionSort)
}
