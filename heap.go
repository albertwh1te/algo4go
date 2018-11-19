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
	if (2*i + 1) < heap.size {
		return 2*i + 1
	}
	return 0
}

// get heap right element
func (heap MaxHeap) right(i int) int {
	if (2*i + 2) < heap.size {
		return 2*i + 2
	}
	return 0
}

// Insert element into Maxheap
// TODO:fix me !
// func (heap *MaxHeap) Insert(n int) {
// 	(*heap).arr = append((*heap).arr, n)
// 	(*heap).size++
// }

// MaxHeapify assumes that the binary trees rooted at LEFT and RIGHT are max-heaps,
// but that node at i might be smaller than its children.
func (heap *MaxHeap) MaxHeapify(i int) {
	largest := i
	r := heap.right(i)
	l := heap.left(i)
	if l > 0 && heap.arr[l] > heap.arr[i] {
		largest = l
	} else if r > 0 && heap.arr[r] > heap.arr[i] {
		largest = r
	}
	if largest != i {
		swap(&heap.arr[i], &heap.arr[largest])
		i = largest
		heap.MaxHeapify(i)
	}
}
