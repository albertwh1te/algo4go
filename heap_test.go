package algo4go

import (
	"fmt"
	"testing"
)

// TestHeap generate test slice and check it's partition
func TestHeap(t *testing.T) {
	const UPPER int = 12
	test := randomSlice(UPPER)
	var heap MaxHeap
	heap.arr = test
	heap.size = len(test)
	t.Run("basic setting", func(t *testing.T) {
		heapifyIndex := random(0, heap.size-1)
		fmt.Println("the origin heap is: ", heap.arr, "\nheapify index is", heapifyIndex)
		heap.MaxHeapify(heapifyIndex)
		fmt.Println("the heapify heap is: ", heap.arr)
	})
}
