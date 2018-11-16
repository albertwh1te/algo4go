package algo4go

import (
	"fmt"
	"testing"
)

// TestHeap generate test slice and check it's partition
func TestHeap(t *testing.T) {
	const UPPER int = 20
	test := randomSlice(UPPER)
	heap := new(MaxHeap)
	for i := 0; i < len(test); i++ {
		heap.Insert(test[i])
	}
	fmt.Println("the origin heap is: ", heap.arr)
}
