package sort

import (
	"sort"
	"testing"

	"github.com/MarkWh1te/algo4go/util"
)

// TestHeap generate test slice and check it's partition
func TestHeap(t *testing.T) {
	util.Log("Heap Test start")
	// init some viriable
	var test []int
	times := 200
	t.Run("test heap insert", func(t *testing.T) {
		for time := 1; time < times; time++ {
			// rest values
			test = util.RandomSlice(7, 103)
			for i := 0; i < len(test); i++ {
				MaxHeapInsert(&test, i)
			}
			if !(util.MaxIntSlice(test) == test[0]) {
				t.Errorf("max heap build fail :( biggest in origin is %v top on max heap is %v ", util.MaxIntSlice(test), test[0])
				t.Errorf("test is %v", test)
			}
		}
	})
	t.Run("test heapsort", func(t *testing.T) {
		var test []int
		// add empty test to keep 100% test coverage
		HeapSort(&test)
		times := 200
		for time := 1; time < times; time++ {
			test = util.RandomSlice(7, 103)
			HeapSort(&test)
			if !sort.IntsAreSorted(test) {
				t.Errorf("%v is not sorted", test)
			}
		}
	})
}
