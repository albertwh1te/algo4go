package algo4go

// MaxHeap always has maxium at root
type MaxHeap struct {
	arr  []int
	size int
}

// get heap parent element
func (heap MaxHeap) parent(i int) int {
	return i - 1/2
}

// get heap left element
func (heap MaxHeap) left(i int) int {
	if (2*i + 1) > heap.size {
		return 2*i + 1
	}
	return 0
}

// get heap right element
func (heap MaxHeap) right(i int) int {
	if (2*i + 2) > heap.size {
		return 2*i + 2
	}
	return 0
}

// Insert element into Maxheap
func (heap *MaxHeap) Insert(n int) {
	(*heap).arr = append((*heap).arr, n)
	(*heap).size++
}

// Heapify particulry element
// func (heap *MaxHeap) MaxHeapify(index int) {
// 	for parent(index) {
// 		if
// 	}
// }
