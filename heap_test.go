package algo4go

import (
	"testing"
)

// TestHeap generate test slice and check it's partition
func TestHeap(t *testing.T) {
	log("Heap Test start")
	// init some viriable
	var heap1 MaxHeap
	var test []int
	times := 200

	t.Run("heap insert", func(t *testing.T) {
		for time := 1; time < times; time++ {
			// rest values
			test = randomSlice(7, 103)
			heap1 = MaxHeap{[]int{}, 0}
			for i := 0; i < len(test); i++ {
				heap1.HeapInsert(test[i])
			}
			if !(maxIntSlice(test) == heap1.arr[0]) {
				t.Errorf("max heap build fail :( biggest in origin is %v top on max heap is %v ", maxIntSlice(test), heap1.arr[0])
				t.Errorf("heap is %v", heap1.arr)
				t.Errorf("test is %v", test)
			}
		}
	})
}
