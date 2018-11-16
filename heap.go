package algo4go

// MaxHeap always has maxium at root
type MaxHeap struct {
	arr  []int
	size int
}

// Insert element into Maxheap
func (heap *MaxHeap) Insert(n int) {
	(*heap).arr = append((*heap).arr, n)
	(*heap).size++
}
