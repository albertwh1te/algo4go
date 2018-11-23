package algo4go

import (
	"sort"
	"testing"
)

// TestHeap generate test slice and check it's partition
func TestHeap(t *testing.T) {
	log("Heap Test start")
	// init some viriable
	var test []int
	times := 200
	t.Run("heap insert", func(t *testing.T) {
		for time := 1; time < times; time++ {
			// rest values
			test = randomSlice(7, 103)
			for i := 0; i < len(test); i++ {
				MaxHeapInsert(&test, i)
			}
			if !(maxIntSlice(test) == test[0]) {
				t.Errorf("max heap build fail :( biggest in origin is %v top on max heap is %v ", maxIntSlice(test), test[0])
				t.Errorf("test is %v", test)
			}
		}
	})
	t.Run("heapsort test", func(t *testing.T) {
		var test []int
		// add empty test to keep 100% test coverage
		HeapSort(&test)
		times := 200
		for time := 1; time < times; time++ {
			test = randomSlice(7, 103)
			HeapSort(&test)
			if !sort.IntsAreSorted(test) {
				t.Errorf("%v is not sorted", test)
			}
		}
	})

}
